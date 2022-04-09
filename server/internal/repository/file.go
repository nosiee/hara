package repository

import (
	"database/sql"
	"hara/internal/models"
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
	_, err := repo.db.Exec("INSERT INTO files(filename,filetype,lifetime) VALUES($1, $2, $3)", file.Filename, file.Filetype, file.Lifetime)
	return err
}

func (repo FileRepository) IsExists(filename string) bool {
	var ID int
	_ = repo.db.QueryRow("SELECT id FROM files WHERE filename=$1", filename).Scan(&ID)
	return ID != 0
}
