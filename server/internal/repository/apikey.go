package repository

import (
	"database/sql"
	"hara/internal/models"
	"time"

	"github.com/sirupsen/logrus"
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
	_, err := repo.db.Exec("INSERT INTO apikeys(uuid, key, maxquota, quota, updatetime) VALUES($1, $2, $3, $4, $5)", key.OwnerUUID, key.Key, key.MaxQuota, key.Quota, key.Updatetime)
	return err
}

func (repo ApiKeyRepository) ChangeKeyID(uuid, key string) error {
	_, err := repo.db.Exec("UPDATE apikeys SET key = $1 WHERE uuid=$2", key, uuid)
	return err
}

func (repo ApiKeyRepository) UserHasKey(uuid string) (bool, error) {
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

func (repo ApiKeyRepository) GetKey(uuid string) (*models.ApiKey, error) {
	var key models.ApiKey
	var ID int

	err := repo.db.QueryRow("SELECT * FROM apikeys WHERE uuid=$1", uuid).Scan(&ID, &key.OwnerUUID, &key.Key, &key.MaxQuota, &key.Quota, &key.Updatetime)
	return &key, err
}

func (repo ApiKeyRepository) GetQuota(key string) (uint, uint, error) {
	var maxquota, quota uint

	err := repo.db.QueryRow("SELECT maxquota,quota FROM apikeys WHERE key=$1", key).Scan(&maxquota, &quota)
	return maxquota, quota, err
}

func (repo ApiKeyRepository) IncrementQuota(key string) error {
	_, err := repo.db.Exec("UPDATE apikeys SET quota = quota + 1 WHERE key=$1", key)
	return err
}

func (repo ApiKeyRepository) SetQuota(key string, quota uint) error {
	_, err := repo.db.Exec("UPDATE apikeys SET quota = $1 WHERE key=$2", quota, key)
	return err
}

func (repo ApiKeyRepository) SetUpdatetime(key string, time int64) error {
	_, err := repo.db.Exec("UPDATE apikeys SET updatetime = $1 WHERE key = $2", time, key)
	return err
}

func (repo ApiKeyRepository) GetUpdatetime(key string) (int64, error) {
	var updatetime int64

	err := repo.db.QueryRow("SELECT updatetime FROM apikeys WHERE key = $1", key).Scan(&updatetime)
	return updatetime, err
}

func (repo ApiKeyRepository) UpdateAllQuota() {
	var updatetime int64
	var key string

	for {
		rows, err := repo.db.Query("SELECT key, updatetime FROM apikeys")
		if err != nil {
			logrus.Fatal(err)
			break
		}

		for rows.Next() {
			rows.Scan(&key, &updatetime)

			now := time.Now().Unix()
			if now >= updatetime && updatetime > 0 {
				if err := repo.SetQuota(key, 0); err != nil {
					logrus.Fatal(err)
					break
				}

				if err := repo.SetUpdatetime(key, 0); err != nil {
					logrus.Fatal(err)
					break
				}
			}
		}

		time.Sleep(time.Hour)
	}
}
