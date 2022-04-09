package controllers

import (
	"fmt"
	"hara/internal/config"
	"hara/internal/convert"
	"hara/internal/models"
	"mime/multipart"
	"os"

	"github.com/gin-gonic/gin"
)

func (c Controllers) ImageController(ctx *gin.Context) {
	ioptIface, _ := ctx.Get("options")
	fileIface, _ := ctx.Get("file")

	options := ioptIface.(convert.ConversionOptions)
	file := fileIface.(*multipart.FileHeader)
	fpath := fmt.Sprintf("%s/%s", config.Values.UploadImagePath, file.Filename)

	ctx.SaveUploadedFile(file, fpath)

	ofile, err := convert.ConvertImage(fpath, options)
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	os.Remove(fpath)

	f := models.NewFile(ofile, "image", options.Lifetime)
	if err = c.FileRepository.Add(f); err != nil {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.String(200, GenerateFileUrl(ctx, "i", ofile))
}

func (c Controllers) GetImage(ctx *gin.Context) {
	fname := ctx.Param("filename")
	fpath := fmt.Sprintf("%s/%s", config.Values.OutputImagePath, fname)

	if ok := c.FileRepository.IsExists(fname); !ok {
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

func (c Controllers) VideoController(ctx *gin.Context) {
	voptIface, _ := ctx.Get("options")
	fileIface, _ := ctx.Get("file")

	options := voptIface.(convert.ConversionOptions)
	file := fileIface.(*multipart.FileHeader)
	fpath := fmt.Sprintf("%s/%s", config.Values.UploadVideoPath, file.Filename)

	ctx.SaveUploadedFile(file, fpath)

	ofile, err := convert.ConvertVideo(fpath, options)
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	os.Remove(fpath)

	f := models.NewFile(ofile, "video", options.Lifetime)
	if err = c.FileRepository.Add(f); err != nil {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.String(200, GenerateFileUrl(ctx, "v", ofile))
}

func (c Controllers) GetVideo(ctx *gin.Context) {
	fname := ctx.Param("filename")
	fpath := fmt.Sprintf("%s/%s", config.Values.OutputVideoPath, fname)

	if ok := c.FileRepository.IsExists(fname); !ok {
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
