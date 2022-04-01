package middleware

import (
	"errors"
	"regexp"
)

const (
	minUsernameLength = 4
	maxUsernameLength = 20

	minPasswordLenght = 8
	maxPasswordLength = 32
)

var emailRegex = regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)

var (
	ErrUsernameLength = errors.New("username doesn't match required length")
	ErrPasswordLength = errors.New("password doesn't match required length")
	ErrEmailRegex     = errors.New("email doesn't match requeired format")
)
