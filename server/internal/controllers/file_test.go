package controllers

import (
	"hara/internal/config"
	"hara/internal/convert"
	mock "hara/internal/convert/mock"
	repository "hara/internal/repository/mock"
	"hara/internal/testhelpers"
	"mime/multipart"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestImageControllers(t *testing.T) {
	c := NewControllers(nil, repository.NewFileMockRepository(), nil, mock.NewMockConverter())

	correctCtx, correctRec := testhelpers.CreateContext(httptest.NewRequest("POST", "/", nil), map[string]any{
		"options": convert.ConversionOptions{},
		"file":    &multipart.FileHeader{Filename: "test.jpg"},
	})

	errConvertImageCtx, errConvertImageRec := testhelpers.CreateContext(httptest.NewRequest("POST", "/", nil), map[string]any{
		"options": convert.ConversionOptions{},
		"file":    &multipart.FileHeader{Filename: "errorconvert.jpg"},
	})

	errAddCtx, errAddRec := testhelpers.CreateContext(httptest.NewRequest("POST", "/", nil), map[string]any{
		"options": convert.ConversionOptions{},
		"file":    &multipart.FileHeader{Filename: "erroradd.jpg"},
	})

	testCases := []testhelpers.ContextCase{
		testhelpers.CreateContextCase(correctCtx, correctRec, true, "Correct"),
		testhelpers.CreateContextCase(errConvertImageCtx, errConvertImageRec, false, "ErrorInConvertImage"),
		testhelpers.CreateContextCase(errAddCtx, errAddRec, false, "ErrorInAdd"),
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			c.ImageController(tc.Context)
			tc.CheckCase(t)
		})
	}
}

func TestVideoController(t *testing.T) {
	c := NewControllers(nil, repository.NewFileMockRepository(), nil, mock.NewMockConverter())

	correctCtx, correctRec := testhelpers.CreateContext(httptest.NewRequest("POST", "/", nil), map[string]any{
		"options": convert.ConversionOptions{},
		"file":    &multipart.FileHeader{Filename: "test.mp4"},
	})

	errConvertVideoCtx, errConvertVideoRec := testhelpers.CreateContext(httptest.NewRequest("POST", "/", nil), map[string]any{
		"options": convert.ConversionOptions{},
		"file":    &multipart.FileHeader{Filename: "errorconvert.mp4"},
	})

	errAddCtx, errAddRec := testhelpers.CreateContext(httptest.NewRequest("POST", "/", nil), map[string]any{
		"options": convert.ConversionOptions{},
		"file":    &multipart.FileHeader{Filename: "erroradd.mp4"},
	})

	testCases := []testhelpers.ContextCase{
		testhelpers.CreateContextCase(correctCtx, correctRec, true, "Correct"),
		testhelpers.CreateContextCase(errConvertVideoCtx, errConvertVideoRec, false, "ErrorInConvertVideo"),
		testhelpers.CreateContextCase(errAddCtx, errAddRec, false, "ErrorInAdd"),
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			c.VideoController(tc.Context)
			tc.CheckCase(t)
		})
	}
}

func TestGetImage(t *testing.T) {
	c := NewControllers(nil, repository.NewFileMockRepository(), nil, nil)

	config.Values.OutputImagePath = "../../testdata/images"

	correctCtx, correctRec := testhelpers.CreateContext(httptest.NewRequest("POST", "http://localhost:8080/api/i/pepe_jpg.jpg", nil), nil)
	correctCtx.Params = append(correctCtx.Params, gin.Param{
		Key:   "filename",
		Value: "pepe_jpg.jpg",
	})

	errIsExistsCtx, errIsExistsRec := testhelpers.CreateContext(httptest.NewRequest("POST", "http://localhost:8080/api/i/errisexists.jpg", nil), nil)
	errIsExistsCtx.Params = append(errIsExistsCtx.Params, gin.Param{
		Key:   "filename",
		Value: "errisexists.jpg",
	})

	errOpenCtx, errOpenRec := testhelpers.CreateContext(httptest.NewRequest("POST", "http://localhost:8080/api/i/erropen.jpg", nil), nil)
	errOpenCtx.Params = append(errOpenCtx.Params, gin.Param{
		Key:   "filename",
		Value: "erropen.jpg",
	})

	testCases := []testhelpers.ContextCase{
		testhelpers.CreateContextCase(correctCtx, correctRec, true, "Correct"),
		testhelpers.CreateContextCase(errIsExistsCtx, errIsExistsRec, false, "ErrorInIsExists"),
		testhelpers.CreateContextCase(errOpenCtx, errOpenRec, false, "ErrorInOsOpen"),
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			c.GetImage(tc.Context)
			tc.CheckCase(t)
		})
	}
}

func TestGetVideo(t *testing.T) {
	c := NewControllers(nil, repository.NewFileMockRepository(), nil, nil)

	config.Values.OutputVideoPath = "../../testdata/images"

	correctCtx, correctRec := testhelpers.CreateContext(httptest.NewRequest("POST", "http://localhost:8080/api/v/pepe_gif.gif", nil), nil)
	correctCtx.Params = append(correctCtx.Params, gin.Param{
		Key:   "filename",
		Value: "pepe_gif.gif",
	})

	errIsExistsCtx, errIsExistsRec := testhelpers.CreateContext(httptest.NewRequest("POST", "http://localhost:8080/api/v/errisexists.mp4", nil), nil)
	errIsExistsCtx.Params = append(errIsExistsCtx.Params, gin.Param{
		Key:   "filename",
		Value: "errisexists.mp4",
	})

	errOpenCtx, errOpenRec := testhelpers.CreateContext(httptest.NewRequest("POST", "http://localhost:8080/api/v/erropen.mp4", nil), nil)
	errOpenCtx.Params = append(errOpenCtx.Params, gin.Param{
		Key:   "filename",
		Value: "erropen.mp4",
	})

	testCases := []testhelpers.ContextCase{
		testhelpers.CreateContextCase(correctCtx, correctRec, true, "Correct"),
		testhelpers.CreateContextCase(errIsExistsCtx, errIsExistsRec, false, "ErrorInIsExists"),
		testhelpers.CreateContextCase(errOpenCtx, errOpenRec, false, "ErrorInOsOpen"),
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			c.GetVideo(tc.Context)
			tc.CheckCase(t)
		})
	}
}
