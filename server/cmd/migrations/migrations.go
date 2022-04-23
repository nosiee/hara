package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	var dburl, mtype, mpath string
	var retries int
	var maxRetries = 5
	var err error

	dburl = os.Getenv("DATABASE_URL")
	flag.StringVar(&mtype, "mtype", "", "")
	flag.StringVar(&mpath, "mpath", "", "")
	flag.Parse()

	for {
		if retries >= maxRetries {
			panic(fmt.Sprintf("Couldn't connect to the database after %d. Last error: %v", retries, err))
		}

		m, err := migrate.New(fmt.Sprintf("file://%s", mpath), dburl)
		if err != nil {
			fmt.Println(err)
			time.Sleep(10 * time.Second)

			retries++
			continue
		}

		switch mtype {
		case "up":
			m.Up()
			println("Migrations up done.")
			return
		case "down":
			m.Down()
			println("Migrations down done.")
			return
		default:
			panic("Unknow migration type")
		}
	}
}
