package repository

import (
	"errors"
	"hara/internal/models"
)

type UserMockRepository struct{}

func NewUserMockRepository() *UserMockRepository {
	return &UserMockRepository{}
}

func (repo UserMockRepository) Add(user models.User) error {
	if user.Username == "errorAdd" {
		return errors.New("Test error")
	}
	return nil
}

func (repo UserMockRepository) FindByUsername(username string) *models.User {
	if username == "nosuchuser" {
		return nil
	}

	return &models.User{
		UUID:     "uuid",
		Username: "nosiee",
		Hash:     "2432612431322437745532724f64467876346b774b57564d7579474f2e3455616b7867386b53596c644367556d3533553445794574487a634d766d57",
		Email:    "nosiee@example.com",
	}
}
