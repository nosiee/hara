package middleware

import (
	"hara/internal/config"
	"hara/internal/controllers"
	"hara/internal/testhelpers"
	"testing"
	"time"
)

func TestSignUpFormProvided(t *testing.T) {
	correctFormCtx, correctFormRec := testhelpers.CreateContextWithRequest(testhelpers.CreateFormRequest(map[string]string{
		"username": "nosiee",
		"password": "123456789",
		"email":    "nosiee@example.com",
	}))

	withoutUsernameCtx, withoutUsernameRec := testhelpers.CreateContextWithRequest(testhelpers.CreateFormRequest(map[string]string{
		"password": "123456789",
		"email":    "nosiee@example.com",
	}))

	withoutPasswordCtx, withoutPasswordRec := testhelpers.CreateContextWithRequest(testhelpers.CreateFormRequest(map[string]string{
		"username": "nosiee",
		"email":    "nosiee@example.com",
	}))

	withoutEmailCtx, withoutEmailRec := testhelpers.CreateContextWithRequest(testhelpers.CreateFormRequest(map[string]string{
		"username": "nosiee",
		"password": "123456789",
	}))

	testCases := []testhelpers.ContextCase{
		testhelpers.CreateContextCase(correctFormCtx, correctFormRec, true, "CorrectForm"),
		testhelpers.CreateContextCase(withoutUsernameCtx, withoutUsernameRec, false, "WithoutUsername"),
		testhelpers.CreateContextCase(withoutPasswordCtx, withoutPasswordRec, false, "WithoutPassword"),
		testhelpers.CreateContextCase(withoutEmailCtx, withoutEmailRec, false, "WithoutEmail"),
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			SignUpFormProvided(tc.Context)
			tc.CheckCase(t)
		})
	}
}

func TestSignInFormProvided(t *testing.T) {
	correctFormCtx, correctFormRec := testhelpers.CreateContextWithRequest(testhelpers.CreateFormRequest(map[string]string{
		"username": "nosiee",
		"password": "123456789",
	}))

	withoutUsernameCtx, withoutUsernameRec := testhelpers.CreateContextWithRequest(testhelpers.CreateFormRequest(map[string]string{
		"password": "123456789",
	}))

	withoutPasswordCtx, withoutPasswordRec := testhelpers.CreateContextWithRequest(testhelpers.CreateFormRequest(map[string]string{
		"username": "nosiee",
	}))

	testCases := []testhelpers.ContextCase{
		testhelpers.CreateContextCase(correctFormCtx, correctFormRec, true, "CorrectForm"),
		testhelpers.CreateContextCase(withoutUsernameCtx, withoutUsernameRec, false, "WithoutUsername"),
		testhelpers.CreateContextCase(withoutPasswordCtx, withoutPasswordRec, false, "WithoutPassword"),
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			SignInFormProvided(tc.Context)
			tc.CheckCase(t)
		})
	}
}

func TestSignUpFormValidate(t *testing.T) {
	correctFormCtx, correctFormRec := testhelpers.CreateContextWithRequest(testhelpers.CreateFormRequest(map[string]string{
		"username": "nosiee",
		"password": "123456789",
		"email":    "nosiee@example.com",
	}))

	incorrectUsernameCtx, incorrectUsernameRec := testhelpers.CreateContextWithRequest(testhelpers.CreateFormRequest(map[string]string{
		"username": "n",
		"password": "123456789",
		"email":    "nosiee@example.com",
	}))

	incorrectPasswordCtx, incorrectPasswordRec := testhelpers.CreateContextWithRequest(testhelpers.CreateFormRequest(map[string]string{
		"username": "nosiee",
		"password": "1234",
		"email":    "nosiee@example.com",
	}))

	incorrectEmailCtx, incorrectEmailRec := testhelpers.CreateContextWithRequest(testhelpers.CreateFormRequest(map[string]string{
		"username": "nosiee",
		"password": "123456789",
		"email":    "noseie",
	}))

	testCases := []testhelpers.ContextCase{
		testhelpers.CreateContextCase(correctFormCtx, correctFormRec, true, "CorrectForm"),
		testhelpers.CreateContextCase(incorrectUsernameCtx, incorrectUsernameRec, false, "IncorrectUsername"),
		testhelpers.CreateContextCase(incorrectPasswordCtx, incorrectPasswordRec, false, "IncorrectPassword"),
		testhelpers.CreateContextCase(incorrectEmailCtx, incorrectEmailRec, false, "IncorrectEmail"),
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			SignUpFormValidate(tc.Context)
			tc.CheckCase(t)
		})
	}
}

func TestSignInFormValidate(t *testing.T) {
	correctFormCtx, correctFormRec := testhelpers.CreateContextWithRequest(testhelpers.CreateFormRequest(map[string]string{
		"username": "nosiee",
		"password": "123456789",
	}))

	incorrectUsernameCtx, incorrectUsernameRec := testhelpers.CreateContextWithRequest(testhelpers.CreateFormRequest(map[string]string{
		"username": "n",
		"password": "123456789",
	}))

	incorrectPasswordCtx, incorrectPasswordRec := testhelpers.CreateContextWithRequest(testhelpers.CreateFormRequest(map[string]string{
		"username": "nosiee",
		"password": "1234",
	}))

	testCases := []testhelpers.ContextCase{
		testhelpers.CreateContextCase(correctFormCtx, correctFormRec, true, "CorrectForm"),
		testhelpers.CreateContextCase(incorrectUsernameCtx, incorrectUsernameRec, false, "IncorrectUsername"),
		testhelpers.CreateContextCase(incorrectPasswordCtx, incorrectPasswordRec, false, "IncorrectPassword"),
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			SignInFormValidate(tc.Context)
			tc.CheckCase(t)
		})
	}
}

func TestIsAuthorized(t *testing.T) {
	config.LoadFromFile("../../testdata/configs/config_test_correct.toml")
	token, err := controllers.GenerateJWT("uuid", config.Values.HS512Key, time.Now().Add(1*365*24*time.Hour).Unix())
	if err != nil {
		t.Fatal(err)
	}

	correctTokenCtx, correctTokenRec := testhelpers.CreateContextWithRequest(testhelpers.CreateCookieRequest("jwt", token))
	incorrectCookieKeyCtx, incorrectCookieKeyRec := testhelpers.CreateContextWithRequest(testhelpers.CreateCookieRequest("test", token))
	incorrectTokenCtx, incorrectTokenRec := testhelpers.CreateContextWithRequest(testhelpers.CreateCookieRequest("jwt", "definitelynotatoken"))

	testCases := []testhelpers.ContextCase{
		testhelpers.CreateContextCase(correctTokenCtx, correctTokenRec, true, "CorrectToken"),
		testhelpers.CreateContextCase(incorrectCookieKeyCtx, incorrectCookieKeyRec, false, "IncorrectCookieKey"),
		testhelpers.CreateContextCase(incorrectTokenCtx, incorrectTokenRec, false, "IncorrectToken"),
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			IsAuthorized(tc.Context)
			tc.CheckCase(t)
		})
	}
}
