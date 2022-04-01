package db

func CreateNewUser(username, hash, email string) error {
	_, err := db.Exec("INSERT INTO users(username, hash, email) VALUES($1, $2, $3)", username, hash, email)
	return err
}
