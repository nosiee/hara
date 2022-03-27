package config

import (
	"os"
	"testing"
)

func TestLoadFromEnv(t *testing.T) {
	correctEnvs := make(map[string]string, 5)
	correctEnvs["API_ENDPOINT"] = ":8080"
	correctEnvs["UPLOAD_IMAGE_PATH"] = "upload/images"
	correctEnvs["UPLOAD_VIDEO_PATH"] = "upload/videos"
	correctEnvs["OUTPUT_IMAGE_PATH"] = "output/images"
	correctEnvs["OUTPUT_VIDEO_PATH"] = "output/videos"

	incorrectEnvs := make(map[string]string, 1)
	incorrectEnvs["API_ENDPOINT"] = ":8080"

	testCases := []struct {
		envVars  map[string]string
		correct  bool
		caseName string
	}{
		{
			correctEnvs,
			true,
			"CurrectCaseEnv",
		},
		{
			incorrectEnvs,
			false,
			"IncorrectCaseEnv",
		},
	}

	for _, c := range testCases {
		for k, v := range c.envVars {
			os.Setenv(k, v)
		}

		err := LoadFromEnv()

		for k, _ := range c.envVars {
			os.Unsetenv(k)
		}

		if c.correct && err != nil {
			t.Fatalf("%s: want err == nil, got %s", c.caseName, err)
		}

		if !c.correct && err == nil {
			t.Fatalf("%s: want err != nil, got nil", c.caseName)
		}
	}

}

func TestLoadFromFile(t *testing.T) {
	testCases := []struct {
		fpath    string
		correct  bool
		caseName string
	}{
		{
			"testconfigs/config_test_correct.toml",
			true,
			"CorrectCaseFile",
		},
		{
			"testconfigs/config_test_incorrect.toml",
			false,
			"IncorrectCaseFile",
		},
	}

	for _, c := range testCases {
		t.Run(c.caseName, func(t *testing.T) {
			err := LoadFromFile(c.fpath)

			if c.correct && err != nil {
				t.Fatalf("%s: want err == nil, got %v", c.caseName, err)
			}

			if !c.correct && err == nil {
				t.Fatalf("%s: want err != nil, got nil", c.caseName)
			}
		})
	}
}
