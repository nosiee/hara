package middleware

import (
	"hara/internal/config"
	"hara/internal/controllers"
	"unicode/utf8"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func SignUpFormProvided(ctx *gin.Context) {
	_, unameOk := ctx.GetPostForm("username")
	_, passwdOk := ctx.GetPostForm("password")
	_, emailOk := ctx.GetPostForm("email")

	if !unameOk {
		ctx.AbortWithStatusJSON(400, gin.H{
			"error": "username required but not provided",
		})
		return
	}

	if !passwdOk {
		ctx.AbortWithStatusJSON(400, gin.H{
			"error": "password required but not provided",
		})
		return
	}

	if !emailOk {
		ctx.AbortWithStatusJSON(400, gin.H{
			"error": "email required but not provided",
		})
		return
	}
}

func SignInFormProvided(ctx *gin.Context) {
	_, unameOk := ctx.GetPostForm("username")
	_, passwdOk := ctx.GetPostForm("password")

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
}

func SignUpFormValidate(ctx *gin.Context) {
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

	if !emailRegex.MatchString(email) {
		ctx.AbortWithStatusJSON(400, gin.H{
			"error": ErrEmailRegex.Error(),
		})
	}
}

func SignInFormValidate(ctx *gin.Context) {
	username, _ := ctx.GetPostForm("username")
	password, _ := ctx.GetPostForm("password")

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
}

func IsAuthorized(ctx *gin.Context) {
	cookie, err := ctx.Cookie("jwt")
	if err != nil {
		ctx.AbortWithStatusJSON(401, gin.H{
			"error": "not authorized",
		})
		return
	}

	if _, err := jwt.Parse(cookie, func(t *jwt.Token) (interface{}, error) {
		return []byte(config.Values.HS512Key), nil
	}); err != nil {
		errLogger.Println(err)

		ctx.AbortWithStatusJSON(401, gin.H{
			"error": err.Error(),
		})
		return
	}

	id, _ := controllers.ExtractUserIDFromJWT(cookie)
	ctx.Set("uuid", id)
}
