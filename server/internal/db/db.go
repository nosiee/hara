package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

var db *sql.DB

func Connnect(url string) error {
	var err error

	db, err = sql.Open("postgres", url)
	if err != nil {
		return err
	}

	if err = db.Ping(); err != nil {
		return err
	}

	go deleteExpiredFilesTicker()

	return nil
}

func Close() {
	closeTicker <- struct{}{}
	db.Close()
}
