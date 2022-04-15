package models

type User struct {
	UUID     string
	Username string
	Hash     string
	Email    string
}

type UserRepository interface {
	Add(user User) error
	FindByUsername(username string) *User
}

func NewUser(uuid, username, hash, email string) User {
	return User{
		uuid,
		username,
		hash,
		email,
	}
}
