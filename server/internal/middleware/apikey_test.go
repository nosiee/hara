package middleware

import (
	"net/http/httptest"
	"testing"
)

func TestApiKeyProvided(t *testing.T) {
	correctApiKeyCtx, correctApiKeyRec := createContextWithRequest(httptest.NewRequest("POST", "http://localhost:8080/api/convert/image?key=dd87b95c-b667-11ec-b909-0242ac120002", nil))
	incorrectApiKeyCtx, incorrectApiKeyRec := createContextWithRequest(httptest.NewRequest("POST", "http://localhost:8080/api/convert/image?key=dd87b95c-b667", nil))
	incorrectUrlQueryCtx, incorrectUrlQueryRec := createContextWithRequest(httptest.NewRequest("POST", "http://localhost:8080?key=;", nil))
	withoutApiKeyCtx, withoutApiKeyRec := createContextWithRequest(httptest.NewRequest("POST", "http://localhost:8080/api/convert/image", nil))

	testCases := []contextCase{
		createContextCase(correctApiKeyCtx, correctApiKeyRec, true, "CorrectApiKey"),
		createContextCase(incorrectApiKeyCtx, incorrectApiKeyRec, false, "IncorrectApiKey"),
		createContextCase(withoutApiKeyCtx, withoutApiKeyRec, false, "WithoutApiKey"),
		createContextCase(incorrectUrlQueryCtx, incorrectUrlQueryRec, false, "WithoutApiKey"),
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ApiKeyProvided(tc.context)
			tc.checkCase(t)
		})
	}
}

func TestApiKeyValidate(t *testing.T) {
	// TODO: We need to find a way to use mock db instead of real postgresql
	println("Not implemented yet")
}
