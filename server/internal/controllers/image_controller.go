package controllers

import (
	"fmt"
	"hara/internal/config"
	"hara/internal/convert"
	"mime/multipart"

	"github.com/gin-gonic/gin"
)

func ImageController(ctx *gin.Context) {
	ioptIface, _ := ctx.Get("imageOptions")
	fileIface, _ := ctx.Get("file")

	iopt := ioptIface.(convert.ConversionImageOptions)
	file := fileIface.(*multipart.FileHeader)
	fpath := fmt.Sprintf("%s/%s", config.Values.UploadImagePath, file.Filename)

	ctx.SaveUploadedFile(file, fpath)

	if err := convert.ConvertImage(fpath, iopt); err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	// TODO: generate temporary link to the converted file
	ctx.String(200, "OK")
}
