package api

import (
	"hara/internal/convert"

	"github.com/gin-gonic/gin"
)

type Server struct {
	gin       *gin.Engine
	converter *convert.Converter
	inputPath string
}

func NewServer(inputPath string, converter *convert.Converter) *Server {
	return &Server{
		gin.New(),
		converter,
		inputPath,
	}
}

func (server *Server) Run(endpoint string) {
	server.gin.POST("/api/convert/video", server.convertVideo)
	server.gin.POST("/api/convert/image", server.convertImage)

	server.gin.Run(endpoint)
}
