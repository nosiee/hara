package controllers

import "github.com/gin-gonic/gin"

func GetApiKey(ctx *gin.Context) {
	token, _ := ctx.Cookie("jwt")
	uuid, err := ExtractUserIDFromJWT(token)

	println(uuid)

	if err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
}
