package middleware

import (
	"hara/internal/convert"
	"net/http/httptest"
	"net/url"
	"strconv"
	"testing"
)

func TestImageFileFieldProvided(t *testing.T) {
	correctImgCtx, correctImgRec := createContextWithRequest(createFileFormRequest("file", "1.jpg"))
	incorrectImgExtCtx, incorrectImgExtRec := createContextWithRequest(createFileFormRequest("file", "1.exe"))
	incorrectImgFieldCtx, incorrectImgFieldRec := createContextWithRequest(createFileFormRequest("not_file", "1.jpg"))

	testCases := []contextCase{
		createContextCase(correctImgCtx, correctImgRec, true, "CorrectImageExtension"),
		createContextCase(incorrectImgExtCtx, incorrectImgExtRec, false, "IncorrectImageExtension"),
		createContextCase(incorrectImgFieldCtx, incorrectImgFieldRec, false, "IncorrectImageField"),
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ImageFileFieldProvided(tc.context)
			tc.checkCase(t)
		})
	}
}

func TestVideoFileFieldProvided(t *testing.T) {
	correctVidCtx, correctVidRec := createContextWithRequest(createFileFormRequest("file", "1.mp4"))
	incorrectVidExtCtx, incorrectVidExtRec := createContextWithRequest(createFileFormRequest("file", "1.exe"))
	incorrectVidFieldCtx, incorrectVidFieldRec := createContextWithRequest(createFileFormRequest("not_file", "1.mp4"))

	testCases := []contextCase{
		createContextCase(correctVidCtx, correctVidRec, true, "CorrectVideoExtension"),
		createContextCase(incorrectVidExtCtx, incorrectVidExtRec, false, "IncorrectVideoExtension"),
		createContextCase(incorrectVidFieldCtx, incorrectVidFieldRec, false, "IncorrectVideoield"),
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			VideoFileFieldProvided(tc.context)
			tc.checkCase(t)
		})
	}
}

func TestSupportedImageFileExtension(t *testing.T) {
	correctImgCtx, correctImgRec := createContextWithRequest(httptest.NewRequest("POST", "http://localhost:8080/api/convert/image?key=apikey&ext=jpg", nil))
	incorrectImgCtx, incorrectImgRec := createContextWithRequest(httptest.NewRequest("POST", "http://localhost:8080/api/convert/image?key=apikey&ext=exe", nil))

	testCases := []contextCase{
		createContextCase(correctImgCtx, correctImgRec, true, "CorrectImageExtension"),
		createContextCase(incorrectImgCtx, incorrectImgRec, false, "IncorrectImageExtension"),
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			SupportedImageFileExtension(tc.context)
			tc.checkCase(t)
		})
	}
}

func TestSupportedVideoFileExtension(t *testing.T) {
	correctVidCtx, correctVidRec := createContextWithRequest(httptest.NewRequest("POST", "http://localhost:8080/api/convert/video?key=apikey&ext=mp4", nil))
	incorrectVidCtx, incorrectVidRec := createContextWithRequest(httptest.NewRequest("POST", "http://localhost:8080/api/convert/video?key=apikey&ext=exe", nil))

	testCases := []contextCase{
		createContextCase(correctVidCtx, correctVidRec, true, "CorrectVideoExtension"),
		createContextCase(incorrectVidCtx, incorrectVidRec, false, "IncorrectVideoExtension"),
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			SupportedVideoFileExtension(tc.context)
			tc.checkCase(t)
		})
	}
}

func TestValidateLifetime(t *testing.T) {
	correctLifetimeCtx, correctLifetimeRec := createContextWithRequest(httptest.NewRequest("POST", "http://localhost:8080/api/convert/image?key=apikey&ext=jpg&lifetime=3600", nil))
	incorrectLifetimeCtx, incorrectLifetimeRec := createContextWithRequest(httptest.NewRequest("POST", "http://localhost:8080/api/convert/image?key=apikey&ext=jpg&lifetime=120", nil))
	withoutLifetimeCtx, withoutLifetimeRec := createContextWithRequest(httptest.NewRequest("POST", "http://localhost:8080/api/convert/image?key=apikey&ext=jpg", nil))

	testCases := []contextCase{
		createContextCase(correctLifetimeCtx, correctLifetimeRec, true, "CorrectLifetime"),
		createContextCase(incorrectLifetimeCtx, incorrectLifetimeRec, false, "IncorrrectLifetime"),
		createContextCase(withoutLifetimeCtx, withoutLifetimeRec, false, "WithoutLifetime"),
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ValidateLifetime(tc.context)
			tc.checkCase(t)
		})
	}
}

func TestExtractConversionOptions(t *testing.T) {
	imageOptionsCtx, _ := createContextWithRequest(httptest.NewRequest("POST", "http://localhost:8080/api/covert/image?key=apikey&ext=jpg&lifetime=3600", nil))
	videoOptionsCtx, _ := createContextWithRequest(httptest.NewRequest("POST", "http://localhost:8080/api/covert/video?key=apikey&ext=mp4&lifetime=4800", nil))

	testCases := []contextCase{
		createContextCase(imageOptionsCtx, nil, true, "ImageOptions"),
		createContextCase(videoOptionsCtx, nil, true, "VideoOptions"),
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ExtractConversionOptions(tc.context)

			values, _ := url.ParseQuery(tc.context.Request.URL.RawQuery)
			options, ok := tc.context.Get("options")

			if !ok {
				t.Fatalf("%s context.Get('options') is nil", tc.name)
			}

			if values.Get("ext") != options.(convert.ConversionOptions).Extension {
				t.Fatalf("%s extensions don't match", tc.name)
			}

			lifetime, _ := strconv.ParseUint(values.Get("lifetime"), 10, 64)

			if lifetime != options.(convert.ConversionOptions).Lifetime {
				t.Fatalf("%s lifetime don't match", tc.name)
			}
		})
	}
}
