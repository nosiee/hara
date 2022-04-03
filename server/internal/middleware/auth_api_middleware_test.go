package middleware

import (
	"fmt"
	"hara/internal/config"
	"hara/internal/controllers"
	"testing"
)

func TestAuthFormProvided(t *testing.T) {
	testCases := []requestCase{
		{
			createFormRequest(map[string]string{
				"username": "nosiee",
				"password": "123456789",
				"email":    "nosiee@example.com",
			}, "/api/auth/signup"),
			true,
			"CorrectSignUpForm",
		},
		{
			createFormRequest(map[string]string{
				"password": "123456789",
			}, "/api/auth/signin"),
			false,
			"IncorrectSignInForm",
		},
		{
			createFormRequest(map[string]string{
				"username": "nosiee",
				"password": "123456789",
			}, "/api/auth/signup"),
			false,
			"IncorrectSignUpForm",
		},
		{
			createFormRequest(map[string]string{
				"username": "nosiee",
				"password": "123456789",
			}, "/api/auth/signin"),
			true,
			"CorrectSignInForm",
		},
		{
			createFormRequest(map[string]string{
				"username": "nosiee",
			}, "/api/auth/signin"),
			false,
			"IncorrectSignInForm",
		},
	}

	for _, c := range testCases {
		t.Run(c.name, func(t *testing.T) {
			testRequest(t, c, AuthFormProvided)
		})
	}
}

func TestAuthFormValidate(t *testing.T) {
	testCases := []requestCase{
		{
			createFormRequest(map[string]string{
				"username": "nosiee",
				"password": "123456789",
				"email":    "nosiee@example.com",
			}, "/api/auth/signup"),
			true,
			"CorrectSignUpForm",
		},
		{
			createFormRequest(map[string]string{
				"username": "nos",
				"password": "123456789",
				"email":    "nosiee@example.com",
			}, "/api/auth/signin"),
			false,
			"IncorrectSignInForm",
		},
		{
			createFormRequest(map[string]string{
				"username": "nosiee",
				"password": "1234",
				"email":    "nosiee@example.com",
			}, "/api/auth/signin"),
			false,
			"IncorrectSignInForm",
		},
		{
			createFormRequest(map[string]string{
				"username": "nosiee",
				"password": "123456789",
				"email":    "nosiee",
			}, "/api/auth/signup"),
			false,
			"IncorrectSignInForm",
		},
	}

	for _, c := range testCases {
		t.Run(c.name, func(t *testing.T) {
			testRequest(t, c, AuthFormValidate)
		})
	}
}

func TestIsAutohorized(t *testing.T) {
	config.LoadFromFile("../../testdata/configs/config_test_correct.toml")
	token, err := controllers.GenerateJWT("uuid", config.Values.JWTKey)
	if err != nil {
		t.Fatal(err)
	}

	testCases := []requestCase{
		{
			createCookieRequest("GET", fmt.Sprintf("jwt=%s", token)),
			true,
			"Autohorized",
		},
		{
			createCookieRequest("GET", ""),
			false,
			"Unauthorized",
		},
		{
			createCookieRequest("GET", fmt.Sprintf("jwt=invalid_token")),
			false,
			"InvalidToken",
		},
	}

	for _, c := range testCases {
		t.Run(c.name, func(t *testing.T) {
			testRequest(t, c, IsAuthorized)
		})
	}
}
