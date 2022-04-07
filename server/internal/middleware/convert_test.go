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
	correctImgCtx, correctImgRec := testhelpers.CreateContextWithRequest(testhelpers.CreateFileFormRequest("file", "1.jpg"))
	incorrectImgExtCtx, incorrectImgExtRec := testhelpers.CreateContextWithRequest(testhelpers.CreateFileFormRequest("file", "1.exe"))
	incorrectImgFieldCtx, incorrectImgFieldRec := testhelpers.CreateContextWithRequest(testhelpers.CreateFileFormRequest("not_file", "1.jpg"))

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
	correctVidCtx, correctVidRec := testhelpers.CreateContextWithRequest(testhelpers.CreateFileFormRequest("file", "1.mp4"))
	incorrectVidExtCtx, incorrectVidExtRec := testhelpers.CreateContextWithRequest(testhelpers.CreateFileFormRequest("file", "1.exe"))
	incorrectVidFieldCtx, incorrectVidFieldRec := testhelpers.CreateContextWithRequest(testhelpers.CreateFileFormRequest("not_file", "1.mp4"))

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
	correctImgCtx, correctImgRec := testhelpers.CreateContextWithRequest(httptest.NewRequest("POST", "http://localhost:8080/api/convert/image?key=apikey&ext=jpg", nil))
	incorrectImgCtx, incorrectImgRec := testhelpers.CreateContextWithRequest(httptest.NewRequest("POST", "http://localhost:8080/api/convert/image?key=apikey&ext=exe", nil))

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
	correctVidCtx, correctVidRec := testhelpers.CreateContextWithRequest(httptest.NewRequest("POST", "http://localhost:8080/api/convert/video?key=apikey&ext=mp4", nil))
	incorrectVidCtx, incorrectVidRec := testhelpers.CreateContextWithRequest(httptest.NewRequest("POST", "http://localhost:8080/api/convert/video?key=apikey&ext=exe", nil))

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
	correctLifetimeCtx, correctLifetimeRec := testhelpers.CreateContextWithRequest(httptest.NewRequest("POST", "http://localhost:8080/api/convert/image?key=apikey&ext=jpg&lifetime=3600", nil))
	incorrectLifetimeCtx, incorrectLifetimeRec := testhelpers.CreateContextWithRequest(httptest.NewRequest("POST", "http://localhost:8080/api/convert/image?key=apikey&ext=jpg&lifetime=120", nil))
	withoutLifetimeCtx, withoutLifetimeRec := testhelpers.CreateContextWithRequest(httptest.NewRequest("POST", "http://localhost:8080/api/convert/image?key=apikey&ext=jpg", nil))

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
	imageOptionsCtx, _ := testhelpers.CreateContextWithRequest(httptest.NewRequest("POST", "http://localhost:8080/api/covert/image?key=apikey&ext=jpg&lifetime=3600", nil))
	videoOptionsCtx, _ := testhelpers.CreateContextWithRequest(httptest.NewRequest("POST", "http://localhost:8080/api/covert/video?key=apikey&ext=mp4&lifetime=4800", nil))

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

			lifetime, _ := strconv.ParseUint(values.Get("lifetime"), 10, 64)

			if lifetime != options.(convert.ConversionOptions).Lifetime {
				t.Fatalf("%s lifetime don't match", tc.Name)
			}
		})
	}
}
