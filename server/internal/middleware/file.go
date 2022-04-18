package middleware

import (
	"hara/internal/convert"
	"net/url"

	"regexp"
	"strconv"

	"github.com/gin-gonic/gin"
)

const (
	hourInSeconds  = 3600
	monthInSeconds = 2592000

	supportedVideoExtensions = "(3g2|3gp|3gpp|avi|cavs|dv|dvr|flv|m2ts|m4v|mkv|mod|mov|mp4|mpeg|mpg|mts|mxf|ogg|rm|webm|wmv)$"
	supportedImageExtensions = "(jpg|jpeg|png|webp|gif|ico|bmp)$"
)

func ImageFileFieldProvided(ctx *gin.Context) {
	f, err := ctx.FormFile("file")
	if err != nil {
		errLogger.Println(err)

		ctx.AbortWithStatusJSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	supportedExt, _ := regexp.Compile(supportedImageExtensions)
	if !supportedExt.MatchString(f.Filename) {
		ctx.AbortWithStatusJSON(400, gin.H{
			"error": "Unsupported file format",
		})
		return
	}

	ctx.Set("file", f)
}

func VideoFileFieldProvided(ctx *gin.Context) {
	f, err := ctx.FormFile("file")
	if err != nil {
		errLogger.Println(err)

		ctx.AbortWithStatusJSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	supportedExt, _ := regexp.Compile(supportedVideoExtensions)
	if !supportedExt.MatchString(f.Filename) {
		ctx.AbortWithStatusJSON(400, gin.H{
			"error": "Unsupported file format",
		})
		return
	}

	ctx.Set("file", f)
}

func SupportedImageFileExtension(ctx *gin.Context) {
	values, _ := url.ParseQuery(ctx.Request.URL.RawQuery)
	supportedExt, _ := regexp.Compile(supportedImageExtensions)

	if !supportedExt.MatchString(values.Get("ext")) {
		ctx.AbortWithStatusJSON(400, gin.H{
			"error": "Unsupported file format",
		})
		return
	}
}

func SupportedVideoFileExtension(ctx *gin.Context) {
	values, _ := url.ParseQuery(ctx.Request.URL.RawQuery)
	supportedExt, _ := regexp.Compile(supportedVideoExtensions)

	if !supportedExt.MatchString(values.Get("ext")) {
		ctx.AbortWithStatusJSON(400, gin.H{
			"error": "Unsupported file format",
		})
		return
	}
}

func ValidateLifetime(ctx *gin.Context) {
	values, _ := url.ParseQuery(ctx.Request.URL.RawQuery)
	lifetime, err := strconv.ParseUint(values.Get("lifetime"), 10, 64)

	if err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{
			"error": "Lifetime parameter should be uint in seconds",
		})
		return
	}

	if lifetime < hourInSeconds || lifetime > monthInSeconds {
		ctx.AbortWithStatusJSON(400, gin.H{
			"error": "Lifetime should be more than an hour, and less than a month",
		})
	}
}

func ExtractConversionOptions(ctx *gin.Context) {
	values, _ := url.ParseQuery(ctx.Request.URL.RawQuery)
	options := convert.UrlQueryToOptions(values)

	ctx.Set("options", options)
}
