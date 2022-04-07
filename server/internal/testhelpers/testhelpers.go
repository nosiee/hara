package testhelpers

import (
	"bytes"
	"fmt"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

func init() {
	gin.SetMode(gin.TestMode)
}

type ContextCase struct {
	Context  *gin.Context
	Recorder *httptest.ResponseRecorder
	Correct  bool
	Name     string
}

type UrlCase struct {
	Proto     string
	Scheme    string
	Host      string
	ApiPrefix string
	Value     string
	Name      string
}

func CreateContextCase(context *gin.Context, recorder *httptest.ResponseRecorder, correct bool, name string) ContextCase {
	return ContextCase{
		context,
		recorder,
		correct,
		name,
	}
}

func CreateUrlCase(proto, scheme, host, apiPrefix, value, name string) UrlCase {
	return UrlCase{
		proto,
		scheme,
		host,
		apiPrefix,
		value,
		name,
	}
}

func (c ContextCase) CheckCase(t *testing.T) {
	if c.Correct && c.Recorder.Code != 200 {
		t.Fatalf("%s want status code 200, got %d", c.Name, c.Recorder.Code)
	}

	if !c.Correct && c.Recorder.Code == 200 {
		t.Fatalf("%s want status code != 200, got 200", c.Name)
	}
}

func CreateContextWithRequest(request *http.Request) (*gin.Context, *httptest.ResponseRecorder) {
	recorder := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(recorder)
	context.Request = request

	return context, recorder
}

func CreateFormRequest(form map[string]string) *http.Request {
	data := url.Values{}

	for k, v := range form {
		data.Set(k, v)
	}

	req := httptest.NewRequest("POST", "/", strings.NewReader(data.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	return req
}

func CreateCookieRequest(cookieKey, cookieValue string) *http.Request {
	req := httptest.NewRequest("POST", "/", nil)
	req.Header.Set("Cookie", fmt.Sprintf("%s=%s", cookieKey, cookieValue))
	return req
}

func CreateFileFormRequest(formField, formValue string) *http.Request {
	buf := new(bytes.Buffer)

	mw := multipart.NewWriter(buf)
	mw.CreateFormFile(formField, formValue)
	defer mw.Close()

	req := httptest.NewRequest("POST", "/", buf)
	req.Header.Add("Content-Type", mw.FormDataContentType())

	return req
}
