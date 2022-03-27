package controllers

import (
	"fmt"
	"hara/internal/convert"
	"mime/multipart"

	"github.com/gin-gonic/gin"
)

func VideoController(ctx *gin.Context) {
	voptIface, _ := ctx.Get("videoOptions")
	fileIface, _ := ctx.Get("file")

	vopt := voptIface.(convert.ConversionVideoOptions)
	file := fileIface.(*multipart.FileHeader)
	fpath := fmt.Sprintf("input/%s", file.Filename)

	// TODO: load input folder from config(?)
	ctx.SaveUploadedFile(file, fpath)

	if err := convert.ConvertVideo(fpath, vopt); err != nil {
		fmt.Println(err)
		ctx.String(400, "NOT OK")
		return
	}

	ctx.String(200, "OK")
}
