package main

import (
	"flag"
	"hara/internal/api"
	"hara/internal/config"
	"hara/internal/db"
	"os"
)

func main() {
	var file string
	var help bool

	flag.StringVar(&file, "config", "", "Config .toml file")
	flag.BoolVar(&help, "help", false, "Print help information")
	flag.Parse()

	if help {
		flag.Usage()
		os.Exit(1)
	}

	if len(os.Args) == 1 {
		if err := config.LoadFromEnv(); err != nil {
			panic(err)
		}
	} else {
		if err := config.LoadFromFile(file); err != nil {
			panic(err)
		}
	}

	if err := db.Connnect(config.Values.DatabaseURL); err != nil {
		panic(err)
	}

	api.RunServer(config.Values.APIEndPoint)
}
