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
	gen.POST("/api/auth/signup", middleware.AuthFormProvided, middleware.AuthFormValidate, controllers.SignUp)
	gen.POST("/api/auth/signin", middleware.AuthFormValidate, middleware.AuthFormValidate, controllers.SignIn)

	gen.GET("/api/key/get", middleware.IsAuthorized, controllers.GetApiKey)

	gen.POST("/api/convert/image",
		middleware.ApiKeyProvided, middleware.ApiKeyValidate,
		middleware.ConversionOptionsProvided, middleware.ExtractConversionOptions,
		middleware.FileFieldProvided, middleware.SupportedImageFileFormat,
		middleware.ValidateLifetime, controllers.ImageController)

	gen.POST("/api/convert/video",
		middleware.ApiKeyProvided, middleware.ApiKeyValidate,
		middleware.ConversionOptionsProvided, middleware.ExtractConversionOptions,
		middleware.FileFieldProvided, middleware.SupportedVideoFileFormat,
		middleware.ValidateLifetime, controllers.VideoController)

	gen.GET("/api/i/:filename", controllers.GetImage)
	gen.GET("/api/v/:filename", controllers.GetVideo)

	gen.Run(endpoint)
}
