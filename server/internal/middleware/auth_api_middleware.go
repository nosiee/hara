package middleware

import (
	"fmt"
	"hara/internal/config"
	"unicode/utf8"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func AuthFormProvided(ctx *gin.Context) {
	_, unameOk := ctx.GetPostForm("username")
	_, passwdOk := ctx.GetPostForm("password")
	_, emailOk := ctx.GetPostForm("email")

	if !unameOk {
		ctx.AbortWithStatusJSON(400, gin.H{
			"error": "username required but not provided",
		})
	}

	if !passwdOk {
		ctx.AbortWithStatusJSON(400, gin.H{
			"error": "password required but not provided",
		})
	}

	fmt.Println(ctx.Request.RequestURI)

	if ctx.Request.RequestURI == "/api/auth/signup" {
		if !emailOk {
			ctx.AbortWithStatusJSON(400, gin.H{
				"error": "email required but not provided",
			})
		}
	}
}

func AuthFormValidate(ctx *gin.Context) {
	username, _ := ctx.GetPostForm("username")
	password, _ := ctx.GetPostForm("password")
	email, _ := ctx.GetPostForm("email")

	if utf8.RuneCountInString(username) < minUsernameLength || utf8.RuneCountInString(username) > maxUsernameLength {
		ctx.AbortWithStatusJSON(400, gin.H{
			"error": ErrUsernameLength.Error(),
		})
	}

	if utf8.RuneCountInString(password) < minPasswordLenght || utf8.RuneCountInString(password) > maxPasswordLength {
		ctx.AbortWithStatusJSON(400, gin.H{
			"error": ErrPasswordLength.Error(),
		})
	}

	if ctx.Request.RequestURI == "/api/auth/signup" {
		if !emailRegex.MatchString(email) {
			ctx.AbortWithStatusJSON(400, gin.H{
				"error": ErrEmailRegex.Error(),
			})
		}
	}
}

func IsAuthorized(ctx *gin.Context) {
	token, err := ctx.Cookie("jwt")
	if err != nil {
		ctx.AbortWithStatusJSON(401, gin.H{
			"error": "not authorized",
		})
		return
	}

	if _, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return []byte(config.Values.JWTKey), nil
	}); err != nil {
		ctx.AbortWithStatusJSON(401, gin.H{
			"error": err.Error(),
		})
		return
	}
}
