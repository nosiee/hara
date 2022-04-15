package repository

import (
	"database/sql"
	"hara/internal/models"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db,
	}
}

func (repo UserRepository) Add(user models.User) error {
	_, err := repo.db.Exec("INSERT INTO users(uuid, username, hash, email) VALUES($1, $2, $3, $4)", user.UUID, user.Username, user.Hash, user.Email)
	return err
}

func (repo UserRepository) FindByUsername(username string) *models.User {
	var user models.User

	repo.db.QueryRow("SELECT uuid, hash FROM users WHERE username=$1", username).Scan(&user.UUID, &user.Hash)
	return &user
}
