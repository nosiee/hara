package controllers

import (
	"database/sql"
	"hara/internal/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func (c Controllers) GetApiKey(ctx *gin.Context) {
	id, _ := ctx.Get("uuid")

	ok, err := c.ApikeyRepository.UserHasKey(id.(string))
	if err != nil && err != sql.ErrNoRows {
		logrus.WithFields(logrus.Fields{
			"remote-addr": ctx.Request.RemoteAddr,
			"uuid":        id,
		}).Error(err)

		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	} else if ok {
		logrus.WithFields(logrus.Fields{
			"remote-addr": ctx.Request.RemoteAddr,
			"uuid":        id,
		}).Warn("Already has a key")

		ctx.JSON(400, gin.H{
			"error": "You already have a key",
		})
		return
	}

	apikey := models.NewApiKey(id.(string), uuid.NewString(), 100, 0, 0)

	if err = c.ApikeyRepository.Add(apikey); err != nil {
		logrus.WithFields(logrus.Fields{
			"remote-addr": ctx.Request.RemoteAddr,
			"uuid":        id,
		}).Error(err)

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

	logrus.WithFields(logrus.Fields{
		"remote-addr": ctx.Request.RemoteAddr,
		"uuid":        id,
		"api-key":     apikey.Key,
	}).Info("New key has been added")
}

func (c Controllers) ResetApiKey(ctx *gin.Context) {
	id, _ := ctx.Get("uuid")
	var err error

	ok, err := c.ApikeyRepository.UserHasKey(id.(string))
	if err != nil && err != sql.ErrNoRows {
		logrus.WithFields(logrus.Fields{
			"remote-addr": ctx.Request.RemoteAddr,
			"uuid":        id,
		}).Error(err)

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
		logrus.WithFields(logrus.Fields{
			"remote-addr": ctx.Request.RemoteAddr,
			"uuid":        id,
		}).Error(err)

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

	logrus.WithFields(logrus.Fields{
		"remote-addr": ctx.Request.RemoteAddr,
		"uuid":        id,
		"api-key":     apikey.Key,
	}).Info("Api key has been reseted")
}

func (c Controllers) RemindApiKey(ctx *gin.Context) {
	id, _ := ctx.Get("uuid")
	var err error

	ok, err := c.ApikeyRepository.UserHasKey(id.(string))
	if err != nil && err != sql.ErrNoRows {
		logrus.WithFields(logrus.Fields{
			"remote-addr": ctx.Request.RemoteAddr,
			"uuid":        id,
		}).Error(err)

		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	if !ok {
		logrus.WithFields(logrus.Fields{
			"remote-addr": ctx.Request.RemoteAddr,
			"uuid":        id,
		}).Warning("You don't have any key")

		ctx.JSON(400, gin.H{
			"error": "You don't have any key",
		})
		return
	}

	apikey, err := c.ApikeyRepository.GetKey(id.(string))
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"remote-addr": ctx.Request.RemoteAddr,
			"uuid":        id,
		}).Error(err)

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

	logrus.WithFields(logrus.Fields{
		"remote-addr": ctx.Request.RemoteAddr,
		"uuid":        id,
		"api-key":     apikey.Key,
	}).Info("Reminded about api key")
}
