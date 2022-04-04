package middleware

import (
	"bytes"
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

func createCookieRequest(method, cookie string) *http.Request {
	req := httptest.NewRequest(method, "/", nil)
	req.Header.Set("Cookie", cookie)
	return req
}

func createFormRequest(f map[string]string, reqUrl string) *http.Request {
	form := url.Values{}

	for k, v := range f {
		form.Add(k, v)
	}

	req := httptest.NewRequest("POST", reqUrl, strings.NewReader(form.Encode()))
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

func testRequest(t *testing.T, c requestCase, f func(*gin.Context)) *gin.Context {
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
