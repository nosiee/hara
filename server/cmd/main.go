package main

import (
	"flag"
	"hara/internal/api"
	"hara/internal/config"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		var confFile string

		flag.StringVar(&confFile, "config", "", "Config .toml file")
		flag.Parse()

		config.LoadFromFile(confFile)
	} else {
		config.LoadFromEnv("API_ENDPOINT", "UPLOAD_VIDEO_PATH", "UPLOAD_IMAGE_PATH", "OUTPUT_VIDEO_PATH", "OUTPUT_IMAGE_PATH")
	}

	api.RunServer(config.APIEndPoint)
}
