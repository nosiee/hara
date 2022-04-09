package middleware

import (
	"hara/internal/testhelpers"
	"net/http/httptest"
	"testing"
)

func TestApiKeyProvided(t *testing.T) {
	correctApiKeyCtx, correctApiKeyRec := testhelpers.CreateContextWithRequest(httptest.NewRequest("POST", "http://localhost:8080/api/convert/image?key=dd87b95c-b667-11ec-b909-0242ac120002", nil))
	incorrectApiKeyCtx, incorrectApiKeyRec := testhelpers.CreateContextWithRequest(httptest.NewRequest("POST", "http://localhost:8080/api/convert/image?key=dd87b95c-b667", nil))
	incorrectUrlQueryCtx, incorrectUrlQueryRec := testhelpers.CreateContextWithRequest(httptest.NewRequest("POST", "http://localhost:8080?key=;", nil))
	withoutApiKeyCtx, withoutApiKeyRec := testhelpers.CreateContextWithRequest(httptest.NewRequest("POST", "http://localhost:8080/api/convert/image", nil))

	testCases := []testhelpers.ContextCase{
		testhelpers.CreateContextCase(correctApiKeyCtx, correctApiKeyRec, true, "CorrectApiKey"),
		testhelpers.CreateContextCase(incorrectApiKeyCtx, incorrectApiKeyRec, false, "IncorrectApiKey"),
		testhelpers.CreateContextCase(withoutApiKeyCtx, withoutApiKeyRec, false, "WithoutApiKey"),
		testhelpers.CreateContextCase(incorrectUrlQueryCtx, incorrectUrlQueryRec, false, "WithoutApiKey"),
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			ApiKeyProvided(tc.Context)
			tc.CheckCase(t)
		})
	}
}

func TestApiKeyValidate(t *testing.T) {
	println("Not implemented yet")
}
