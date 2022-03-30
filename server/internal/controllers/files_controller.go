package controllers

import (
	"fmt"
	"hara/internal/config"
	"os"

	"github.com/gin-gonic/gin"
)

func GetImage(ctx *gin.Context) {
	fpath := fmt.Sprintf("%s/%s", config.Values.OutputImagePath, ctx.Param("filename"))
	file, err := os.Open(fpath)

	if err != nil {
		ctx.JSON(404, gin.H{
			"error": "File not found",
		})
		return
	}

	contentType, err := GetFileContentType(file)
	if err != nil {
		ctx.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.Header("Content-Type", contentType)
	ctx.File(fpath)
}

func GetVideo(ctx *gin.Context) {
	fpath := fmt.Sprintf("%s/%s", config.Values.OutputVideoPath, ctx.Param("filename"))
	file, err := os.Open(fpath)

	if err != nil {
		ctx.JSON(404, gin.H{
			"error": "File not found",
		})
		return
	}

	contentType, err := GetFileContentType(file)
	if err != nil {
		ctx.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.Header("Content-Type", contentType)
	ctx.File(fpath)
}
