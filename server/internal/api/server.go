package api

import "github.com/gin-gonic/gin"

type Server struct {
	inputPath  string
	outputPath string
	gin        *gin.Engine
}

func NewServer(inputPath string, outputPath string) *Server {
	return &Server{
		inputPath,
		outputPath,
		gin.New(),
	}
}

func (server *Server) Run(endpoint string) {
	server.gin.POST("/api/convert/video", server.convertVideo)
	server.gin.POST("/api/convert/image", server.convertImage)

	server.gin.Run(endpoint)
}
