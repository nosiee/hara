package models

type ApiKey struct {
	OwnerUUID string
	Key       string
	MaxQuotas uint
	Quotas    uint
}

type ApiKeyRepository interface {
	Add(key ApiKey) error
	UserHaveKey(uuid string) (bool, error)
	IsExists(key string) (bool, error)
}

func NewApiKey(ownerUUID, key string, maxquotas, quotas uint) ApiKey {
	return ApiKey{
		ownerUUID,
		key,
		maxquotas,
		quotas,
	}
}
