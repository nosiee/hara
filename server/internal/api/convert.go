package api

import (
	"encoding/json"
	"fmt"
	"hara/internal/convert"

	"github.com/gin-gonic/gin"
)

func (server *Server) convertVideo(ctx *gin.Context) {
	var options convert.ConversionVideoOptions

	// TODO: All this should be in separate function
	optionsField, ok := ctx.GetPostForm("options")
	if !ok {
		server.sendError(ctx, 400, "Request cannot be processed. Required options not provided")
		return
	}

	if err := json.Unmarshal([]byte(optionsField), &options); err != nil {
		server.sendError(ctx, 400, err.Error())
		return
	}

	file, err := ctx.FormFile("file")
	if err != nil {
		server.sendError(ctx, 400, err.Error())
		return
	}

	fpath := fmt.Sprintf("%s/%s", server.inputPath, file.Filename)
	ctx.SaveUploadedFile(file, fpath)

	if err := server.converter.ConvertVideo(fpath, options); err != nil {
		server.sendError(ctx, 500, err.Error())
		return
	}

	// TODO: Remove input file
	ctx.String(200, "OK")
}

func (server *Server) convertImage(ctx *gin.Context) {
	var options convert.ConversionImageOptions

	// TODO: All this should be in separate function
	optionsField, ok := ctx.GetPostForm("options")
	if !ok {
		server.sendError(ctx, 400, "Request cannot be processed. Required options not provided")
		return
	}

	if err := json.Unmarshal([]byte(optionsField), &options); err != nil {
		server.sendError(ctx, 400, err.Error())
		return
	}

	file, err := ctx.FormFile("file")
	if err != nil {
		server.sendError(ctx, 400, err.Error())
		return
	}

	fpath := fmt.Sprintf("%s/%s", server.inputPath, file.Filename)
	ctx.SaveUploadedFile(file, fpath)

	if err := server.converter.ConvertImage(fpath, options); err != nil {
		server.sendError(ctx, 500, err.Error())
		return
	}

	// TODO: Remove input file
	ctx.String(200, "OK")
}
