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
	gen.POST("/api/auth/signup", middleware.SignUpFormProvided, middleware.SignUpFormValidate, controllers.SignUp)
	gen.POST("/api/auth/signin", middleware.SignInFormProvided, middleware.SignInFormValidate, controllers.SignIn)

	gen.GET("/api/key/get", middleware.IsAuthorized, controllers.GetApiKey)

	gen.POST("/api/convert/image",
		middleware.ApiKeyProvided, middleware.ApiKeyValidate, middleware.ExtractConversionOptions,
		middleware.ImageFileFieldProvided, middleware.SupportedImageFileExtension,
		middleware.ValidateLifetime, controllers.ImageController)

	gen.POST("/api/convert/video",
		middleware.ApiKeyProvided, middleware.ApiKeyValidate, middleware.ExtractConversionOptions,
		middleware.VideoFileFieldProvided, middleware.SupportedVideoFileExtension,
		middleware.ValidateLifetime, controllers.VideoController)

	gen.GET("/api/i/:filename", controllers.GetImage)
	gen.GET("/api/v/:filename", controllers.GetVideo)

	gen.Run(endpoint)
}
