package middleware

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

type contextCase struct {
	context  *gin.Context
	recorder *httptest.ResponseRecorder
	correct  bool
	name     string
}

func createContextCase(context *gin.Context, recorder *httptest.ResponseRecorder, correct bool, name string) contextCase {
	return contextCase{
		context,
		recorder,
		correct,
		name,
	}
}

func (c contextCase) checkCase(t *testing.T) {
	if c.correct && c.recorder.Code != 200 {
		t.Fatalf("%s want status code 200, got %d", c.name, c.recorder.Code)
	}

	if !c.correct && c.recorder.Code == 200 {
		t.Fatalf("%s want status code != 200, got 200", c.name)
	}
}

func createContextWithRequest(request *http.Request) (*gin.Context, *httptest.ResponseRecorder) {
	recorder := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(recorder)
	context.Request = request

	return context, recorder
}

func createFormRequest(form map[string]string) *http.Request {
	data := url.Values{}

	for k, v := range form {
		data.Set(k, v)
	}

	req := httptest.NewRequest("POST", "/", strings.NewReader(data.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	return req
}

func createCookieRequest(cookieKey, cookieValue string) *http.Request {
	req := httptest.NewRequest("POST", "/", nil)
	req.Header.Set("Cookie", fmt.Sprintf("%s=%s", cookieKey, cookieValue))
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
