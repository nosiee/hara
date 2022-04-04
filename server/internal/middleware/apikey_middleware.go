package middleware

import (
	"hara/internal/db"

	"github.com/gin-gonic/gin"
)

func ApiKeyProvided(ctx *gin.Context) {
	if len(ctx.Param("key")) != apiKeyLength {
		ctx.AbortWithStatusJSON(401, gin.H{
			"error": "api key is incorrect",
		})
	}
}

func ApiKeyValidate(ctx *gin.Context) {
	var ok bool
	var err error

	if ok, err = db.IsKeyExists(ctx.Param("key")); !ok {
		ctx.AbortWithStatusJSON(401, gin.H{
			"error": "api key is invalid",
		})
		return
	}

	if err != nil {
		ctx.AbortWithStatusJSON(401, gin.H{
			"error": err.Error(),
		})
		return
	}
}

func ApiQuotas(ctx *gin.Context) {
	var maxquotas, quotas int
	var err error
	key := ctx.Param("key")

	if maxquotas, quotas, err = db.GetKeyQuotas(key); err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	if quotas < maxquotas {
		if err = db.UpdateKeyQuotas(key, quotas+1); err != nil {
			ctx.AbortWithStatusJSON(500, gin.H{
				"error": err.Error(),
			})
		}
	} else {
		ctx.AbortWithStatusJSON(400, gin.H{
			"error": "quotas have been exhausted",
		})
	}
}
