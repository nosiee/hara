package config

import (
	"os"
	"testing"
)

func TestLoadFromEnv(t *testing.T) {
	correctEnvs := make(map[string]string, 6)
	correctEnvs["API_ENDPOINT"] = ":8080"
	correctEnvs["UPLOAD_IMAGE_PATH"] = "upload/images"
	correctEnvs["UPLOAD_VIDEO_PATH"] = "upload/videos"
	correctEnvs["OUTPUT_IMAGE_PATH"] = "output/images"
	correctEnvs["OUTPUT_VIDEO_PATH"] = "output/videos"
	correctEnvs["DATABASE_URL"] = "postgresql://localhost/mydb?user=other&password=secret"

	incorrectEnvs := make(map[string]string, 1)
	incorrectEnvs["API_ENDPOINT"] = ":8080"

	testCases := []struct {
		envVars map[string]string
		correct bool
		name    string
	}{
		{
			correctEnvs,
			true,
			"CurrectEnv",
		},
		{
			incorrectEnvs,
			false,
			"IncorrectEnv",
		},
	}

	for _, c := range testCases {
		for k, v := range c.envVars {
			os.Setenv(k, v)
		}

		err := LoadFromEnv()

		for k := range c.envVars {
			os.Unsetenv(k)
		}

		if c.correct && err != nil {
			t.Fatalf("%s: want err == nil, got %s", c.name, err)
		}

		if !c.correct && err == nil {
			t.Fatalf("%s: want err != nil, got nil", c.name)
		}
	}

}

func TestLoadFromFile(t *testing.T) {
	testCases := []struct {
		fpath   string
		correct bool
		name    string
	}{
		{
			"../../testdata/configs/config_test_correct.toml",
			true,
			"CorrectFile",
		},
		{
			"../../testdata/configs/config_test_incorrect.toml",
			false,
			"IncorrectFile",
		},
	}

	for _, c := range testCases {
		t.Run(c.name, func(t *testing.T) {
			err := LoadFromFile(c.fpath)

			if c.correct && err != nil {
				t.Fatalf("%s: want err == nil, got %v", c.name, err)
			}

			if !c.correct && err == nil {
				t.Fatalf("%s: want err != nil, got nil", c.name)
			}
		})
	}
}
