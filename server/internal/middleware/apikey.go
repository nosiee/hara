package middleware

import (
	"hara/internal/db"
	"net/url"

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

	ctx.Set("apikey", query.Get("key"))
}

func ApiKeyValidate(ctx *gin.Context) {
	key, _ := ctx.Get("apikey")

	if ok, err := db.IsKeyExists(key.(string)); !ok {
		ctx.AbortWithStatusJSON(401, gin.H{
			"error": "api key is invalid",
		})
	} else if err != nil {
		ctx.AbortWithStatusJSON(401, gin.H{
			"error": err.Error(),
		})
	}
}
