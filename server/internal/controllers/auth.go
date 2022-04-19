package controllers

import (
	"encoding/hex"
	"hara/internal/config"
	"hara/internal/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

func (c Controllers) SignUp(ctx *gin.Context) {
	username, _ := ctx.GetPostForm("username")
	password, _ := ctx.GetPostForm("password")
	email, _ := ctx.GetPostForm("email")
	uuid := uuid.NewString()

	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	user := models.NewUser(uuid, username, hex.EncodeToString(hash[:]), email)
	if err := c.UserRepository.Add(user); err != nil {
		logrus.WithFields(logrus.Fields{
			"remote-addr": ctx.Request.RemoteAddr,
		}).Error(err)

		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	exp := time.Now().Add(1 * 365 * 24 * time.Hour).Unix()

	token, err := GenerateJWT(uuid, config.Values.HS512Key, exp)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"remote-addr": ctx.Request.RemoteAddr,
		}).Error(err)

		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.SetCookie("jwt", token, int(exp), "/", "", false, true)
	ctx.JSON(200, gin.H{
		"message": "ok",
	})

	logrus.WithFields(logrus.Fields{
		"remote-addr": ctx.Request.RemoteAddr,
		"username":    user.Username,
		"uuid":        user.UUID,
	}).Info("New user has been added")
}

func (c Controllers) SignIn(ctx *gin.Context) {
	username, _ := ctx.GetPostForm("username")
	password, _ := ctx.GetPostForm("password")

	user := c.UserRepository.FindByUsername(username)
	if user == nil {
		logrus.WithFields(logrus.Fields{
			"remote-addr": ctx.Request.RemoteAddr,
		}).Warn("Incorrect username or password")

		ctx.JSON(401, gin.H{
			"error": "Username or password is incorrect",
		})
		return
	}

	decodedHash, _ := hex.DecodeString(user.Hash)
	if err := bcrypt.CompareHashAndPassword(decodedHash, []byte(password)); err != nil {
		logrus.WithFields(logrus.Fields{
			"remote-addr": ctx.Request.RemoteAddr,
		}).Warn("Incorrect username or password")

		ctx.JSON(401, gin.H{
			"error": "Username or password is incorrect",
		})
		return
	}

	exp := time.Now().Add(1 * 365 * 24 * time.Hour).Unix()

	token, err := GenerateJWT(user.UUID, config.Values.HS512Key, exp)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"remote-addr": ctx.Request.RemoteAddr,
		}).Error(err)

		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.SetCookie("jwt", token, int(exp), "/", "", false, true)
	ctx.JSON(200, gin.H{
		"message": "ok",
	})

	logrus.WithFields(logrus.Fields{
		"remote-addr": ctx.Request.RemoteAddr,
		"username":    user.Username,
		"uuid":        user.UUID,
	}).Info("Successful authorized")
}
