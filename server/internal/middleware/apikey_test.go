package middleware

import (
	repository "hara/internal/repository/mock"
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
	mock := repository.NewApiKeyMockRepository()
	f := ApiKeyValidate(mock)

	correctApiKeyCtx, correctApiKeyRec := testhelpers.CreateContextWithRequest(httptest.NewRequest("POST", "http://localhost:8080/api/convert/image?key=correct", nil))
	incorrectApiKeyCtx, incorrectApiKeyRec := testhelpers.CreateContextWithRequest(httptest.NewRequest("POST", "http://localhsot:8080/api/convert/image?key=incorrect", nil))
	errorApiKeyCtx, errorApiKeyRec := testhelpers.CreateContextWithRequest(httptest.NewRequest("POST", "http://localhsot:8080/api/convert/image?key=error", nil))

	testCases := []testhelpers.ContextCase{
		testhelpers.CreateContextCase(correctApiKeyCtx, correctApiKeyRec, true, "CorrectApiKey"),
		testhelpers.CreateContextCase(incorrectApiKeyCtx, incorrectApiKeyRec, false, "IncorrectApiKey"),
		testhelpers.CreateContextCase(errorApiKeyCtx, errorApiKeyRec, false, "ErrorApiKey"),
	}

	for _, tc := range testCases {
		f(tc.Context)
		tc.CheckCase(t)
	}
}

func TestApiKeyQuota(t *testing.T) {
	mock := repository.NewApiKeyMockRepository()
	f := ApiKeyQuota(mock)

	exceededApiKeyCtx, exceededApiKeyRec := testhelpers.CreateContextWithRequest(httptest.NewRequest("POST", "http://localhsot:8080/api/convert/image?key=exceeded", nil))
	exceededApiKeyCtx_, exceededApiKeyRec_ := testhelpers.CreateContextWithRequest(httptest.NewRequest("POST", "http://localhsot:8080/api/convert/image?key=exceeded_", nil))

	errGetUpdatetimeCtx, errGetUpdatetimeRec := testhelpers.CreateContextWithRequest(httptest.NewRequest("POST", "http://localhsot:8080/api/convert/image?key=errorgetupdatetime", nil))
	errIncrementQuotaCtx, errIncrementQuotaRec := testhelpers.CreateContextWithRequest(httptest.NewRequest("POST", "http://localhsot:8080/api/convert/image?key=errorincrementquota", nil))
	errGetQuotaCtx, errGetQuotaRec := testhelpers.CreateContextWithRequest(httptest.NewRequest("POST", "http://localhsot:8080/api/convert/image?key=errorquota", nil))
	errSetUpdatetimeCtx, errSetupdatetimeRec := testhelpers.CreateContextWithRequest(httptest.NewRequest("POST", "http://localhsot:8080/api/convert/image?key=errorsetupdatetime", nil))

	testCases := []testhelpers.ContextCase{
		testhelpers.CreateContextCase(errGetUpdatetimeCtx, errGetUpdatetimeRec, false, "ErrorGetUpdatetime"),
		testhelpers.CreateContextCase(errGetQuotaCtx, errGetQuotaRec, false, "ErrorGetQuota"),
		testhelpers.CreateContextCase(exceededApiKeyCtx, exceededApiKeyRec, false, "ExceededQuota"),
		testhelpers.CreateContextCase(errIncrementQuotaCtx, errIncrementQuotaRec, false, "ErrorIncrementQuota"),
		testhelpers.CreateContextCase(errSetUpdatetimeCtx, errSetupdatetimeRec, false, "ErrorSetUpdatetime"),
		testhelpers.CreateContextCase(exceededApiKeyCtx_, exceededApiKeyRec_, false, "ExceededQuota#2"),
	}

	for _, tc := range testCases {
		f(tc.Context)
		tc.CheckCase(t)
	}
}
