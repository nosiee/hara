package api

import "github.com/gin-gonic/gin"

func (server *Server) sendError(ctx *gin.Context, httpcode int, message string) {
	ctx.JSON(httpcode, gin.H{
		"error": message,
	})
}
