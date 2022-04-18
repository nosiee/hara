package controllers

import (
	"database/sql"
	"hara/internal/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (c Controllers) GetApiKey(ctx *gin.Context) {
	token, _ := ctx.Cookie("jwt")
	id, err := ExtractUserIDFromJWT(token)
	if err != nil {
		c.ErrLogger.Println(err)

		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	ok, err := c.ApikeyRepository.UserHasKey(id)
	if err != nil && err != sql.ErrNoRows {
		c.ErrLogger.Println(err)

		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	} else if ok {
		ctx.JSON(400, gin.H{
			"error": "You already have a key",
		})
		return
	}

	apikey := models.NewApiKey(id, uuid.NewString(), 100, 0, 0)

	if err = c.ApikeyRepository.Add(apikey); err != nil {
		c.ErrLogger.Println(err)

		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"video-url": GenerateAPIUrl(ctx, "video", apikey.Key),
		"image-url": GenerateAPIUrl(ctx, "image", apikey.Key),
		"key":       apikey.Key,
	})
}
