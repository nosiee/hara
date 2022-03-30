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

type requestCase struct {
	request *http.Request
	correct bool
	name    string
}

type contextCase struct {
	context any
	correct bool
	name    string
}

func init() {
	gin.SetMode(gin.TestMode)
}

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

func testFieldProvided(t *testing.T, c requestCase, f func(*gin.Context)) *gin.Context {
	rec := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(rec)
	ctx.Request = c.request

	f(ctx)

	if c.correct && rec.Code != 200 {
		t.Fatalf("%s: want 200, got %d, body %s", c.name, rec.Code, rec.Body.String())
	}

	if !c.correct && rec.Code == 200 {
		t.Fatalf("%s: want != 200, got %d, body %s", c.name, rec.Code, rec.Body.String())
	}

	return ctx
}

func testSupportedFileFormat(t *testing.T, c contextCase, f func(*gin.Context)) {
	rec := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(rec)

	switch c.context.(type) {
	case convert.ConversionImageOptions:
		ctx.Set("options", c.context.(convert.ConversionImageOptions))
		ctx.Set("file", &multipart.FileHeader{Filename: c.context.(convert.ConversionImageOptions).Extension})
	case convert.ConversionVideoOptions:
		ctx.Set("options", c.context.(convert.ConversionVideoOptions))
		ctx.Set("file", &multipart.FileHeader{Filename: c.context.(convert.ConversionVideoOptions).Extension})
	}

	f(ctx)

	if c.correct && rec.Code != 200 {
		t.Fatalf("%s: want 200, got %d, body %s", c.name, rec.Code, rec.Body.String())
	}

	if !c.correct && rec.Code == 200 {
		t.Fatalf("%s: want != 200, got %d, body %s", c.name, rec.Code, rec.Body.String())
	}
}

func TestOptionsFieldProvided(t *testing.T) {
	testCases := []requestCase{
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
			testFieldProvided(t, c, OptionsFieldProvided)
		})
	}
}

func TestFileFieldProvided(t *testing.T) {
	testCases := []requestCase{
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
			testFieldProvided(t, c, FileFieldProvided)
		})
	}
}

func TestValidateVideoOptionsJson(t *testing.T) {
	var options convert.ConversionVideoOptions
	optionsJson, _ := json.Marshal(options)

	testCases := []requestCase{
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
			ctx := testFieldProvided(t, c, ValidateVideoOptionsJson)

			got, ok := ctx.Get("options")
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

	testCases := []requestCase{
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
			ctx := testFieldProvided(t, c, ValidateImageOptionsJson)

			got, ok := ctx.Get("options")
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
	correctVideoFormat := convert.ConversionVideoOptions{Extension: "mp4"}
	incorrectVideoFormat := convert.ConversionVideoOptions{Extension: "exe"}

	testCases := []contextCase{
		{
			correctVideoFormat,
			true,
			"CorrectVideoFormat",
		},
		{
			incorrectVideoFormat,
			false,
			"IncorrectVideoFormat",
		},
	}

	for _, c := range testCases {
		t.Run(c.name, func(t *testing.T) {
			testSupportedFileFormat(t, c, SupportedVideoFileFormat)
		})
	}
}

func TestSupportedImageFileFormat(t *testing.T) {
	correctImageFormat := convert.ConversionImageOptions{Extension: "jpg"}
	incorrectImageFormat := convert.ConversionImageOptions{Extension: "exe"}

	testCases := []contextCase{
		{
			correctImageFormat,
			true,
			"CorrectImageFormat",
		},
		{
			incorrectImageFormat,
			false,
			"IncorrectImageFormat",
		},
	}

	for _, c := range testCases {
		t.Run(c.name, func(t *testing.T) {
			testSupportedFileFormat(t, c, SupportedImageFileFormat)
		})
	}
}

func TestValidateLifetime(t *testing.T) {
	correctLifetime := convert.ConversionImageOptions{Lifetime: 4200}
	moreThanMonthLifetimeImage := convert.ConversionImageOptions{Lifetime: 934579384759374985}
	lessThanHourLifetimeImage := convert.ConversionImageOptions{Lifetime: 120}

	lessThanHourLifetimeVideo := convert.ConversionVideoOptions{Lifetime: 120}
	moreThanMonthLifetimeVideo := convert.ConversionVideoOptions{Lifetime: 934579384759374985}

	testCases := []struct {
		options  any
		correct  bool
		lifetime uint
		name     string
	}{
		{
			correctLifetime,
			true,
			4200,
			"CorrectLifetime",
		},
		{
			moreThanMonthLifetimeImage,
			false,
			934579384759374985,
			"MoreThanMonthImage",
		},
		{
			lessThanHourLifetimeImage,
			false,
			120,
			"LessThanHourImage",
		},
		{
			moreThanMonthLifetimeVideo,
			false,
			934579384759374985,
			"MoreThanMonthVideo",
		},
		{
			lessThanHourLifetimeVideo,
			false,
			120,
			"LessThanHourVideo",
		},
	}

	for _, c := range testCases {
		t.Run(c.name, func(t *testing.T) {
			ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
			ctx.Set("options", c.options)

			ValidateLifetime(ctx)

			var context any
			var ok bool
			var checkedLifetime uint

			if context, ok = ctx.Get("options"); !ok {
				t.Fatalf("%s no context after ValidateLifetime", c.name)
			}

			switch context.(type) {
			case convert.ConversionImageOptions:
				checkedLifetime = context.(convert.ConversionImageOptions).Lifetime
			case convert.ConversionVideoOptions:
				checkedLifetime = context.(convert.ConversionVideoOptions).Lifetime
			}

			if c.correct && c.lifetime != checkedLifetime {
				t.Fatalf("%s want %d, got %d", c.name, c.lifetime, checkedLifetime)
			}

			if !c.correct {
				if c.lifetime < hourInSeconds && checkedLifetime != hourInSeconds {
					t.Fatalf("%s want %d, got %d", c.name, hourInSeconds, checkedLifetime)
				}

				if c.lifetime > monthInSeconds && checkedLifetime != monthInSeconds {
					t.Fatalf("%s want %d, got %d", c.name, monthInSeconds, checkedLifetime)
				}
			}
		})
	}
}
