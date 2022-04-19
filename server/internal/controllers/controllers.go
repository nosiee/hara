package controllers

import (
	"hara/internal/models"
)

type Controllers struct {
	UserRepository   models.UserRepository
	FileRepository   models.FileRepository
	ApikeyRepository models.ApiKeyRepository
	Converter        models.Converter
}

func NewControllers(usersRepo models.UserRepository, filesRepo models.FileRepository, apikeyRepo models.ApiKeyRepository, converter models.Converter) *Controllers {
	return &Controllers{
		usersRepo,
		filesRepo,
		apikeyRepo,
		converter,
	}
}
