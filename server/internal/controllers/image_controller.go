package controllers

import (
	"fmt"
	"hara/internal/config"
	"hara/internal/convert"
	"hara/internal/db"
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

	ofile, err := convert.ConvertImage(fpath, iopt)
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	db.AddFileLifetime(ofile, iopt.Lifetime)
	// url := GenerateFileUrl(ctx, ofile)
	// os.Remove(fpath)

	ctx.String(200, fmt.Sprintf("http://localhost:8080/api/i/%s", ofile))
}
