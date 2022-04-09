package controllers

import (
	"encoding/hex"
	"hara/internal/config"
	"hara/internal/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func (c Controllers) SignUp(ctx *gin.Context) {
	username, _ := ctx.GetPostForm("username")
	password, _ := ctx.GetPostForm("password")
	email, _ := ctx.GetPostForm("email")
	uuid := uuid.NewString()

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	user := models.NewUser(uuid, username, hex.EncodeToString(hash[:]), email)
	if err := c.UserRepository.Create(user); err != nil {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	exp := time.Now().Add(1 * 365 * 24 * time.Hour).Unix()

	token, err := GenerateJWT(uuid, config.Values.HS512Key, exp)
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.SetCookie("jwt", token, int(exp), "/", "", false, true)
	ctx.JSON(200, gin.H{
		"message": "ok",
	})
}

func (c Controllers) SignIn(ctx *gin.Context) {
	username, _ := ctx.GetPostForm("username")
	password, _ := ctx.GetPostForm("password")

	user := c.UserRepository.FindByUsername(username)
	if user == nil {
		ctx.JSON(401, gin.H{
			"error": "Username or password is incorrect",
		})
		return
	}

	decodedHash, _ := hex.DecodeString(user.Hash)
	if err := bcrypt.CompareHashAndPassword(decodedHash, []byte(password)); err != nil {
		ctx.JSON(401, gin.H{
			"error": "Username or password is incorrect",
		})
		return
	}

	exp := time.Now().Add(1 * 365 * 24 * time.Hour).Unix()

	token, err := GenerateJWT(user.UUID, config.Values.HS512Key, exp)
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.SetCookie("jwt", token, int(exp), "/", "", false, true)
	ctx.JSON(200, gin.H{
		"message": "ok",
	})
}
