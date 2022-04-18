package controllers

import (
	"hara/internal/models"
	"log"
	"os"
)

type Controllers struct {
	UserRepository   models.UserRepository
	FileRepository   models.FileRepository
	ApikeyRepository models.ApiKeyRepository
	Converter        models.Converter
	InfoLogger       *log.Logger
	ErrLogger        *log.Logger
}

func NewControllers(usersRepo models.UserRepository, filesRepo models.FileRepository, apikeyRepo models.ApiKeyRepository, converter models.Converter) *Controllers {
	return &Controllers{
		usersRepo,
		filesRepo,
		apikeyRepo,
		converter,
		log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime),
		log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile),
	}
}
