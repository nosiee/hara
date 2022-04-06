package middleware

import (
	"bytes"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
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

func createFileFormRequest(formField, formValue string) *http.Request {
	buf := new(bytes.Buffer)

	mw := multipart.NewWriter(buf)
	mw.CreateFormFile(formField, formValue)
	defer mw.Close()

	req := httptest.NewRequest("POST", "/", buf)
	req.Header.Add("Content-Type", mw.FormDataContentType())

	return req
}
