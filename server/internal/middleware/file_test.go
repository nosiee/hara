package middleware

import (
	"hara/internal/convert"
	"hara/internal/testhelpers"
	"net/http/httptest"
	"net/url"
	"strconv"
	"testing"
)

func TestImageFileFieldProvided(t *testing.T) {
	correctImgCtx, correctImgRec := testhelpers.CreateContext(testhelpers.CreateFileFormRequest("file", "1.jpg"), nil)
	incorrectImgExtCtx, incorrectImgExtRec := testhelpers.CreateContext(testhelpers.CreateFileFormRequest("file", "1.exe"), nil)
	incorrectImgFieldCtx, incorrectImgFieldRec := testhelpers.CreateContext(testhelpers.CreateFileFormRequest("not_file", "1.jpg"), nil)

	testCases := []testhelpers.ContextCase{
		testhelpers.CreateContextCase(correctImgCtx, correctImgRec, true, "CorrectImageExtension"),
		testhelpers.CreateContextCase(incorrectImgExtCtx, incorrectImgExtRec, false, "IncorrectImageExtension"),
		testhelpers.CreateContextCase(incorrectImgFieldCtx, incorrectImgFieldRec, false, "IncorrectImageField"),
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			ImageFileFieldProvided(tc.Context)
			tc.CheckCase(t)
		})
	}
}

func TestVideoFileFieldProvided(t *testing.T) {
	correctVidCtx, correctVidRec := testhelpers.CreateContext(testhelpers.CreateFileFormRequest("file", "1.mp4"), nil)
	incorrectVidExtCtx, incorrectVidExtRec := testhelpers.CreateContext(testhelpers.CreateFileFormRequest("file", "1.exe"), nil)
	incorrectVidFieldCtx, incorrectVidFieldRec := testhelpers.CreateContext(testhelpers.CreateFileFormRequest("not_file", "1.mp4"), nil)

	testCases := []testhelpers.ContextCase{
		testhelpers.CreateContextCase(correctVidCtx, correctVidRec, true, "CorrectVideoExtension"),
		testhelpers.CreateContextCase(incorrectVidExtCtx, incorrectVidExtRec, false, "IncorrectVideoExtension"),
		testhelpers.CreateContextCase(incorrectVidFieldCtx, incorrectVidFieldRec, false, "IncorrectVideoield"),
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			VideoFileFieldProvided(tc.Context)
			tc.CheckCase(t)
		})
	}
}

func TestSupportedImageFileExtension(t *testing.T) {
	correctImgCtx, correctImgRec := testhelpers.CreateContext(httptest.NewRequest("POST", "http://localhost:8080/api/convert/image?key=apikey&ext=jpg", nil), nil)
	incorrectImgCtx, incorrectImgRec := testhelpers.CreateContext(httptest.NewRequest("POST", "http://localhost:8080/api/convert/image?key=apikey&ext=exe", nil), nil)

	testCases := []testhelpers.ContextCase{
		testhelpers.CreateContextCase(correctImgCtx, correctImgRec, true, "CorrectImageExtension"),
		testhelpers.CreateContextCase(incorrectImgCtx, incorrectImgRec, false, "IncorrectImageExtension"),
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			SupportedImageFileExtension(tc.Context)
			tc.CheckCase(t)
		})
	}
}

func TestSupportedVideoFileExtension(t *testing.T) {
	correctVidCtx, correctVidRec := testhelpers.CreateContext(httptest.NewRequest("POST", "http://localhost:8080/api/convert/video?key=apikey&ext=mp4", nil), nil)
	incorrectVidCtx, incorrectVidRec := testhelpers.CreateContext(httptest.NewRequest("POST", "http://localhost:8080/api/convert/video?key=apikey&ext=exe", nil), nil)

	testCases := []testhelpers.ContextCase{
		testhelpers.CreateContextCase(correctVidCtx, correctVidRec, true, "CorrectVideoExtension"),
		testhelpers.CreateContextCase(incorrectVidCtx, incorrectVidRec, false, "IncorrectVideoExtension"),
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			SupportedVideoFileExtension(tc.Context)
			tc.CheckCase(t)
		})
	}
}

func TestValidateLifetime(t *testing.T) {
	correctLifetimeCtx, correctLifetimeRec := testhelpers.CreateContext(httptest.NewRequest("POST", "http://localhost:8080/api/convert/image?key=apikey&ext=jpg&lifetime=3600", nil), nil)
	incorrectLifetimeCtx, incorrectLifetimeRec := testhelpers.CreateContext(httptest.NewRequest("POST", "http://localhost:8080/api/convert/image?key=apikey&ext=jpg&lifetime=120", nil), nil)
	withoutLifetimeCtx, withoutLifetimeRec := testhelpers.CreateContext(httptest.NewRequest("POST", "http://localhost:8080/api/convert/image?key=apikey&ext=jpg", nil), nil)

	testCases := []testhelpers.ContextCase{
		testhelpers.CreateContextCase(correctLifetimeCtx, correctLifetimeRec, true, "CorrectLifetime"),
		testhelpers.CreateContextCase(incorrectLifetimeCtx, incorrectLifetimeRec, false, "IncorrrectLifetime"),
		testhelpers.CreateContextCase(withoutLifetimeCtx, withoutLifetimeRec, false, "WithoutLifetime"),
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			ValidateLifetime(tc.Context)
			tc.CheckCase(t)
		})
	}
}

func TestExtractConversionOptions(t *testing.T) {
	imageOptionsCtx, _ := testhelpers.CreateContext(httptest.NewRequest("POST", "http://localhost:8080/api/covert/image?key=apikey&ext=jpg&lifetime=3600", nil), nil)
	videoOptionsCtx, _ := testhelpers.CreateContext(httptest.NewRequest("POST", "http://localhost:8080/api/covert/video?key=apikey&ext=mp4&lifetime=4800", nil), nil)

	testCases := []testhelpers.ContextCase{
		testhelpers.CreateContextCase(imageOptionsCtx, nil, true, "ImageOptions"),
		testhelpers.CreateContextCase(videoOptionsCtx, nil, true, "VideoOptions"),
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			ExtractConversionOptions(tc.Context)

			values, _ := url.ParseQuery(tc.Context.Request.URL.RawQuery)
			options, ok := tc.Context.Get("options")

			if !ok {
				t.Fatalf("%s context.Get('options') is nil", tc.Name)
			}

			if values.Get("ext") != options.(convert.ConversionOptions).Extension {
				t.Fatalf("%s extensions don't match", tc.Name)
			}

			lifetime, _ := strconv.ParseUint(values.Get("lifetime"), 10, 32)

			if uint(lifetime) != options.(convert.ConversionOptions).Lifetime {
				t.Fatalf("%s lifetime don't match", tc.Name)
			}
		})
	}
}
