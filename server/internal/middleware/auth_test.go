package middleware

import (
	"hara/internal/config"
	"hara/internal/controllers"
	"hara/internal/testhelpers"
	"testing"
	"time"
)

func TestSignUpFormProvided(t *testing.T) {
	correctFormCtx, correctFormRec := testhelpers.CreateContext(testhelpers.CreateFormRequest(map[string]string{
		"username": "nosiee",
		"password": "123456789",
		"email":    "nosiee@example.com",
	}), nil)

	withoutUsernameCtx, withoutUsernameRec := testhelpers.CreateContext(testhelpers.CreateFormRequest(map[string]string{
		"password": "123456789",
		"email":    "nosiee@example.com",
	}), nil)

	withoutPasswordCtx, withoutPasswordRec := testhelpers.CreateContext(testhelpers.CreateFormRequest(map[string]string{
		"username": "nosiee",
		"email":    "nosiee@example.com",
	}), nil)

	withoutEmailCtx, withoutEmailRec := testhelpers.CreateContext(testhelpers.CreateFormRequest(map[string]string{
		"username": "nosiee",
		"password": "123456789",
	}), nil)

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
	correctFormCtx, correctFormRec := testhelpers.CreateContext(testhelpers.CreateFormRequest(map[string]string{
		"username": "nosiee",
		"password": "123456789",
	}), nil)

	withoutUsernameCtx, withoutUsernameRec := testhelpers.CreateContext(testhelpers.CreateFormRequest(map[string]string{
		"password": "123456789",
	}), nil)

	withoutPasswordCtx, withoutPasswordRec := testhelpers.CreateContext(testhelpers.CreateFormRequest(map[string]string{
		"username": "nosiee",
	}), nil)

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
	correctFormCtx, correctFormRec := testhelpers.CreateContext(testhelpers.CreateFormRequest(map[string]string{
		"username": "nosiee",
		"password": "123456789",
		"email":    "nosiee@example.com",
	}), nil)

	incorrectUsernameCtx, incorrectUsernameRec := testhelpers.CreateContext(testhelpers.CreateFormRequest(map[string]string{
		"username": "n",
		"password": "123456789",
		"email":    "nosiee@example.com",
	}), nil)

	incorrectPasswordCtx, incorrectPasswordRec := testhelpers.CreateContext(testhelpers.CreateFormRequest(map[string]string{
		"username": "nosiee",
		"password": "1234",
		"email":    "nosiee@example.com",
	}), nil)

	incorrectEmailCtx, incorrectEmailRec := testhelpers.CreateContext(testhelpers.CreateFormRequest(map[string]string{
		"username": "nosiee",
		"password": "123456789",
		"email":    "noseie",
	}), nil)

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
	correctFormCtx, correctFormRec := testhelpers.CreateContext(testhelpers.CreateFormRequest(map[string]string{
		"username": "nosiee",
		"password": "123456789",
	}), nil)

	incorrectUsernameCtx, incorrectUsernameRec := testhelpers.CreateContext(testhelpers.CreateFormRequest(map[string]string{
		"username": "n",
		"password": "123456789",
	}), nil)

	incorrectPasswordCtx, incorrectPasswordRec := testhelpers.CreateContext(testhelpers.CreateFormRequest(map[string]string{
		"username": "nosiee",
		"password": "1234",
	}), nil)

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

	correctTokenCtx, correctTokenRec := testhelpers.CreateContext(testhelpers.CreateCookieRequest("jwt", token), nil)
	incorrectCookieKeyCtx, incorrectCookieKeyRec := testhelpers.CreateContext(testhelpers.CreateCookieRequest("test", token), nil)
	incorrectTokenCtx, incorrectTokenRec := testhelpers.CreateContext(testhelpers.CreateCookieRequest("jwt", "definitelynotatoken"), nil)

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
