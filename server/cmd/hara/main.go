package main

import (
	"database/sql"
	"flag"
	"fmt"
	"hara/internal/api"
	"hara/internal/config"
	"hara/internal/controllers"
	"hara/internal/convert"
	"hara/internal/repository"
	"os"
	"time"

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

	db, err := connectToDB(config.Values.DatabaseDriver, config.Values.DatabaseURL, 5)
	if err != nil {
		logrus.Fatal(err)
	}

	userRepo := repository.NewUserRepository(db)
	fileRepo := repository.NewFileRepository(db)
	apikeyRepo := repository.NewApiKeyRepository(db)
	converter := convert.NewConverter()
	converter.Initialize()

	controllers := controllers.NewControllers(userRepo, fileRepo, apikeyRepo, converter)

	go fileRepo.DeleteExpired()
	go apikeyRepo.UpdateAllQuota()

	createMediaFolders(config.Values.UploadImagePath, config.Values.UploadVideoPath, config.Values.OutputImagePath, config.Values.OutputVideoPath)

	server := api.NewServer(controllers)
	server.RunServer(config.Values.APIEndPoint)
}

func connectToDB(driver, dburl string, maxRetries int) (*sql.DB, error) {
	var db *sql.DB
	var err error
	var retries int

	for {
		if retries >= maxRetries {
			return nil, fmt.Errorf("Couldn't connect to the database after %d retries. Last error: %v", retries, err)
		}

		db, err = sql.Open(driver, dburl)
		if err != nil {
			logrus.Error(err)
			time.Sleep(time.Second * 10)

			retries++
			continue
		}

		if err := db.Ping(); err != nil {
			logrus.Error(err)
			time.Sleep(time.Second * 10)

			retries++
			continue
		}

		return db, nil
	}
}

func createMediaFolders(uplImg, uplVid, outImg, outVid string) {
	if err := os.MkdirAll(uplImg, 0700); err != nil {
		logrus.Error(err)
	}
	if err := os.MkdirAll(uplVid, 0700); err != nil {
		logrus.Error(err)
	}

	if err := os.MkdirAll(outImg, 0700); err != nil {
		logrus.Error(err)
	}
	if err := os.MkdirAll(outVid, 0700); err != nil {
		logrus.Error(err)
	}
}
