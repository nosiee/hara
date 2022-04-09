package repository

import (
	"database/sql"
	"hara/internal/models"
)

type ApiKeyRepository struct {
	db *sql.DB
}

func NewApiKeyRepository(db *sql.DB) *ApiKeyRepository {
	return &ApiKeyRepository{
		db,
	}
}

func (repo ApiKeyRepository) Add(key models.ApiKey) error {
	_, err := repo.db.Exec("INSERT INTO apikeys(uuid, key, maxquotas, quotas) VALUES($1, $2, $3, $4)", key.OwnerUUID, key.Key, key.MaxQuotas, key.Quotas)
	return err
}

func (repo ApiKeyRepository) UserHaveKey(uuid string) (bool, error) {
	var ID int
	var key string

	err := repo.db.QueryRow("SELECT id, key FROM apikeys WHERE uuid=$1", uuid).Scan(&ID, &key)
	return ID != 0, err
}

func (repo ApiKeyRepository) IsExists(key string) (bool, error) {
	var ID int

	err := repo.db.QueryRow("SELECT id FROM apikeys WHERE key=$1", key).Scan(&ID)
	return ID != 0, err
}
