package middleware

import (
	"encoding/json"
	"fmt"
	"hara/internal/convert"
	"mime/multipart"
	"regexp"

	"github.com/gin-gonic/gin"
)

const (
	hourInSeconds  = 3600
	monthInSeconds = 2592000
)

func OptionsFieldProvided(ctx *gin.Context) {
	if _, ok := ctx.GetPostForm("options"); !ok {
		ctx.AbortWithStatusJSON(400, gin.H{
			"error": "Request cannot be processed. Options field not provided",
		})
	}
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

func ValidateVideoOptionsJson(ctx *gin.Context) {
	var vopt convert.ConversionVideoOptions

	optionsField, _ := ctx.GetPostForm("options")
	if err := json.Unmarshal([]byte(optionsField), &vopt); err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{
			"error": fmt.Sprintf("Request cannot be processed. %s", err),
		})
	}

	ctx.Set("options", vopt)
}

func ValidateImageOptionsJson(ctx *gin.Context) {
	var iopt convert.ConversionImageOptions

	optionsField, _ := ctx.GetPostForm("options")
	if err := json.Unmarshal([]byte(optionsField), &iopt); err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{
			"error": fmt.Sprintf("Request cannot be processed. %s", err),
		})
	}

	ctx.Set("options", iopt)
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

	if !supportedFileRegexp.MatchString(vopt.(convert.ConversionVideoOptions).Extension) {
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

	if !supportedFileRegexp.MatchString(iopt.(convert.ConversionImageOptions).Extension) {
		ctx.AbortWithStatusJSON(400, gin.H{
			"error": "Unsupported file format",
		})
	}
}

func ValidateLifetime(ctx *gin.Context) {
	options, _ := ctx.Get("options")

	switch options.(type) {
	case convert.ConversionImageOptions:
		iopt := options.(convert.ConversionImageOptions)

		if iopt.Lifetime < hourInSeconds {
			iopt.Lifetime = hourInSeconds
		} else if iopt.Lifetime > monthInSeconds {
			iopt.Lifetime = monthInSeconds
		}

		ctx.Set("options", iopt)

	case convert.ConversionVideoOptions:
		vopt := options.(convert.ConversionVideoOptions)

		if vopt.Lifetime < hourInSeconds {
			vopt.Lifetime = hourInSeconds
		} else if vopt.Lifetime > monthInSeconds {
			vopt.Lifetime = monthInSeconds
		}

		ctx.Set("options", vopt)
	}
}
