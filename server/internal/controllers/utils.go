package controllers

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func GenerateFileUrl(ctx *gin.Context, apiPrefix, ofile string) string {
	proto := "http://"

	if ctx.Request.Proto == "HTTP/2" {
		proto = "https://"
	}

	return fmt.Sprintf("%s%s/api/%s/%s", proto, ctx.Request.Host, apiPrefix, ofile)
}

func GetFileContentType(f *os.File) (string, error) {
	buffer := make([]byte, 512)

	if _, err := f.Read(buffer); err != nil {
		return "", err
	}

	contentType := http.DetectContentType(buffer)
	return contentType, nil
}
