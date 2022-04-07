package controllers

import (
	"fmt"
	"hara/internal/config"
	"io"
	"net/http"

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

func GenerateAPIUrl(ctx *gin.Context, apiPrefix, key string) string {
	scheme := "http://"
	host := ctx.Request.Host

	if ctx.Request.Proto == "HTTP/2" {
		scheme = "https://"
	}

	return fmt.Sprintf("%s%s/api/convert/%s?key=%s", scheme, host, apiPrefix, key)
}

func GetFileContentType(reader io.Reader) (string, error) {
	buffer := make([]byte, 512)

	if _, err := reader.Read(buffer); err != nil {
		return "", err
	}

	contentType := http.DetectContentType(buffer)
	return contentType, nil
}

func GenerateJWT(uuid string, key string) (string, error) {
	if len(key) != HS512KeySize {
		return "", jwt.ErrInvalidKey
	}

	payload := jwt.MapClaims{}
	payload["uuid"] = uuid
	payload["exp"] = JWTExp

	header := jwt.NewWithClaims(jwt.SigningMethodHS512, payload)

	return header.SignedString([]byte(key))
}

func ExtractUserIDFromJWT(t string) (string, error) {
	token, err := jwt.Parse(t, func(t *jwt.Token) (interface{}, error) {
		return []byte(config.Values.HS512Key), nil
	})

	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return "", token.Claims.Valid()
	}

	return claims["uuid"].(string), nil
}
