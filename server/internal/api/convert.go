package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"hara/internal/convert"

	"github.com/gin-gonic/gin"
)

func (server *Server) convertVideo(ctx *gin.Context) {
	// TODO: validate options.Name field
	var options convert.ConversionVideoOptions

	fpath, err := server.getConversionOptions(ctx, &options)
	if err != nil {
		server.sendError(ctx, 400, err.Error())
		return
	}

	if err := server.converter.ConvertVideo(fpath, options); err != nil {
		server.sendError(ctx, 500, err.Error())
		return
	}

	// TODO: Remove input file
	// Generate temp link to a converted file and pass it
	ctx.String(200, "OK")
}

func (server *Server) convertImage(ctx *gin.Context) {
	// TODO: validate options.Name field
	var options convert.ConversionImageOptions

	fpath, err := server.getConversionOptions(ctx, &options)
	if err != nil {
		server.sendError(ctx, 400, err.Error())
		return
	}

	if err := server.converter.ConvertImage(fpath, options); err != nil {
		server.sendError(ctx, 500, err.Error())
		return
	}

	// TODO: Remove input file
	// Generate temp link to a converted file and pass it
	ctx.String(200, "OK")
}

func (server *Server) getConversionOptions(ctx *gin.Context, options any) (string, error) {
	optionsField, ok := ctx.GetPostForm("options")
	if !ok {
		return "", errors.New("Request cannot be processed. Required options not provided")
	}

	if err := json.Unmarshal([]byte(optionsField), &options); err != nil {
		return "", err
	}

	file, err := ctx.FormFile("file")
	if err != nil {
		return "", err
	}

	fpath := fmt.Sprintf("%s/%s", server.inputPath, file.Filename)
	ctx.SaveUploadedFile(file, fpath)

	return fpath, nil
}
