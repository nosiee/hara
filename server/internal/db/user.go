package db

func CreateNewUser(uuid, username, hash, email string) error {
	_, err := db.Exec("INSERT INTO users(uuid, username, hash, email) VALUES($1, $2, $3, $4)", uuid, username, hash, email)
	return err
}

func FindUser(username string) (string, string, bool) {
	var ID int
	var uuid, hash string

	db.QueryRow("SELECT id, uuid, hash FROM users WHERE username=$1", username).Scan(&ID, &uuid, &hash)
	return uuid, hash, ID != 0
}
