package controllers

import (
	"fmt"
	"hara/internal/db"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetApiKey(ctx *gin.Context) {
	token, _ := ctx.Cookie("jwt")
	id, err := ExtractUserIDFromJWT(token)

	if err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	apiKey := uuid.NewString()
	// NOTE: 5 is just for tests. dont forget to change it
	if err = db.AddNewApiKey(id, apiKey, 5); err != nil {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"key": apiKey,
		"url": fmt.Sprintf("http://localhost:8080/api/convert/video?key=%s", apiKey),
	})
}
