package controllers

import (
	"database/sql"
	"hara/internal/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (c Controllers) GetApiKey(ctx *gin.Context) {
	id, _ := ctx.Get("uuid")

	ok, err := c.ApikeyRepository.UserHasKey(id.(string))
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

	apikey := models.NewApiKey(id.(string), uuid.NewString(), 100, 0, 0)

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

func (c Controllers) ResetApiKey(ctx *gin.Context) {
	id, _ := ctx.Get("uuid")
	var err error

	ok, err := c.ApikeyRepository.UserHasKey(id.(string))
	if err != nil && err != sql.ErrNoRows {
		c.ErrLogger.Println(err)

		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	apikey := models.NewApiKey(id.(string), uuid.NewString(), 100, 0, 0)

	if !ok {
		err = c.ApikeyRepository.Add(apikey)
	} else {
		err = c.ApikeyRepository.ChangeKeyID(id.(string), apikey.Key)
	}

	if err != nil {
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

func (c Controllers) RemindApiKey(ctx *gin.Context) {
	id, _ := ctx.Get("uuid")
	var err error

	ok, err := c.ApikeyRepository.UserHasKey(id.(string))
	if err != nil && err != sql.ErrNoRows {
		c.ErrLogger.Println(err)

		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	if !ok {
		ctx.JSON(400, gin.H{
			"error": "You don't have any key",
		})
		return
	}

	apikey, err := c.ApikeyRepository.GetKey(id.(string))
	if err != nil {
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
