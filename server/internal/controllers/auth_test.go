package controllers

import (
	"hara/internal/config"
	repository "hara/internal/repository/mock"
	"hara/internal/testhelpers"
	"testing"
)

func TestSingUp(t *testing.T) {
	c := NewControllers(repository.NewUserMockRepository(), nil, nil, nil)

	correctCtx, correctRec := testhelpers.CreateContext(testhelpers.CreateFormRequest(map[string]string{
		"username": "nosiee",
		"password": "123456789",
		"email":    "nosiee@example.com",
	}), nil)

	errAddCtx, errAddRec := testhelpers.CreateContext(testhelpers.CreateFormRequest(map[string]string{
		"username": "errorAdd",
		"password": "123456789",
		"email":    "nosiee@example.com",
	}), nil)

	wrongHSKeyCtx, wrongHSKeyRec := testhelpers.CreateContext(testhelpers.CreateFormRequest(map[string]string{
		"username": "nosieee",
		"password": "123456789",
		"email":    "nosiee@example.com",
	}), nil)

	testCases := []testhelpers.ContextCase{
		testhelpers.CreateContextCase(correctCtx, correctRec, true, "CorrectForm"),
		testhelpers.CreateContextCase(errAddCtx, errAddRec, false, "ErrInAdd"),
		testhelpers.CreateContextCase(wrongHSKeyCtx, wrongHSKeyRec, false, "WrongHS512Key"),
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			if tc.Name == "WrongHS512Key" {
				config.Values.HS512Key = "invalidkey"
			}

			c.SignUp(tc.Context)
			tc.CheckCase(t)
		})
	}
}

func TestSignIn(t *testing.T) {
	if err := config.LoadFromFile("../../testdata/configs/config_test_correct.toml"); err != nil {
		t.Fatal(err)
	}

	c := NewControllers(repository.NewUserMockRepository(), nil, nil, nil)

	correctCtx, correctRec := testhelpers.CreateContext(testhelpers.CreateFormRequest(map[string]string{
		"username": "nosiee",
		"password": "123456789",
	}), nil)

	incorrectCtx, incorrectRec := testhelpers.CreateContext(testhelpers.CreateFormRequest(map[string]string{
		"username": "nosiee",
		"password": "incorrectpassword",
	}), nil)

	noSuchUserCtx, noSuchUserRec := testhelpers.CreateContext(testhelpers.CreateFormRequest(map[string]string{
		"username": "nosuchuser",
		"password": "123456789",
	}), nil)

	wrongHSKeyCtx, wrongHSKeyRec := testhelpers.CreateContext(testhelpers.CreateFormRequest(map[string]string{
		"username": "nosiee",
		"password": "123456789",
	}), nil)

	testCases := []testhelpers.ContextCase{
		testhelpers.CreateContextCase(correctCtx, correctRec, true, "CorrectForm"),
		testhelpers.CreateContextCase(incorrectCtx, incorrectRec, false, "IncorrectPassword"),
		testhelpers.CreateContextCase(noSuchUserCtx, noSuchUserRec, false, "NoSuchUser"),
		testhelpers.CreateContextCase(wrongHSKeyCtx, wrongHSKeyRec, false, "WrongHS512Key"),
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			if tc.Name == "WrongHS512Key" {
				config.Values.HS512Key = "invalidkey"
			}

			c.SignIn(tc.Context)
			tc.CheckCase(t)
		})
	}
}
