package controllers

import (
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GenerateFileUrl(ctx *gin.Context, apiPrefix, ofile string) string {
	proto := "http://"

	if ctx.Request.Proto == "HTTP/2" {
		proto = "https://"
	}

	return fmt.Sprintf("%s%s/api/%s/%s", proto, ctx.Request.Host, apiPrefix, ofile)
}

func GetFileContentType(reader io.Reader) (string, error) {
	buffer := make([]byte, 512)

	if _, err := reader.Read(buffer); err != nil {
		return "", err
	}

	contentType := http.DetectContentType(buffer)
	return contentType, nil
}
