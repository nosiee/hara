package models

type ApiKey struct {
	OwnerUUID  string
	Key        string
	MaxQuota   uint
	Quota      uint
	Updatetime int64
}

type ApiKeyRepository interface {
	Add(key ApiKey) error
	ChangeKeyID(uuid, key string) error
	UserHasKey(uuid string) (bool, error)
	IsExists(key string) (bool, error)
	GetKey(uuid string) (*ApiKey, error)
	GetQuota(key string) (uint, uint, error)
	IncrementQuota(key string) error
	SetUpdatetime(key string, updatetime int64) error
	GetUpdatetime(key string) (int64, error)
	UpdateAllQuota()
}

func NewApiKey(ownerUUID, key string, maxquota, quota uint, updatetime int64) ApiKey {
	return ApiKey{
		ownerUUID,
		key,
		maxquota,
		quota,
		updatetime,
	}
}
