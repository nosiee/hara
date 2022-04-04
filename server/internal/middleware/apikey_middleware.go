package middleware

import (
	"hara/internal/db"

	"github.com/gin-gonic/gin"
)

func ApiKeyProvided(ctx *gin.Context) {
	// TODO: 36 should be a constant
	if len(ctx.Param("key")) != 36 {
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
	println("TODO")
}
