package api

import (
	"encoding/json"
	"fmt"
	"hara/internal/convert"

	"github.com/gin-gonic/gin"
)

func (server *Server) convertVideo(ctx *gin.Context) {
	var options convert.ConversionVideoOptions

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

	out, err := convert.ConvertVideo(fpath, server.outputPath, &options)
	if err != nil {
		server.sendError(ctx, 500, err.Error())
		return
	}

	ctx.String(200, out)
}

func (sever *Server) convertImage(ctx *gin.Context) {
	ctx.String(200, "Not implemented yet /api/convert/image")
}
