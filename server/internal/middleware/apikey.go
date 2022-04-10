package middleware

import (
	"hara/internal/models"
	"net/url"
	"time"

	"github.com/gin-gonic/gin"
)

func ApiKeyProvided(ctx *gin.Context) {
	query, err := url.ParseQuery(ctx.Request.URL.RawQuery)
	if err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	if len(query.Get("key")) != apiKeyLength {
		ctx.AbortWithStatusJSON(401, gin.H{
			"error": "api key is incorrect",
		})
	}
}

func ApiKeyValidate(apikeyRepo models.ApiKeyRepository) func(*gin.Context) {
	return func(ctx *gin.Context) {
		query, _ := url.ParseQuery(ctx.Request.URL.RawQuery)

		if ok, err := apikeyRepo.IsExists(query.Get("key")); err != nil {
			ctx.AbortWithStatusJSON(500, gin.H{
				"error": err.Error(),
			})
		} else if !ok {
			ctx.AbortWithStatusJSON(401, gin.H{
				"error": "api key is invalid",
			})
		}
	}
}

func ApiKeyQuota(apikeyRepo models.ApiKeyRepository) func(*gin.Context) {
	return func(ctx *gin.Context) {
		query, _ := url.ParseQuery(ctx.Request.URL.RawQuery)
		key := query.Get("key")

		ut, err := apikeyRepo.GetUpdatetime(key)
		if err != nil {
			ctx.AbortWithStatusJSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}

		if ut > 0 {
			ctx.AbortWithStatusJSON(400, gin.H{
				"error": "Quota has been exceeded",
			})
			return
		}

		maxquota, quota, err := apikeyRepo.GetQuota(key)
		if err != nil {
			ctx.AbortWithStatusJSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}

		if quota < maxquota {
			if err := apikeyRepo.IncrementQuota(key); err != nil {
				ctx.AbortWithStatusJSON(500, gin.H{
					"error": err.Error(),
				})
			}
			return
		}

		if err := apikeyRepo.SetUpdatetime(key, time.Now().Add(24*time.Hour).Unix()); err != nil {
			ctx.AbortWithStatusJSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.AbortWithStatusJSON(400, gin.H{
			"error": "Quota has been exceeded",
		})
	}
}
