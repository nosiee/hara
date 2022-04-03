package db

import (
	"fmt"
	"hara/internal/config"
	"os"
	"time"
)

var closeTicker chan struct{}

func AddFileLifetime(fpath, ftype, deleteDate string) error {
	_, err := db.Exec("INSERT INTO lifetimes(filename,filetype,deletedate) VALUES($1, $2, $3)", fpath, ftype, deleteDate)
	return err
}

func IsFileExists(fname string) bool {
	var ID int
	_ = db.QueryRow("SELECT id FROM lifetimes WHERE filename=$1", fname).Scan(&ID)
	return ID != 0
}

func deleteExpiredFilesTicker() {
	ticker := time.NewTicker(time.Hour)
	closeTicker = make(chan struct{})

	for {
		select {
		case <-ticker.C:
			deleteExpiredFiles()
		case <-closeTicker:
			return
		}
	}
}

func deleteExpiredFiles() error {
	r, err := db.Query("SELECT * FROM lifetimes")
	if err != nil {
		return err
	}

	var ID int
	var filename, filetype, deleteDate, fpath string

	for r.Next() {
		if err := r.Scan(&ID, &filename, &filetype, &deleteDate); err != nil {
			return err
		}

		now := time.Now()
		date, _ := time.Parse(time.RFC3339, deleteDate)

		if now.After(date) {
			if _, err := db.Exec(fmt.Sprintf("DELETE FROM lifetimes WHERE ID=%d", ID)); err != nil {
				return err
			}

			switch filetype {
			case "image":
				fpath = fmt.Sprintf("%s/%s", config.Values.OutputImagePath, filename)
			case "video":
				fpath = fmt.Sprintf("%s/%s", config.Values.OutputVideoPath, filename)
			}

			os.Remove(fpath)
		}
	}

	return nil
}
