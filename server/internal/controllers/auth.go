package controllers

import (
	"encoding/hex"
	"hara/internal/config"
	"hara/internal/db"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(ctx *gin.Context) {
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

	if err := db.CreateNewUser(uuid, username, hex.EncodeToString(hash[:]), email); err != nil {
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

func SignIn(ctx *gin.Context) {
	username, _ := ctx.GetPostForm("username")
	password, _ := ctx.GetPostForm("password")

	uuid, hash, ok := db.FindUser(username)
	if !ok {
		ctx.JSON(401, gin.H{
			"error": "Username or password is incorrect",
		})
		return
	}

	decodedHash, _ := hex.DecodeString(hash)
	if err := bcrypt.CompareHashAndPassword(decodedHash, []byte(password)); err != nil {
		ctx.JSON(401, gin.H{
			"error": "Username or password is incorrect",
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
