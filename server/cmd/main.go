package main

import (
	"database/sql"
	"flag"
	"hara/internal/api"
	"hara/internal/config"
	"hara/internal/controllers"
	"hara/internal/repository"
	"os"

	_ "github.com/lib/pq"
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

	db, err := sql.Open("postgres", config.Values.DatabaseURL)
	if err != nil {
		panic(err)
	}

	userRepo := repository.NewUserRepository(db)
	fileRepo := repository.NewFileRepository(db)
	apikeyRepo := repository.NewApiKeyRepository(db)

	controllers := controllers.NewControllers(userRepo, fileRepo, apikeyRepo)

	server := api.NewServer(controllers)
	server.RunServer(config.Values.APIEndPoint)
}
