package controllers

import (
	"fmt"
	"hara/internal/config"
	"hara/internal/convert"
	"hara/internal/db"
	"mime/multipart"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func VideoController(ctx *gin.Context) {
	voptIface, _ := ctx.Get("options")
	fileIface, _ := ctx.Get("file")

	vopt := voptIface.(convert.ConversionVideoOptions)
	file := fileIface.(*multipart.FileHeader)
	fpath := fmt.Sprintf("%s/%s", config.Values.UploadVideoPath, file.Filename)

	ctx.SaveUploadedFile(file, fpath)

	ofile, err := convert.ConvertVideo(fpath, vopt)
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	os.Remove(fpath)
	deleteDate := time.Now().Add(time.Duration(vopt.Lifetime) * time.Second)
	if err = db.AddFileLifetime(ofile, "video", deleteDate.Format(time.RFC3339)); err != nil {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.String(200, GenerateFileUrl(ctx, "v", ofile))
}

func GetVideo(ctx *gin.Context) {
	fname := ctx.Param("filename")
	fpath := fmt.Sprintf("%s/%s", config.Values.OutputVideoPath, fname)

	if ok := db.IsFileExists(fname); !ok {
		ctx.JSON(404, gin.H{
			"error": "File not found",
		})
		return
	}

	file, err := os.Open(fpath)
	if err != nil {
		ctx.JSON(404, gin.H{
			"error": err.Error(),
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
