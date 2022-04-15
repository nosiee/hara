package controllers

import (
	"hara/internal/config"
	repository "hara/internal/repository/mock"
	"hara/internal/testhelpers"
	"testing"
	"time"
)

func TestGetApiKey(t *testing.T) {
	c := NewControllers(nil, nil, repository.NewApiKeyMockRepository(), nil)

	if err := config.LoadFromFile("../../testdata/configs/config_test_correct.toml"); err != nil {
		t.Fatal(err)
	}

	correctToken, err := GenerateJWT("correct", config.Values.HS512Key, time.Now().Add(time.Hour).Unix())
	if err != nil {
		t.Fatal(err)
	}

	hasKeyToken, err := GenerateJWT("has", config.Values.HS512Key, time.Now().Add(time.Hour).Unix())
	if err != nil {
		t.Fatal(err)
	}

	errUserHasKeyToken, err := GenerateJWT("errorUserHasKey", config.Values.HS512Key, time.Now().Add(time.Hour).Unix())
	if err != nil {
		t.Fatal(err)
	}

	errAddToken, err := GenerateJWT("errorAdd", config.Values.HS512Key, time.Now().Add(time.Hour).Unix())
	if err != nil {
		t.Fatal(err)
	}

	correctJwtCtx, correctJwtRec := testhelpers.CreateContext(testhelpers.CreateCookieRequest("jwt", correctToken), nil)
	hasKeyCtx, hasKeyRec := testhelpers.CreateContext(testhelpers.CreateCookieRequest("jwt", hasKeyToken), nil)
	incorrectJwtCtx, incorrectJwtRec := testhelpers.CreateContext(testhelpers.CreateCookieRequest("jwt", "definitelynotatoken"), nil)
	errUserHasKeyCtx, errUserHasKeyRec := testhelpers.CreateContext(testhelpers.CreateCookieRequest("jwt", errUserHasKeyToken), nil)
	errAddCtx, errAddRec := testhelpers.CreateContext(testhelpers.CreateCookieRequest("jwt", errAddToken), nil)

	testCases := []testhelpers.ContextCase{
		testhelpers.CreateContextCase(correctJwtCtx, correctJwtRec, true, "CorrectJWT"),
		testhelpers.CreateContextCase(incorrectJwtCtx, incorrectJwtRec, false, "IncorrectJWT"),
		testhelpers.CreateContextCase(hasKeyCtx, hasKeyRec, false, "UserAlreadyHasKey"),
		testhelpers.CreateContextCase(errUserHasKeyCtx, errUserHasKeyRec, false, "ErrorInUserHasKey"),
		testhelpers.CreateContextCase(errAddCtx, errAddRec, false, "ErrorInApiKeyRepositoryAdd"),
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			c.GetApiKey(tc.Context)
			tc.CheckCase(t)
		})
	}
}
