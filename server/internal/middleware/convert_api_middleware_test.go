package middleware

import (
	"encoding/json"
	"hara/internal/convert"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestOptionsFieldProvided(t *testing.T) {
	testCases := []requestCase{
		{
			createFormRequest(map[string]string{
				"options": "{value}",
			}, "/"),
			true,
			"CorrectFieldName",
		},
		{
			createFormRequest(map[string]string{
				"not_options": "{value}",
			}, "/"),
			false,
			"IncorrectFieldName",
		},
	}

	for _, c := range testCases {
		t.Run(c.name, func(t *testing.T) {
			testRequest(t, c, OptionsFieldProvided)
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

func TestValidateVideoOptionsJson(t *testing.T) {
	var options convert.ConversionVideoOptions
	optionsJson, _ := json.Marshal(options)

	testCases := []requestCase{
		{
			createFormRequest(map[string]string{
				"options": string(optionsJson),
			}, "/"),
			true,
			"CorrectJson",
		},
		{
			createFormRequest(map[string]string{
				"options": "not_json",
			}, "/"),
			false,
			"IncorrectJson",
		},
	}

	for _, c := range testCases {
		t.Run(c.name, func(t *testing.T) {
			ctx := testRequest(t, c, ValidateVideoOptionsJson)

			got, ok := ctx.Get("options")
			if !ok {
				t.Fatalf("%s: options not set", c.name)
			}

			if got.(convert.ConversionVideoOptions) != options {
				t.Fatalf("%s: options not equal options", c.name)
			}
		})
	}
}

func TestValidateImageOptionsJson(t *testing.T) {
	var options convert.ConversionImageOptions
	optionsJson, _ := json.Marshal(options)

	testCases := []requestCase{
		{
			createFormRequest(map[string]string{
				"options": string(optionsJson),
			}, "/"),
			true,
			"CorrectJson",
		},
		{
			createFormRequest(map[string]string{
				"options": "not_json",
			}, "/"),
			false,
			"IncorrectJson",
		},
	}

	for _, c := range testCases {
		t.Run(c.name, func(t *testing.T) {
			ctx := testRequest(t, c, ValidateImageOptionsJson)

			got, ok := ctx.Get("options")
			if !ok {
				t.Fatalf("%s: options not set", c.name)
			}

			if got.(convert.ConversionImageOptions) != options {
				t.Fatalf("%s: options not equal options", c.name)
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
				if (c.lifetime < hourInSeconds || c.lifetime > monthInSeconds) && checkedLifetime != hourInSeconds {
					t.Fatalf("%s want %d, got %d", c.name, hourInSeconds, checkedLifetime)
				}
			}
		})
	}
}
