package api

import (
	"hara/internal/controllers"
	"hara/internal/middleware"

	"github.com/gin-gonic/gin"
)

type Server struct {
	gin       *gin.Engine
	inputPath string
}

func NewServer(inputPath string) Server {
	return Server{
		gin.New(),
		inputPath,
	}
}

func (server Server) Run(endpoint string) {
	server.gin.POST("/api/convert/video", middleware.OptionsFieldProvided, middleware.FileFieldProvided, middleware.ValidateVideoOptionsJson, middleware.SupportedVideoFileFormat, controllers.VideoController)
	server.gin.POST("/api/convert/image", middleware.OptionsFieldProvided, middleware.FileFieldProvided, middleware.ValidateImageOptionsJson, middleware.SupportedImageFileFormat, controllers.ImageController)

	server.gin.Run(endpoint)
}
