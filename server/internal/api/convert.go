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
	opath := fmt.Sprintf("%s/%s", server.outputPath, options.Output.Name)
	ctx.SaveUploadedFile(file, fpath)

	if err := server.converter.ConvertVideo(fpath, opath, options); err != nil {
		server.sendError(ctx, 500, err.Error())
		return
	}

	ctx.String(200, "OK")
}

func (sever *Server) convertImage(ctx *gin.Context) {
	ctx.String(200, "Not implemented yet /api/convert/image")
}
