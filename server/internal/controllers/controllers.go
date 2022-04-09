package controllers

import "hara/internal/models"

type Controllers struct {
	UserRepository   models.UserRepository
	FileRepository   models.FileRepository
	ApikeyRepository models.ApiKeyRepository
}

func NewControllers(usersRepo models.UserRepository, filesRepo models.FileRepository, apikeyRepo models.ApiKeyRepository) *Controllers {
	return &Controllers{
		usersRepo,
		filesRepo,
		apikeyRepo,
	}
}
