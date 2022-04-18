package api

import (
	"hara/internal/controllers"
	"hara/internal/middleware"

	"github.com/gin-gonic/gin"
)

type Server struct {
	engine      *gin.Engine
	controllers *controllers.Controllers
}

func NewServer(controllers *controllers.Controllers) *Server {
	return &Server{
		gin.New(),
		controllers,
	}
}

func (serv Server) RunServer(endpoint string) {
	serv.engine.POST("/api/auth/signup", middleware.SignUpFormProvided, middleware.SignUpFormValidate, serv.controllers.SignUp)
	serv.engine.POST("/api/auth/signin", middleware.SignInFormProvided, middleware.SignInFormValidate, serv.controllers.SignIn)

	// TODO: /api/key/reset
	serv.engine.GET("/api/key/get", middleware.IsAuthorized, serv.controllers.GetApiKey)

	serv.engine.POST("/api/convert/image",
		middleware.ApiKeyProvided, middleware.ApiKeyValidate(serv.controllers.ApikeyRepository),
		middleware.ApiKeyQuota(serv.controllers.ApikeyRepository), middleware.ExtractConversionOptions,
		middleware.ImageFileFieldProvided, middleware.SupportedImageFileExtension,
		middleware.ValidateLifetime, serv.controllers.ImageController)

	serv.engine.POST("/api/convert/video",
		middleware.ApiKeyProvided, middleware.ApiKeyValidate(serv.controllers.ApikeyRepository),
		middleware.ApiKeyQuota(serv.controllers.ApikeyRepository), middleware.ExtractConversionOptions,
		middleware.VideoFileFieldProvided, middleware.SupportedVideoFileExtension,
		middleware.ValidateLifetime, serv.controllers.VideoController)

	serv.engine.GET("/api/i/:filename", serv.controllers.GetImage)
	serv.engine.GET("/api/v/:filename", serv.controllers.GetVideo)

	serv.engine.Run(endpoint)
}
