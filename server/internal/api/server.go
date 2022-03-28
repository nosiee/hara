package api

import (
	"fmt"
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
	gen.GET("/api/i/:filename", testImage)
	gen.GET("/api/v/:filename", testVideo)

	gen.Run(endpoint)
}

func testImage(c *gin.Context) {
	fpath := fmt.Sprintf("output/images/%s", c.Param("filename"))

	c.Header("Content-Type", "image/jpg")
	c.File(fpath)
}

func testVideo(c *gin.Context) {
	fpath := fmt.Sprintf("output/videos/%s", c.Param("filename"))

	c.Header("Content-Type", "video/webm")
	c.File(fpath)
}
