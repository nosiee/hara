package api

import "github.com/gin-gonic/gin"

func (server *Server) convertVideo(ctx *gin.Context) {
	ctx.String(200, "Not implemented yet /api/convert/video")
}

func (sever *Server) convertImage(ctx *gin.Context) {
	ctx.String(200, "Not implemented yet /api/convert/image")
}
