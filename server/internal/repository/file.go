package repository

import (
	"database/sql"
	"fmt"
	"hara/internal/models"
	"os"
	"time"
)

type FileRepository struct {
	db *sql.DB
}

func NewFileRepository(db *sql.DB) *FileRepository {
	return &FileRepository{
		db,
	}
}

func (repo FileRepository) Add(file models.File) error {
	_, err := repo.db.Exec("INSERT INTO files(filename,fullpath,deletetime) VALUES($1, $2, $3)", file.Filename, file.Fullpath, file.Deletetime)
	return err
}

func (repo FileRepository) IsExists(filename string) bool {
	var ID int
	_ = repo.db.QueryRow("SELECT id FROM files WHERE filename=$1", filename).Scan(&ID)
	return ID != 0
}

func (repo FileRepository) DeleteExpired() {
	var ID, deletetime int64
	var fullpath string

	for {
		rows, err := repo.db.Query("SELECT id,fullpath,deletetime FROM files")
		if err != nil {
			fmt.Println(err)
			continue
		}

		for rows.Next() {
			rows.Scan(&ID, &fullpath, &deletetime)

			now := time.Now().Unix()
			if now >= deletetime {
				repo.db.Exec("DELETE FROM files where id=$1", ID)
				os.Remove(fullpath)
			}
		}

		time.Sleep(time.Hour)
	}
}
