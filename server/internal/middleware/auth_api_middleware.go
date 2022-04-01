package middleware

import (
	"unicode/utf8"

	"github.com/gin-gonic/gin"
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
