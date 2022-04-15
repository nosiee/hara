package repository

import (
	"errors"
	"hara/internal/models"
)

type ApiKeyMockRepository struct{}

func NewApiKeyMockRepository() *ApiKeyMockRepository {
	return &ApiKeyMockRepository{}
}

func (mock ApiKeyMockRepository) Add(key models.ApiKey) error {
	if key.OwnerUUID == "errorAdd" {
		return errors.New("Test error")
	}

	return nil
}

func (mock ApiKeyMockRepository) UserHasKey(uuid string) (bool, error) {
	switch uuid {
	case "has":
		return true, nil
	case "errorUserHasKey":
		return false, errors.New("Test error")
	default:
		return false, nil
	}
}

func (mock ApiKeyMockRepository) IsExists(key string) (bool, error) {
	switch key {
	case "correct":
		return true, nil
	case "incorrect":
		return false, nil
	case "error":
		return false, errors.New("Test error")
	default:
		return false, nil
	}
}

func (mock ApiKeyMockRepository) GetQuota(key string) (uint, uint, error) {
	switch key {
	case "errorquota":
		return 0, 0, errors.New("Test error")
	case "errorsetupdatetime":
		return 100, 102, nil
	case "exceeded_":
		return 100, 102, nil
	default:
		return 100, 0, nil
	}
}

func (mock ApiKeyMockRepository) IncrementQuota(key string) error {
	switch key {
	case "errorincrementquota":
		return errors.New("Test error")
	default:
		return nil
	}
}

func (mock ApiKeyMockRepository) SetQuota(key string, quota uint) error {
	return nil
}

func (mock ApiKeyMockRepository) SetUpdatetime(key string, time int64) error {
	switch key {
	case "errorsetupdatetime":
		return errors.New("Test error")
	case "exceeded_":
		return nil
	default:
		return nil
	}
}

func (mock ApiKeyMockRepository) GetUpdatetime(key string) (int64, error) {
	switch key {
	case "correct":
		return 0, nil
	case "errorgetupdatetime":
		return 0, errors.New("Test error")
	case "exceeded":
		return 100, nil
	case "exceeded_":
		return 0, nil
	default:
		return 0, nil
	}
}

func (mock ApiKeyMockRepository) UpdateAllQuota() {

}
