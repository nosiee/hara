package controllers

import (
	"encoding/hex"
	"hara/internal/config"
	"hara/internal/db"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(ctx *gin.Context) {
	username, _ := ctx.GetPostForm("username")
	password, _ := ctx.GetPostForm("password")
	email, _ := ctx.GetPostForm("email")
	uuid := GenerateRandomUUID()

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

	token, err := GenerateJWT(uuid, config.Values.HS512Key)
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.SetCookie("jwt", token, int(JWTExp), "/", "", false, true)
	ctx.JSON(200, gin.H{
		"message": "ok",
	})
}
