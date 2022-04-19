package controllers

import (
	repository "hara/internal/repository/mock"
	"hara/internal/testhelpers"
	"net/http/httptest"
	"testing"
)

func TestGetApiKey(t *testing.T) {
	c := NewControllers(nil, nil, repository.NewApiKeyMockRepository(), nil)

	correctCtx, correctRec := testhelpers.CreateContext(httptest.NewRequest("POST", "/", nil), map[string]any{
		"uuid": "correct",
	})
	hasKeyCtx, hasKeyRec := testhelpers.CreateContext(httptest.NewRequest("POST", "/", nil), map[string]any{
		"uuid": "has",
	})
	errUserHasKeyCtx, errUserHasKeyRec := testhelpers.CreateContext(httptest.NewRequest("POST", "/", nil), map[string]any{
		"uuid": "errorUserHasKey",
	})
	errAddCtx, errAddRec := testhelpers.CreateContext(httptest.NewRequest("POST", "/", nil), map[string]any{
		"uuid": "errorAdd",
	})

	testCases := []testhelpers.ContextCase{
		testhelpers.CreateContextCase(correctCtx, correctRec, true, "Correct"),
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

func TestResetApiKey(t *testing.T) {
	c := NewControllers(nil, nil, repository.NewApiKeyMockRepository(), nil)

	correctCtx, correctRec := testhelpers.CreateContext(httptest.NewRequest("POST", "/", nil), map[string]any{
		"uuid": "has",
	})

	errUserHasKeyCtx, errUserHasKeyRec := testhelpers.CreateContext(httptest.NewRequest("POST", "/", nil), map[string]any{
		"uuid": "errorUserHasKey",
	})

	errAddCtx, errAddRec := testhelpers.CreateContext(httptest.NewRequest("POST", "/", nil), map[string]any{
		"uuid": "errorAdd",
	})

	errChangeKeyCtx, errChangeKeyRec := testhelpers.CreateContext(httptest.NewRequest("POST", "/", nil), map[string]any{
		"uuid": "errorChangeKeyID",
	})

	testCases := []testhelpers.ContextCase{
		testhelpers.CreateContextCase(correctCtx, correctRec, true, "Correct"),
		testhelpers.CreateContextCase(errUserHasKeyCtx, errUserHasKeyRec, false, "ErrorInUserHasKey"),
		testhelpers.CreateContextCase(errAddCtx, errAddRec, false, "ErrorInApiKeyRepositoryAdd"),
		testhelpers.CreateContextCase(errChangeKeyCtx, errChangeKeyRec, false, "ErrorInChangeKeyID"),
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			c.ResetApiKey(tc.Context)
			tc.CheckCase(t)
		})
	}
}

func TestRemindApiKey(t *testing.T) {
	c := NewControllers(nil, nil, repository.NewApiKeyMockRepository(), nil)

	correctCtx, correctRec := testhelpers.CreateContext(httptest.NewRequest("POST", "/", nil), map[string]any{
		"uuid": "has",
	})

	errUserHasKeyCtx, errUserHasKeyRec := testhelpers.CreateContext(httptest.NewRequest("POST", "/", nil), map[string]any{
		"uuid": "errorUserHasKey",
	})

	hasNoKeyCtx, hasNoKeyRec := testhelpers.CreateContext(httptest.NewRequest("POST", "/", nil), map[string]any{
		"uuid": "hasNoKey",
	})

	errGetKeyCtx, errGetKeyRec := testhelpers.CreateContext(httptest.NewRequest("POST", "/", nil), map[string]any{
		"uuid": "errorGetKey",
	})

	testCases := []testhelpers.ContextCase{
		testhelpers.CreateContextCase(correctCtx, correctRec, true, "Correct"),
		testhelpers.CreateContextCase(errUserHasKeyCtx, errUserHasKeyRec, false, "ErrorInUserHasKey"),
		testhelpers.CreateContextCase(hasNoKeyCtx, hasNoKeyRec, false, "hasNoKey"),
		testhelpers.CreateContextCase(errGetKeyCtx, errGetKeyRec, false, "errorInGetKey"),
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			c.RemindApiKey(tc.Context)
			tc.CheckCase(t)
		})
	}
}
