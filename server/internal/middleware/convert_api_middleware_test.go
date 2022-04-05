package middleware

import (
	"hara/internal/convert"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestOptionsFieldProvided(t *testing.T) {
	testCases := []requestCase{
		{
			createFormRequest(nil, "https://localhost:8080/api/convert/image?key=testkey&ext=mp4"),
			true,
			"CorrectUrlQuery",
		},
		{
			createFormRequest(nil, "https://localhost:8080/api/convert/image"),
			false,
			"IncorrectUrlQuery",
		},
	}

	for _, c := range testCases {
		t.Run(c.name, func(t *testing.T) {
			testRequest(t, c, ConversionOptionsProvided)
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
			testRequest(t, c, FileFieldProvided)
		})
	}
}

func TestSupportedVideoFileFormat(t *testing.T) {
	correctVideoFormat := convert.ConversionOptions{Extension: "mp4"}
	incorrectVideoFormat := convert.ConversionOptions{Extension: "exe"}

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
	correctImageFormat := convert.ConversionOptions{Extension: "jpg"}
	incorrectImageFormat := convert.ConversionOptions{Extension: "exe"}

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
	correctLifetime := convert.ConversionOptions{Lifetime: 4200}
	moreThanMonthLifetimeImage := convert.ConversionOptions{Lifetime: 934579384759374985}
	lessThanHourLifetimeImage := convert.ConversionOptions{Lifetime: 120}

	lessThanHourLifetimeVideo := convert.ConversionOptions{Lifetime: 120}
	moreThanMonthLifetimeVideo := convert.ConversionOptions{Lifetime: 934579384759374985}

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

			checkedLifetime = context.(convert.ConversionOptions).Lifetime

			if c.correct && c.lifetime != checkedLifetime {
				t.Fatalf("%s want %d, got %d", c.name, c.lifetime, checkedLifetime)
			}

			if !c.correct {
				if (c.lifetime < hourInSeconds || c.lifetime > monthInSeconds) && checkedLifetime != hourInSeconds {
					t.Fatalf("%s want %d, got %d", c.name, hourInSeconds, checkedLifetime)
				}
			}
		})
	}
}
