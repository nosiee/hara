package controllers

import (
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
	if err = db.AddNewApiKey(id, apiKey, 100); err != nil {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"vid-url": GenerateAPIUrl(ctx, "video", apiKey),
		"img-url": GenerateAPIUrl(ctx, "image", apiKey),
		"key":     apiKey,
	})
}
