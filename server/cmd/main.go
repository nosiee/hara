package main

import (
	"database/sql"
	"flag"
	"hara/internal/api"
	"hara/internal/config"
	"hara/internal/controllers"
	"hara/internal/convert"
	"hara/internal/repository"
	"os"

	"github.com/sirupsen/logrus"

	_ "github.com/lib/pq"
)

func main() {
	var file string
	var help bool

	flag.StringVar(&file, "config", "", "Config .toml file")
	flag.BoolVar(&help, "help", false, "Print help information")
	flag.Parse()

	logrus.SetReportCaller(true)
	logrus.SetLevel(logrus.TraceLevel)

	if help {
		flag.Usage()
		os.Exit(1)
	}

	if len(os.Args) == 1 {
		if err := config.LoadFromEnv(); err != nil {
			logrus.Panic(err)
		}
	} else {
		if err := config.LoadFromFile(file); err != nil {
			logrus.Panic(err)
		}
	}

	db, err := sql.Open("postgres", config.Values.DatabaseURL)
	if err != nil {
		logrus.Panic(err)
	}

	if err := db.Ping(); err != nil {
		logrus.Panic(err)
	}

	userRepo := repository.NewUserRepository(db)
	fileRepo := repository.NewFileRepository(db)
	apikeyRepo := repository.NewApiKeyRepository(db)
	converter := convert.NewConverter()
	converter.Initialize()

	controllers := controllers.NewControllers(userRepo, fileRepo, apikeyRepo, converter)

	go fileRepo.DeleteExpired()
	go apikeyRepo.UpdateAllQuota()

	server := api.NewServer(controllers)
	server.RunServer(config.Values.APIEndPoint)
}
