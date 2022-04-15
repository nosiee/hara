package controllers

import (
	"fmt"
	"hara/internal/config"
	"hara/internal/testhelpers"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func init() {
	gin.SetMode(gin.TestMode)
}

func TestGenerateFileUrl(t *testing.T) {
	testCases := []testhelpers.UrlCase{
		testhelpers.CreateUrlCase("HTTP/1.1", "http://", "localhost:8080", "i", "test_image.jpg", "ImageUrl"),
		testhelpers.CreateUrlCase("HTTP/2", "https://", "localhost:8080", "v", "test_video.webm", "VideoUrl"),
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			fileUrl := fmt.Sprintf("%s%s/api/%s/%s", tc.Scheme, tc.Host, tc.ApiPrefix, tc.Value)

			ctx, _ := testhelpers.CreateContext(httptest.NewRequest("GET", fileUrl, nil), nil)
			ctx.Request.Proto = tc.Proto

			url := GenerateFileUrl(ctx, tc.ApiPrefix, tc.Value)
			if fileUrl != url {
				t.Fatalf("%s want %s, got %s", tc.Name, fileUrl, url)
			}
		})
	}
}

func TestGenerateAPIUrl(t *testing.T) {
	testCases := []testhelpers.UrlCase{
		testhelpers.CreateUrlCase("HTTP/1.1", "http://", "localhost:8080", "image", uuid.NewString(), "ImageUrl"),
		testhelpers.CreateUrlCase("HTTP/2", "https://", "localhost:8080", "video", uuid.NewString(), "VideoUrl"),
	}

	for _, tc := range testCases {
		apiUrl := fmt.Sprintf("%s%s/api/convert/%s?key=%s", tc.Scheme, tc.Host, tc.ApiPrefix, tc.Value)

		ctx, _ := testhelpers.CreateContext(httptest.NewRequest("GET", apiUrl, nil), nil)
		ctx.Request.Proto = tc.Proto

		url := GenerateAPIUrl(ctx, tc.ApiPrefix, tc.Value)
		if apiUrl != url {
			t.Fatalf("%s want %s, got %s", tc.Name, apiUrl, url)
		}
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

func TestGenerateJWT(t *testing.T) {
	config.LoadFromFile("../../testdata/configs/config_test_correct.toml")

	testCases := []struct {
		key     string
		correct bool
		name    string
	}{
		{
			config.Values.HS512Key,
			true,
			"CorrectJWTKey",
		},
		{
			"not_valid_key",
			false,
			"IncorrectJWTKey",
		},
	}

	for _, c := range testCases {
		t.Run(c.name, func(t *testing.T) {
			_, err := GenerateJWT("someid", c.key, time.Now().Add(1*365*24*time.Hour).Unix())

			if c.correct && err != nil {
				t.Fatalf("%s want err == nil, got %v", c.name, err)
			}

			if !c.correct && err == nil {
				t.Fatalf("%s want err != nil, got nil", c.name)
			}
		})
	}
}

func TestExtractUserIDFromJWT(t *testing.T) {
	config.LoadFromFile("../../testdata/configs/config_test_correct.toml")

	id := uuid.NewString()
	token, err := GenerateJWT(id, config.Values.HS512Key, time.Now().Add(1*365*24*time.Hour).Unix())
	if err != nil {
		t.Fatal(err)
	}

	expToken, err := GenerateJWT(id, config.Values.HS512Key, 120)
	if err != nil {
		t.Fatal(err)
	}

	testCases := []struct {
		token   string
		uuid    string
		correct bool
		name    string
	}{
		{
			token,
			id,
			true,
			"CorrectToken",
		},
		{
			"definitelynotatoken",
			id,
			false,
			"IncorrectToken",
		},
		{
			expToken,
			id,
			false,
			"ExpiredToken",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			id, err := ExtractUserIDFromJWT(tc.token)

			if tc.correct && err != nil {
				t.Fatalf("%s got %v", tc.name, err)
			}

			if !tc.correct && err == nil {
				t.Fatalf("%s want err != nil", tc.name)
			}

			if tc.correct && (tc.uuid != id) {
				t.Fatalf("%s want %s, got %s", tc.name, tc.uuid, id)
			}
		})
	}
}
