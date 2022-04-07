package middleware

import (
	"hara/internal/config"
	"hara/internal/controllers"
	"testing"
)

func TestSignUpFormProvided(t *testing.T) {
	correctFormCtx, correctFormRec := createContextWithRequest(createFormRequest(map[string]string{
		"username": "nosiee",
		"password": "123456789",
		"email":    "nosiee@example.com",
	}))

	withoutUsernameCtx, withoutUsernameRec := createContextWithRequest(createFormRequest(map[string]string{
		"password": "123456789",
		"email":    "nosiee@example.com",
	}))

	withoutPasswordCtx, withoutPasswordRec := createContextWithRequest(createFormRequest(map[string]string{
		"username": "nosiee",
		"email":    "nosiee@example.com",
	}))

	withoutEmailCtx, withoutEmailRec := createContextWithRequest(createFormRequest(map[string]string{
		"username": "nosiee",
		"password": "123456789",
	}))

	testCases := []contextCase{
		createContextCase(correctFormCtx, correctFormRec, true, "CorrectForm"),
		createContextCase(withoutUsernameCtx, withoutUsernameRec, false, "WithoutUsername"),
		createContextCase(withoutPasswordCtx, withoutPasswordRec, false, "WithoutPassword"),
		createContextCase(withoutEmailCtx, withoutEmailRec, false, "WithoutEmail"),
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			SignUpFormProvided(tc.context)
			tc.checkCase(t)
		})
	}
}

func TestSignInFormProvided(t *testing.T) {
	correctFormCtx, correctFormRec := createContextWithRequest(createFormRequest(map[string]string{
		"username": "nosiee",
		"password": "123456789",
	}))

	withoutUsernameCtx, withoutUsernameRec := createContextWithRequest(createFormRequest(map[string]string{
		"password": "123456789",
	}))

	withoutPasswordCtx, withoutPasswordRec := createContextWithRequest(createFormRequest(map[string]string{
		"username": "nosiee",
	}))

	testCases := []contextCase{
		createContextCase(correctFormCtx, correctFormRec, true, "CorrectForm"),
		createContextCase(withoutUsernameCtx, withoutUsernameRec, false, "WithoutUsername"),
		createContextCase(withoutPasswordCtx, withoutPasswordRec, false, "WithoutPassword"),
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			SignInFormProvided(tc.context)
			tc.checkCase(t)
		})
	}
}

func TestSignUpFormValidate(t *testing.T) {
	correctFormCtx, correctFormRec := createContextWithRequest(createFormRequest(map[string]string{
		"username": "nosiee",
		"password": "123456789",
		"email":    "nosiee@example.com",
	}))

	incorrectUsernameCtx, incorrectUsernameRec := createContextWithRequest(createFormRequest(map[string]string{
		"username": "n",
		"password": "123456789",
		"email":    "nosiee@example.com",
	}))

	incorrectPasswordCtx, incorrectPasswordRec := createContextWithRequest(createFormRequest(map[string]string{
		"username": "nosiee",
		"password": "1234",
		"email":    "nosiee@example.com",
	}))

	incorrectEmailCtx, incorrectEmailRec := createContextWithRequest(createFormRequest(map[string]string{
		"username": "nosiee",
		"password": "123456789",
		"email":    "noseie",
	}))

	testCases := []contextCase{
		createContextCase(correctFormCtx, correctFormRec, true, "CorrectForm"),
		createContextCase(incorrectUsernameCtx, incorrectUsernameRec, false, "IncorrectUsername"),
		createContextCase(incorrectPasswordCtx, incorrectPasswordRec, false, "IncorrectPassword"),
		createContextCase(incorrectEmailCtx, incorrectEmailRec, false, "IncorrectEmail"),
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			SignUpFormValidate(tc.context)
			tc.checkCase(t)
		})
	}
}

func TestSignInFormValidate(t *testing.T) {
	correctFormCtx, correctFormRec := createContextWithRequest(createFormRequest(map[string]string{
		"username": "nosiee",
		"password": "123456789",
	}))

	incorrectUsernameCtx, incorrectUsernameRec := createContextWithRequest(createFormRequest(map[string]string{
		"username": "n",
		"password": "123456789",
	}))

	incorrectPasswordCtx, incorrectPasswordRec := createContextWithRequest(createFormRequest(map[string]string{
		"username": "nosiee",
		"password": "1234",
	}))

	testCases := []contextCase{
		createContextCase(correctFormCtx, correctFormRec, true, "CorrectForm"),
		createContextCase(incorrectUsernameCtx, incorrectUsernameRec, false, "IncorrectUsername"),
		createContextCase(incorrectPasswordCtx, incorrectPasswordRec, false, "IncorrectPassword"),
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			SignInFormValidate(tc.context)
			tc.checkCase(t)
		})
	}
}

func TestIsAuthorized(t *testing.T) {
	config.LoadFromFile("../../testdata/configs/config_test_correct.toml")
	token, err := controllers.GenerateJWT("uuid", config.Values.HS512Key)
	if err != nil {
		t.Fatal(err)
	}

	correctTokenCtx, correctTokenRec := createContextWithRequest(createCookieRequest("jwt", token))
	incorrectCookieKeyCtx, incorrectCookieKeyRec := createContextWithRequest(createCookieRequest("test", token))
	incorrectTokenCtx, incorrectTokenRec := createContextWithRequest(createCookieRequest("jwt", "definitelynotatoken"))

	testCases := []contextCase{
		createContextCase(correctTokenCtx, correctTokenRec, true, "CorrectToken"),
		createContextCase(incorrectCookieKeyCtx, incorrectCookieKeyRec, false, "IncorrectCookieKey"),
		createContextCase(incorrectTokenCtx, incorrectTokenRec, false, "IncorrectToken"),
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			IsAuthorized(tc.context)
			tc.checkCase(t)
		})
	}
}
