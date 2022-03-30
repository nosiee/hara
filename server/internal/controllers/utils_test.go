package controllers

import (
	"fmt"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestGenerateFileUrl(t *testing.T) {
	testCases := []struct {
		proto     string
		scheme    string
		apiPrefix string
		fpath     string
		name      string
	}{
		{
			"HTTP/1.1",
			"http://",
			"i",
			"test_image.jpg",
			"CorrectImageUrl",
		},
		{
			"HTTP/2",
			"https://",
			"v",
			"test_video.webm",
			"CorrectVideoUrl",
		},
	}

	host := "localhost:8080"

	for _, c := range testCases {
		t.Run(c.name, func(t *testing.T) {
			fileUrl := fmt.Sprintf("%s%s/api/%s/%s", c.scheme, host, c.apiPrefix, c.fpath)
			ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

			ctx.Request = httptest.NewRequest("GET", fileUrl, nil)
			ctx.Request.Proto = c.proto

			url := GenerateFileUrl(ctx, c.apiPrefix, c.fpath)

			if fileUrl != url {
				t.Fatalf("%s want %s, got %s", c.name, fileUrl, url)
			}
		})
	}
}

func TestGetFileContentType(t *testing.T) {
	testCases := []struct {
		fpath       string
		contentType string
		correct     bool
		name        string
	}{
		{
			"../../testdata/images/pepe_jpg.jpg",
			"image/jpeg",
			true,
			"CorrectJpg",
		},
		{
			"../../testdata/images/pepe_gif.gif",
			"image/gif",
			true,
			"CorrectGif",
		},
		{
			"../../testdata/images/no_such_file.test",
			"",
			false,
			"IncorrectFile",
		},
	}

	for _, c := range testCases {
		t.Run(c.name, func(t *testing.T) {
			f, _ := os.Open(c.fpath)

			contentType, err := GetFileContentType(f)
			if c.correct && err != nil {
				t.Fatalf("%s want err == nil, got %v", c.name, err)
			}

			if contentType != c.contentType {
				t.Fatalf("%s want %s, got %s", c.name, c.contentType, contentType)
			}
		})
	}
}
