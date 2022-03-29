package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func GenerateFileUrl(ctx *gin.Context, apiPrefix, ofile string) string {
	proto := "http://"

	if ctx.Request.Proto == "HTTP/2" {
		proto = "https://"
	}

	return fmt.Sprintf("%s%s/api/%s/%s", proto, ctx.Request.Host, apiPrefix, ofile)
}
