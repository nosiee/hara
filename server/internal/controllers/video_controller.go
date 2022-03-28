package controllers

import (
	"fmt"
	"hara/internal/config"
	"hara/internal/convert"
	"mime/multipart"

	"github.com/gin-gonic/gin"
)

func VideoController(ctx *gin.Context) {
	voptIface, _ := ctx.Get("videoOptions")
	fileIface, _ := ctx.Get("file")

	vopt := voptIface.(convert.ConversionVideoOptions)
	file := fileIface.(*multipart.FileHeader)
	fpath := fmt.Sprintf("%s/%s", config.Values.UploadVideoPath, file.Filename)

	ctx.SaveUploadedFile(file, fpath)

	if err := convert.ConvertVideo(fpath, vopt); err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	// TODO: generate temporary link to the converted file
	ctx.String(200, "OK")
}
