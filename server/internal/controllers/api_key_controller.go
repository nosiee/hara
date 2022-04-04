package controllers

import (
	"hara/internal/db"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetApiKey(ctx *gin.Context) {
	token, _ := ctx.Cookie("jwt")
	id, err := ExtractUserIDFromJWT(token)
	apiKey := uuid.NewString()

	if err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err = db.AddNewApiKey(id, apiKey, 100); err != nil {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
	}

	ctx.String(200, apiKey)
}
