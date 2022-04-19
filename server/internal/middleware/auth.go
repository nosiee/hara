package middleware

import (
	"hara/internal/config"
	"hara/internal/controllers"
	"unicode/utf8"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/sirupsen/logrus"
)

func SignUpFormProvided(ctx *gin.Context) {
	_, unameOk := ctx.GetPostForm("username")
	_, passwdOk := ctx.GetPostForm("password")
	_, emailOk := ctx.GetPostForm("email")

	if !unameOk {
		logrus.WithFields(logrus.Fields{
			"remote-addr": ctx.Request.RemoteAddr,
		}).Warning("Username required but not provided")

		ctx.AbortWithStatusJSON(400, gin.H{
			"error": "Username required but not provided",
		})
		return
	}

	if !passwdOk {
		logrus.WithFields(logrus.Fields{
			"remote-addr": ctx.Request.RemoteAddr,
		}).Warning("Password required but not provided")

		ctx.AbortWithStatusJSON(400, gin.H{
			"error": "Password required but not provided",
		})
		return
	}

	if !emailOk {
		logrus.WithFields(logrus.Fields{
			"remote-addr": ctx.Request.RemoteAddr,
		}).Warning("Email required but not provided")

		ctx.AbortWithStatusJSON(400, gin.H{
			"error": "Email required but not provided",
		})
		return
	}
}

func SignInFormProvided(ctx *gin.Context) {
	_, unameOk := ctx.GetPostForm("username")
	_, passwdOk := ctx.GetPostForm("password")

	if !unameOk {
		logrus.WithFields(logrus.Fields{
			"remote-addr": ctx.Request.RemoteAddr,
		}).Warning("Username required but not provided")

		ctx.AbortWithStatusJSON(400, gin.H{
			"error": "Username required but not provided",
		})
	}

	if !passwdOk {
		logrus.WithFields(logrus.Fields{
			"remote-addr": ctx.Request.RemoteAddr,
		}).Warning("Password required but not provided")

		ctx.AbortWithStatusJSON(400, gin.H{
			"error": "Password required but not provided",
		})
	}
}

func SignUpFormValidate(ctx *gin.Context) {
	username, _ := ctx.GetPostForm("username")
	password, _ := ctx.GetPostForm("password")
	email, _ := ctx.GetPostForm("email")

	if utf8.RuneCountInString(username) < minUsernameLength || utf8.RuneCountInString(username) > maxUsernameLength {
		logrus.WithFields(logrus.Fields{
			"remote-addr": ctx.Request.RemoteAddr,
		}).Error(ErrUsernameLength)

		ctx.AbortWithStatusJSON(400, gin.H{
			"error": ErrUsernameLength.Error(),
		})
	}

	if utf8.RuneCountInString(password) < minPasswordLenght || utf8.RuneCountInString(password) > maxPasswordLength {
		logrus.WithFields(logrus.Fields{
			"remote-addr": ctx.Request.RemoteAddr,
		}).Error(ErrPasswordLength)

		ctx.AbortWithStatusJSON(400, gin.H{
			"error": ErrPasswordLength.Error(),
		})
	}

	if !emailRegex.MatchString(email) {
		logrus.WithFields(logrus.Fields{
			"remote-addr": ctx.Request.RemoteAddr,
		}).Error(ErrEmailRegex)

		ctx.AbortWithStatusJSON(400, gin.H{
			"error": ErrEmailRegex.Error(),
		})
	}
}

func SignInFormValidate(ctx *gin.Context) {
	username, _ := ctx.GetPostForm("username")
	password, _ := ctx.GetPostForm("password")

	if utf8.RuneCountInString(username) < minUsernameLength || utf8.RuneCountInString(username) > maxUsernameLength {
		logrus.WithFields(logrus.Fields{
			"remote-addr": ctx.Request.RemoteAddr,
		}).Error(ErrUsernameLength)

		ctx.AbortWithStatusJSON(400, gin.H{
			"error": ErrUsernameLength.Error(),
		})
	}

	if utf8.RuneCountInString(password) < minPasswordLenght || utf8.RuneCountInString(password) > maxPasswordLength {
		logrus.WithFields(logrus.Fields{
			"remote-addr": ctx.Request.RemoteAddr,
		}).Error(ErrUsernameLength)

		ctx.AbortWithStatusJSON(400, gin.H{
			"error": ErrPasswordLength.Error(),
		})
	}
}

func IsAuthorized(ctx *gin.Context) {
	cookie, err := ctx.Cookie("jwt")
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"remote-addr": ctx.Request.RemoteAddr,
		}).Warn("Not authorized")

		ctx.AbortWithStatusJSON(401, gin.H{
			"error": "not authorized",
		})
		return
	}

	if _, err := jwt.Parse(cookie, func(t *jwt.Token) (interface{}, error) {
		return []byte(config.Values.HS512Key), nil
	}); err != nil {
		logrus.WithFields(logrus.Fields{
			"remote-addr": ctx.Request.RemoteAddr,
		}).Error(err)

		ctx.AbortWithStatusJSON(401, gin.H{
			"error": err.Error(),
		})
		return
	}

	id, _ := controllers.ExtractUserIDFromJWT(cookie)
	ctx.Set("uuid", id)
}
