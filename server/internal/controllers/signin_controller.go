package controllers

import (
	"encoding/hex"
	"hara/internal/config"
	"hara/internal/db"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

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

	token, err := GenerateJWT(uuid, config.Values.JWTKey)
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.SetCookie("jwt", token, int(time.Now().Add(1*365*24*time.Hour).Unix()), "/", "localhost", true, true)
	ctx.JSON(200, gin.H{
		"message": "ok",
	})
}
