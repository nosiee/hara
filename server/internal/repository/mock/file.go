package repository

import (
	"errors"
	"hara/internal/models"
)

type FileMockRepository struct{}

func NewFileMockRepository() *FileMockRepository {
	return &FileMockRepository{}
}

func (repo FileMockRepository) Add(file models.File) error {
	switch file.Filename {
	case "erroradd.jpg":
		return errors.New("Test error")
	case "erroradd.mp4":
		return errors.New("Test error")
	default:
		return nil
	}
}

func (repo FileMockRepository) IsExists(filename string) bool {
	switch filename {
	case "errisexists.jpg":
		return false
	case "errisexists.mp4":
		return false
	default:
		return true
	}
}

func (repo FileMockRepository) DeleteExpired() {}
