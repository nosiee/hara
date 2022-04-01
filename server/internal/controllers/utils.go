package controllers

import (
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GenerateFileUrl(ctx *gin.Context, apiPrefix, ofile string) string {
	scheme := "http://"
	host := ctx.Request.Host

	if ctx.Request.Proto == "HTTP/2" {
		scheme = "https://"
	}

	return fmt.Sprintf("%s%s/api/%s/%s", scheme, host, apiPrefix, ofile)
}

func GetFileContentType(reader io.Reader) (string, error) {
	buffer := make([]byte, 512)

	if _, err := reader.Read(buffer); err != nil {
		return "", err
	}

	contentType := http.DetectContentType(buffer)
	return contentType, nil
}
