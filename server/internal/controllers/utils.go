package controllers

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
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

func GenerateRandomUUID() string {
	u := make([]byte, 32)
	_, _ = rand.Read(u)

	u[8] = (u[8] | 0x80) & 0xBF
	u[6] = (u[6] | 0x40) & 0x4F

	return hex.EncodeToString(u)
}

func GenerateJWT(uuid, key string) (string, error) {
	payload := jwt.MapClaims{}
	payload["uuid"] = uuid
	// TODO: make time constant and use it in signup and signin functions
	payload["exp"] = time.Now().Add(1 * 365 * 24 * time.Hour).Unix()

	header := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	return header.SignedString([]byte(key))
}
