package middleware

import (
	"bytes"
	"encoding/json"
	"hara/internal/convert"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

func createFormRequest(formField, formValue string) *http.Request {
	form := url.Values{}
	form.Add(formField, formValue)

	req := httptest.NewRequest("POST", "/", strings.NewReader(form.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	return req
}

func createFileFormRequest(formField, formValue string) *http.Request {
	buf := new(bytes.Buffer)

	mw := multipart.NewWriter(buf)
	mw.CreateFormFile(formField, formValue)
	defer mw.Close()

	req := httptest.NewRequest("POST", "/", buf)
	req.Header.Add("Content-Type", mw.FormDataContentType())

	return req
}

func TestOptionsFieldProvided(t *testing.T) {
	testCases := []struct {
		request *http.Request
		correct bool
		name    string
	}{
		{
			createFormRequest("options", "{value}"),
			true,
			"CorrectFieldName",
		},
		{
			createFormRequest("not_options", "{value}"),
			false,
			"IncorrectFieldName",
		},
	}

	for _, c := range testCases {
		t.Run(c.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(rec)
			ctx.Request = c.request

			OptionsFieldProvided(ctx)

			if c.correct && rec.Code != 200 {
				t.Fatalf("%s: want 200, got %d, body %s", c.name, rec.Code, rec.Body.String())
			}

			if !c.correct && rec.Code == 200 {
				t.Fatalf("%s: want != 200, got %d, body %s", c.name, rec.Code, rec.Body.String())
			}
		})
	}
}

func TestFileFieldProvided(t *testing.T) {
	testCases := []struct {
		request *http.Request
		correct bool
		name    string
	}{
		{
			createFileFormRequest("file", "1.mp4"),
			true,
			"CorrectFileField",
		},
		{
			createFileFormRequest("not_file", "1.mp4"),
			false,
			"IncorrectFileField",
		},
	}

	for _, c := range testCases {
		t.Run(c.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(rec)
			ctx.Request = c.request

			FileFieldProvided(ctx)

			if c.correct && rec.Code != 200 {
				t.Fatalf("%s: want 200, got %d, body %s", c.name, rec.Code, rec.Body.String())
			}

			if !c.correct && rec.Code == 200 {
				t.Fatalf("%s: want != 200, got %d, body %s", c.name, rec.Code, rec.Body.String())
			}
		})
	}
}

func TestValidateVideoOptionsJson(t *testing.T) {
	var options convert.ConversionVideoOptions
	optionsJson, _ := json.Marshal(options)

	testCases := []struct {
		request *http.Request
		correct bool
		name    string
	}{
		{
			createFormRequest("options", string(optionsJson)),
			true,
			"CorrectJson",
		},
		{
			createFormRequest("options", "not_json"),
			false,
			"IncorrectJson",
		},
	}

	for _, c := range testCases {
		t.Run(c.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(rec)
			ctx.Request = c.request

			ValidateVideoOptionsJson(ctx)

			if c.correct && rec.Code != 200 {
				t.Fatalf("%s: want 200, got %d, body %s", c.name, rec.Code, rec.Body.String())
			}

			if !c.correct && rec.Code == 200 {
				t.Fatalf("%s: want != 200, got %d, body %s", c.name, rec.Code, rec.Body.String())
			}

			got, ok := ctx.Get("videoOptions")
			if !ok {
				t.Fatalf("%s: videoOptions not set", c.name)
			}

			if got.(convert.ConversionVideoOptions) != options {
				t.Fatalf("%s: videoOptions not equal options", c.name)
			}
		})
	}
}

func TestValidateImageOptionsJson(t *testing.T) {
	var options convert.ConversionImageOptions
	optionsJson, _ := json.Marshal(options)

	testCases := []struct {
		request *http.Request
		correct bool
		name    string
	}{
		{
			createFormRequest("options", string(optionsJson)),
			true,
			"CorrectJson",
		},
		{
			createFormRequest("options", "not_json"),
			false,
			"IncorrectJson",
		},
	}

	for _, c := range testCases {
		t.Run(c.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(rec)
			ctx.Request = c.request

			ValidateImageOptionsJson(ctx)

			if c.correct && rec.Code != 200 {
				t.Fatalf("%s: want 200, got %d, body %s", c.name, rec.Code, rec.Body.String())
			}

			if !c.correct && rec.Code == 200 {
				t.Fatalf("%s: want != 200, got %d, body %s", c.name, rec.Code, rec.Body.String())
			}

			got, ok := ctx.Get("imageOptions")
			if !ok {
				t.Fatalf("%s: imageOptions not set", c.name)
			}

			if got.(convert.ConversionImageOptions) != options {
				t.Fatalf("%s: imageOptions not equal options", c.name)
			}
		})
	}
}

func TestSupportedVideoFileFormat(t *testing.T) {
	correctVideoFormat := convert.ConversionVideoOptions{Name: "1.mp4"}
	incorrectVideoFormat := convert.ConversionVideoOptions{Name: "1.exe"}

	testCases := []struct {
		correct bool
		name    string
	}{
		{
			true,
			"CorrectVideoFormat",
		},
		{
			false,
			"IncorrectVideoFormat",
		},
	}

	for _, c := range testCases {
		t.Run(c.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(rec)

			if c.correct {
				ctx.Set("videoOptions", correctVideoFormat)
				ctx.Set("file", &multipart.FileHeader{Filename: correctVideoFormat.Name})
			} else {
				ctx.Set("videoOptions", incorrectVideoFormat)
				ctx.Set("file", &multipart.FileHeader{Filename: incorrectVideoFormat.Name})
			}

			SupportedVideoFileFormat(ctx)

			if c.correct && rec.Code != 200 {
				t.Fatalf("%s: want 200, got %d, body %s", c.name, rec.Code, rec.Body.String())
			}

			if !c.correct && rec.Code == 200 {
				t.Fatalf("%s: want != 200, got %d, body %s", c.name, rec.Code, rec.Body.String())
			}
		})
	}
}

func TestSupportedImageFileFormat(t *testing.T) {
	correctImageFormat := convert.ConversionImageOptions{Name: "1.jpg"}
	incorrectImageFormat := convert.ConversionImageOptions{Name: "1.exe"}

	testCases := []struct {
		correct bool
		name    string
	}{
		{
			true,
			"CorrectImageFormat",
		},
		{
			false,
			"IncorrectImageFormat",
		},
	}

	for _, c := range testCases {
		t.Run(c.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(rec)

			if c.correct {
				ctx.Set("imageOptions", correctImageFormat)
				ctx.Set("file", &multipart.FileHeader{Filename: correctImageFormat.Name})
			} else {
				ctx.Set("imageOptions", incorrectImageFormat)
				ctx.Set("file", &multipart.FileHeader{Filename: incorrectImageFormat.Name})
			}

			SupportedImageFileFormat(ctx)

			if c.correct && rec.Code != 200 {
				t.Fatalf("%s: want 200, got %d, body %s", c.name, rec.Code, rec.Body.String())
			}

			if !c.correct && rec.Code == 200 {
				t.Fatalf("%s: want != 200, got %d, body %s", c.name, rec.Code, rec.Body.String())
			}
		})
	}
}
