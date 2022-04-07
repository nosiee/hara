package db

func AddNewApiKey(uuid, key string, quotas int) error {
	_, err := db.Exec("INSERT INTO apikeys(uuid, key, maxquotas, quotas) VALUES($1, $2, $3, $4)", uuid, key, quotas, 0)
	return err
}

func UserHasKey(uuid string) (string, bool, error) {
	var ID int
	var key string

	err := db.QueryRow("SELECT id, key FROM apikeys WHERE uuid=$1", uuid).Scan(&ID, &key)
	return key, ID != 0, err
}

func IsKeyExists(key string) (bool, error) {
	var ID int

	err := db.QueryRow("SELECT id FROM apikeys WHERE key=$1", key).Scan(&ID)
	return ID != 0, err
}

func GetKeyQuotas(key string) (int, int, error) {
	var maxquotas, quotas int

	err := db.QueryRow("SELECT maxquotas, quotas FROM apikeys WHERE key=$1", key).Scan(&maxquotas, &quotas)
	return maxquotas, quotas, err
}

func UpdateKeyQuotas(key string, quotas int) error {
	_, err := db.Exec("UPDATE apikeys SET quotas=$1 WHERE key=$2", quotas, key)
	return err
}
