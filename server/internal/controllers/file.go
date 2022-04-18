package controllers

import (
	"fmt"
	"hara/internal/config"
	"hara/internal/convert"
	"hara/internal/models"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
)

func (c Controllers) ImageController(ctx *gin.Context) {
	optIface, _ := ctx.Get("options")
	fileIface, _ := ctx.Get("file")

	options := optIface.(convert.ConversionOptions)
	file := fileIface.(*multipart.FileHeader)
	fpath := fmt.Sprintf("%s/%s", config.Values.UploadImagePath, file.Filename)

	ctx.SaveUploadedFile(file, fpath)

	ofile, err := c.Converter.ConvertImage(fpath, options)
	if err != nil {
		c.ErrLogger.Println(err)

		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	os.Remove(fpath)

	f := models.NewFile(ofile, filepath.Join(config.Values.OutputImagePath, ofile), time.Now().Add(time.Duration(options.Lifetime)*time.Second).Unix())
	if err = c.FileRepository.Add(f); err != nil {
		c.ErrLogger.Println(err)

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
		c.ErrLogger.Println(err)

		ctx.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	contentType, _ := GetFileContentType(file)
	ctx.Header("Content-Type", contentType)
	ctx.File(fpath)
}

func (c Controllers) VideoController(ctx *gin.Context) {
	optIface, _ := ctx.Get("options")
	fileIface, _ := ctx.Get("file")

	options := optIface.(convert.ConversionOptions)
	file := fileIface.(*multipart.FileHeader)
	fpath := fmt.Sprintf("%s/%s", config.Values.UploadVideoPath, file.Filename)

	ctx.SaveUploadedFile(file, fpath)

	ofile, err := c.Converter.ConvertVideo(fpath, options)
	if err != nil {
		c.ErrLogger.Println(err)

		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	os.Remove(fpath)

	f := models.NewFile(ofile, filepath.Join(config.Values.OutputVideoPath, ofile), time.Now().Add(time.Duration(options.Lifetime)*time.Second).Unix())
	if err = c.FileRepository.Add(f); err != nil {
		c.ErrLogger.Println(err)

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
		c.ErrLogger.Println(err)

		ctx.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	contentType, _ := GetFileContentType(file)
	ctx.Header("Content-Type", contentType)
	ctx.File(fpath)
}
