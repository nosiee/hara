package db

func AddNewApiKey(uuid, key string, quotas int) error {
	_, err := db.Exec("INSERT INTO apikeys(owneruuid, key, quotas) VALUES($1, $2, $3)", uuid, key, quotas)
	return err
}

func UserHasKey(uuid string) (bool, error) {
	var ID int

	err := db.QueryRow("SELECT id FROM apikeys WHERE owneruuid=$1", uuid).Scan(&ID)
	return ID != 0, err
}
