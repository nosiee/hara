package controllers

import (
	"fmt"
	"hara/internal/convert"
	"mime/multipart"

	"github.com/gin-gonic/gin"
)

func ImageController(ctx *gin.Context) {
	ioptIface, _ := ctx.Get("imageOptions")
	fileIface, _ := ctx.Get("file")

	iopt := ioptIface.(convert.ConversionImageOptions)
	file := fileIface.(*multipart.FileHeader)
	fpath := fmt.Sprintf("input/%s", file.Filename)

	// TODO: load input folder from config(?)
	ctx.SaveUploadedFile(file, fpath)

	if err := convert.ConvertImage(fpath, iopt); err != nil {
		fmt.Println(err)
		ctx.String(400, "NOT OK")
		return
	}

	ctx.String(200, "OK")
}
