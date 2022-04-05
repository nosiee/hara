package middleware

import (
	"fmt"
	"hara/internal/convert"
	"mime/multipart"
	"net/url"
	"regexp"

	"github.com/gin-gonic/gin"
)

const (
	hourInSeconds  = 3600
	monthInSeconds = 2592000
)

func ConversionOptionsProvided(ctx *gin.Context) {
	if u, err := url.ParseQuery(ctx.Request.URL.RawQuery); err != nil || len(u) == 0 {
		ctx.AbortWithStatusJSON(400, gin.H{
			"error": "Request cannot be processed. Conversion options are not provided",
		})
	}
}

func ExtractConversionOptions(ctx *gin.Context) {
	values, _ := url.ParseQuery(ctx.Request.URL.RawQuery)
	options := convert.UrlQueryToOptions(values)

	ctx.Set("options", options)
}

func FileFieldProvided(ctx *gin.Context) {
	f, err := ctx.FormFile("file")
	if err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{
			"error": fmt.Sprintf("Request cannot be processed. %s", err),
		})
	}

	ctx.Set("file", f)
}

func SupportedVideoFileFormat(ctx *gin.Context) {
	var supportedFilePattern = "(3g2|3gp|3gpp|avi|cavs|dv|dvr|flv|m2ts|m4v|mkv|mod|mov|mp4|mpeg|mpg|mts|mxf|ogg|rm|webm|wmv)$"
	supportedFileRegexp, _ := regexp.Compile(supportedFilePattern)
	vopt, _ := ctx.Get("options")
	f, _ := ctx.Get("file")

	if !supportedFileRegexp.MatchString(f.(*multipart.FileHeader).Filename) {
		ctx.AbortWithStatusJSON(400, gin.H{
			"error": "Unsupported file format",
		})
	}

	if !supportedFileRegexp.MatchString(vopt.(convert.ConversionOptions).Extension) {
		ctx.AbortWithStatusJSON(400, gin.H{
			"error": "Unsupported file format",
		})
	}
}

func SupportedImageFileFormat(ctx *gin.Context) {
	var supportedFilePattern = "(jpg|jpeg|png|webp|gif|ico|bmp)$"
	supportedFileRegexp, _ := regexp.Compile(supportedFilePattern)
	iopt, _ := ctx.Get("options")
	f, _ := ctx.Get("file")

	if !supportedFileRegexp.MatchString(f.(*multipart.FileHeader).Filename) {
		ctx.AbortWithStatusJSON(400, gin.H{
			"error": "Unsupported file format",
		})
	}

	if !supportedFileRegexp.MatchString(iopt.(convert.ConversionOptions).Extension) {
		ctx.AbortWithStatusJSON(400, gin.H{
			"error": "Unsupported file format",
		})
	}
}

func ValidateLifetime(ctx *gin.Context) {
	optField, _ := ctx.Get("options")
	options := optField.(convert.ConversionOptions)

	if options.Lifetime < hourInSeconds || options.Lifetime > monthInSeconds {
		options.Lifetime = hourInSeconds

		ctx.Set("options", options)
	}
}
