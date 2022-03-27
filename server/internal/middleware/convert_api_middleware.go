package middleware

import (
	"encoding/json"
	"fmt"
	"hara/internal/convert"
	"regexp"

	"github.com/gin-gonic/gin"
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

	ctx.Set("videoOptions", vopt)
}

func ValidateImageOptionsJson(ctx *gin.Context) {
	var iopt convert.ConversionImageOptions

	optionsField, _ := ctx.GetPostForm("options")
	if err := json.Unmarshal([]byte(optionsField), &iopt); err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{
			"error": fmt.Sprintf("Request cannot be processed. %s", err),
		})
	}

	ctx.Set("imageOptions", iopt)
}

func SupportedVideoFileFormat(ctx *gin.Context) {
	// TODO: Also check input file format!!!
	var supportedFilePattern = "^\\w+.(3g2|3gp|3gpp|avi|cavs|dv|dvr|flv|m2ts|m4v|mkv|mod|mov|mp4|mpeg|mpg|mts|mxf|ogg|rm|webm|wmv)$"
	supportedFileRegexp, _ := regexp.Compile(supportedFilePattern)
	vopt, _ := ctx.Get("videoOptions")

	if !supportedFileRegexp.MatchString(vopt.(convert.ConversionVideoOptions).Name) {
		ctx.AbortWithStatusJSON(400, gin.H{
			"error": "Unsupported file format",
		})
	}
}

func SupportedImageFileFormat(ctx *gin.Context) {
	// TODO: Also check input file format!!!
	// TODO: Add more extensions
	var supportedFilePattern = "^\\w+.(jpg|jpeg|png|webp|gyf)$"
	supportedFileRegexp, _ := regexp.Compile(supportedFilePattern)
	iopt, _ := ctx.Get("imageOptions")

	if !supportedFileRegexp.MatchString(iopt.(convert.ConversionImageOptions).Name) {
		ctx.AbortWithStatusJSON(400, gin.H{
			"error": "Unsupported file format",
		})
	}
}
