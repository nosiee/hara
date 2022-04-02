package config

import (
	"github.com/BurntSushi/toml"
	"github.com/kelseyhightower/envconfig"
)

type values struct {
	APIEndPoint     string `envconfig:"API_ENDPOINT" required:"true"`
	JWTKey          string `envconfig:"JWT_KEY" required:"true"`
	DatabaseURL     string `envconfig:"DATABASE_URL" required:"true"`
	UploadImagePath string `envconfig:"UPLOAD_IMAGE_PATH" required:"true"`
	UploadVideoPath string `envconfig:"UPLOAD_VIDEO_PATH" required:"true"`
	OutputImagePath string `envconfig:"OUTPUT_IMAGE_PATH" required:"true"`
	OutputVideoPath string `envconfig:"OUTPUT_VIDEO_PATH" required:"true"`
}

var Values values

func LoadFromEnv() error {
	err := envconfig.Process("", &Values)
	return err
}

func LoadFromFile(fpath string) error {
	_, err := toml.DecodeFile(fpath, &Values)
	return err
}
