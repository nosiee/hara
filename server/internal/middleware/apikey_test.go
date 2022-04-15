package middleware

import (
	repository "hara/internal/repository/mock"
	"hara/internal/testhelpers"
	"net/http/httptest"
	"testing"
)

func TestApiKeyProvided(t *testing.T) {
	correctApiKeyCtx, correctApiKeyRec := testhelpers.CreateContext(httptest.NewRequest("POST", "http://localhost:8080/api/convert/image?key=dd87b95c-b667-11ec-b909-0242ac120002", nil), nil)
	incorrectApiKeyCtx, incorrectApiKeyRec := testhelpers.CreateContext(httptest.NewRequest("POST", "http://localhost:8080/api/convert/image?key=dd87b95c-b667", nil), nil)
	incorrectUrlQueryCtx, incorrectUrlQueryRec := testhelpers.CreateContext(httptest.NewRequest("POST", "http://localhost:8080?key=;", nil), nil)
	withoutApiKeyCtx, withoutApiKeyRec := testhelpers.CreateContext(httptest.NewRequest("POST", "http://localhost:8080/api/convert/image", nil), nil)

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

	correctApiKeyCtx, correctApiKeyRec := testhelpers.CreateContext(httptest.NewRequest("POST", "http://localhost:8080/api/convert/image?key=correct", nil), nil)
	incorrectApiKeyCtx, incorrectApiKeyRec := testhelpers.CreateContext(httptest.NewRequest("POST", "http://localhsot:8080/api/convert/image?key=incorrect", nil), nil)
	errorApiKeyCtx, errorApiKeyRec := testhelpers.CreateContext(httptest.NewRequest("POST", "http://localhsot:8080/api/convert/image?key=error", nil), nil)

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

	exceededApiKeyCtx, exceededApiKeyRec := testhelpers.CreateContext(httptest.NewRequest("POST", "http://localhsot:8080/api/convert/image?key=exceeded", nil), nil)
	exceededApiKeyCtx_, exceededApiKeyRec_ := testhelpers.CreateContext(httptest.NewRequest("POST", "http://localhsot:8080/api/convert/image?key=exceeded_", nil), nil)

	errGetUpdatetimeCtx, errGetUpdatetimeRec := testhelpers.CreateContext(httptest.NewRequest("POST", "http://localhsot:8080/api/convert/image?key=errorgetupdatetime", nil), nil)
	errIncrementQuotaCtx, errIncrementQuotaRec := testhelpers.CreateContext(httptest.NewRequest("POST", "http://localhsot:8080/api/convert/image?key=errorincrementquota", nil), nil)
	errGetQuotaCtx, errGetQuotaRec := testhelpers.CreateContext(httptest.NewRequest("POST", "http://localhsot:8080/api/convert/image?key=errorquota", nil), nil)
	errSetUpdatetimeCtx, errSetupdatetimeRec := testhelpers.CreateContext(httptest.NewRequest("POST", "http://localhsot:8080/api/convert/image?key=errorsetupdatetime", nil), nil)

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
