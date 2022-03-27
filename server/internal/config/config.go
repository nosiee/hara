package config

import (
	"os"

	"github.com/BurntSushi/toml"
)

var (
	APIEndPoint     string
	UploadVideoPath string
	UploadImagePath string
	OutputVideoPath string
	OutputImagePath string
)

type config struct {
	APIEndPoint     string
	UploadVideoPath string
	UploadImagePath string
	OutputVideoPath string
	OutputImagePath string
}

func LoadFromEnv(apiEnv, uplVidEnv, uplImgEnv, outVidEnv, outImgEnv string) error {
	// TODO: check for empty variables
	APIEndPoint = os.Getenv(apiEnv)
	UploadImagePath = os.Getenv(uplImgEnv)
	UploadVideoPath = os.Getenv(uplVidEnv)
	OutputImagePath = os.Getenv(outImgEnv)
	OutputVideoPath = os.Getenv(outVidEnv)

	return nil
}

func LoadFromFile(fpath string) error {
	var conf config
	_, err := toml.DecodeFile(fpath, &conf)

	APIEndPoint = conf.APIEndPoint
	UploadImagePath = conf.UploadImagePath
	UploadVideoPath = conf.UploadVideoPath
	OutputImagePath = conf.OutputImagePath
	OutputVideoPath = conf.OutputVideoPath

	return err
}
