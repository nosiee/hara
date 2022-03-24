package api

import "github.com/gin-gonic/gin"

type Server struct {
	inputFolder  string
	outputFolder string
	gin          *gin.Engine
}

func NewServer(inputFolder string, outputFolder string) *Server {
	return &Server{
		inputFolder,
		outputFolder,
		gin.New(),
	}
}

func (server *Server) Run(endpoint string) {
	server.gin.POST("/api/convert/video", server.convertVideo)
	server.gin.POST("/api/convert/image", server.convertImage)

	server.gin.Run(endpoint)
}
