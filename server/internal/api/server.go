package api

import (
	"hara/internal/controllers"
	"hara/internal/middleware"

	"github.com/gin-gonic/gin"
)

var (
	gen = gin.New()
)

func RunServer(endpoint string) {
	gen.POST("/api/convert/video", middleware.OptionsFieldProvided, middleware.FileFieldProvided, middleware.ValidateVideoOptionsJson, middleware.SupportedVideoFileFormat, controllers.VideoController)
	gen.POST("/api/convert/image", middleware.OptionsFieldProvided, middleware.FileFieldProvided, middleware.ValidateImageOptionsJson, middleware.SupportedImageFileFormat, controllers.ImageController)

	gen.Run(endpoint)
}
