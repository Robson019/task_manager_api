package handlers

import (
	nativeErrors "errors"
	"net/http"
	"net/http/httptest"
	mocks "task_manager/src/api/handlers/mocks"
	"task_manager/src/core/domain/account"
	"task_manager/src/core/domain/account/role"
	"task_manager/src/core/errors"
	"task_manager/src/core/messages"
	"task_manager/src/utils/tests/filenames"
	"task_manager/src/utils/tests/functions"
	md "task_manager/src/utils/tests/mockData"
	"task_manager/src/utils/tests/routes"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestAccountHandlers_FindProfile(t *testing.T) {
	testCases := findProfileTestCases{}
	t.Run("Test FindProfile() when a profile is returned",
		testCases.testWhenAProfileIsReturned)
	t.Run("Test FindProfile() when an error is returned",
		testCases.testWhenAnErrorIsReturned)
}

type findProfileTestCases struct{}

func (findProfileTestCases) testWhenAProfileIsReturned(t *testing.T) {
	roleBuilder := role.NewBuilder()
	roleBuilder.WithID(md.RoleID).WithName(md.RoleName)
	roleInstance, validationError := roleBuilder.Build()
	assert.Nil(t, validationError)

	accountBuilder := account.NewBuilder()
	accountInstance, validationErr := accountBuilder.WithID(md.AccountID).
		WithEmail(md.AccountEmail).WithRole(*roleInstance).Build()
	assert.Nil(t, validationErr)

	ctrl := gomock.NewController(t)
	accountServicesMock := mocks.NewMockAccountManager(ctrl)
	accountServicesMock.EXPECT().FetchProfileByID(gomock.Any()).Return(accountInstance, nil)

	clientRequest := httptest.NewRequest(http.MethodGet, routes.FindProfile, nil)
	clientRequest.Header.Set(echo.HeaderAuthorization, md.Token)
	recorder := httptest.NewRecorder()
	e := echo.New()
	context := e.NewContext(clientRequest, recorder)

	accountHandlers := NewAccountHandlers(accountServicesMock)
	_ = accountHandlers.FindProfile(context)

	expectedHTTPStatusCode := http.StatusOK
	expectedJSON := functions.ReadJSON(filenames.FindProfile)

	assert.Equal(t, expectedHTTPStatusCode, recorder.Code, "The status codes does not match")
	assert.JSONEq(t, expectedJSON, recorder.Body.String(), "The JSONs does not match")
}

func (findProfileTestCases) testWhenAnErrorIsReturned(t *testing.T) {
	internalError := nativeErrors.New(md.ConnectionError)
	returnedError := errors.NewUnexpectedError(messages.UnexpectedErrorMessage, internalError)

	ctrl := gomock.NewController(t)
	accountServicesMock := mocks.NewMockAccountManager(ctrl)
	accountServicesMock.EXPECT().FetchProfileByID(gomock.Any()).Return(nil, returnedError)

	clientRequest := httptest.NewRequest(http.MethodGet, routes.FindProfile, nil)
	clientRequest.Header.Set(echo.HeaderAuthorization, md.Token)
	recorder := httptest.NewRecorder()
	e := echo.New()
	context := e.NewContext(clientRequest, recorder)

	accountHandlers := NewAccountHandlers(accountServicesMock)
	_ = accountHandlers.FindProfile(context)

	expectedHTTPStatusCode := http.StatusInternalServerError
	expectedJSON := functions.ReadJSON(filenames.AnUnexpectedErrorOccurred)

	assert.Equal(t, expectedHTTPStatusCode, recorder.Code, "The status codes does not match")
	assert.JSONEq(t, expectedJSON, recorder.Body.String(), "The JSONs does not match")
}
