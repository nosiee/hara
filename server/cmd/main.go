package main

import (
	"flag"
	"fmt"
	"hara/internal/api"
	"hara/internal/config"
	"hara/internal/db"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		var confFile string

		flag.StringVar(&confFile, "config", "", "Config .toml file")
		flag.Parse()

		config.LoadFromFile(confFile)
	} else {
		config.LoadFromEnv()
	}

	if err := db.Connnect(config.Values.DatabaseURL); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	api.RunServer(config.Values.APIEndPoint)
}
