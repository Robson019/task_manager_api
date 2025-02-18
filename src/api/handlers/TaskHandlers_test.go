package handlers

import (
	"encoding/json"
	nativeErrors "errors"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	mocks "task_manager/src/api/handlers/mocks"
	"task_manager/src/core/domain/task"
	"task_manager/src/core/errors"
	"task_manager/src/core/messages"
	"task_manager/src/utils/tests/filenames"
	"task_manager/src/utils/tests/functions"
	md "task_manager/src/utils/tests/mockData"
	"task_manager/src/utils/tests/routes"

	"go.uber.org/mock/gomock"
	"testing"
)

func TestTaskHandlers_CreateTask(t *testing.T) {
	testCases := createTaskTestCases{}
	t.Run("Test CreateTask() when task is successfully created",
		testCases.testWhenCreateTaskIsSuccessfullyCreated)
	t.Run("Test CreateTask() when an error is returned",
		testCases.testWhenCreateTaskReturnsAnError)
}

type createTaskTestCases struct{}

func (createTaskTestCases) testWhenCreateTaskIsSuccessfullyCreated(t *testing.T) {
	ctrl := gomock.NewController(t)
	taskServicesMock := mocks.NewMockTaskManager(ctrl)
	taskServicesMock.EXPECT().CreateTask(gomock.Any()).Return(&md.TaskID, nil)

	var taskValues map[string]interface{}
	jsonError := json.Unmarshal([]byte(md.TaskValuesValid), &taskValues)
	assert.Nil(t, jsonError)

	requestBody := strings.NewReader(md.TaskValuesValid)
	clientRequest := httptest.NewRequest(http.MethodPost, routes.CreateTask, requestBody)
	clientRequest.Header.Set(echo.HeaderAuthorization, md.Token)
	clientRequest.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	recorder := httptest.NewRecorder()
	e := echo.New()
	context := e.NewContext(clientRequest, recorder)

	taskHandlers := NewTaskHandlers(taskServicesMock)
	_ = taskHandlers.CreateTask(context)

	expectedHTTPStatusCode := http.StatusCreated
	expectedJSON := functions.ReadJSON(filenames.TaskCreatedSuccess)

	assert.Equal(t, expectedHTTPStatusCode, recorder.Code, md.TheStatusCodesDoesNotMatch)
	assert.JSONEq(t, expectedJSON, recorder.Body.String(), md.TheJSONsDoesNotMatch)
}

func (createTaskTestCases) testWhenCreateTaskReturnsAnError(t *testing.T) {
	internalError := nativeErrors.New(md.ConnectionError)
	returnedError := errors.NewUnexpectedError(messages.UnexpectedErrorMessage, internalError)

	ctrl := gomock.NewController(t)
	taskServicesMock := mocks.NewMockTaskManager(ctrl)
	taskServicesMock.EXPECT().CreateTask(gomock.Any()).Return(nil, returnedError)

	var taskValues map[string]interface{}
	jsonError := json.Unmarshal([]byte(md.TaskValuesValid), &taskValues)
	assert.Nil(t, jsonError)

	requestBody := strings.NewReader(md.TaskValuesValid)
	clientRequest := httptest.NewRequest(http.MethodPost, routes.CreateTask, requestBody)
	clientRequest.Header.Set(echo.HeaderAuthorization, md.Token)
	clientRequest.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	recorder := httptest.NewRecorder()
	e := echo.New()
	context := e.NewContext(clientRequest, recorder)

	taskHandlers := NewTaskHandlers(taskServicesMock)
	_ = taskHandlers.CreateTask(context)

	expectedHTTPStatusCode := http.StatusInternalServerError
	expectedJSON := functions.ReadJSON(filenames.AnUnexpectedErrorOccurred)

	assert.Equal(t, expectedHTTPStatusCode, recorder.Code, md.TheStatusCodesDoesNotMatch)
	assert.JSONEq(t, expectedJSON, recorder.Body.String(), md.TheJSONsDoesNotMatch)
}

func TestTaskHandlers_FindTasks(t *testing.T) {
	testCases := findTasksTestCases{}
	t.Run("Test FindTask() when task is successfully returned",
		testCases.testWhenFindTasksIsSuccessfullyReturned)
	t.Run("Test FindTask() when an error is returned",
		testCases.testWhenFindTasksReturnsAnError)
}

type findTasksTestCases struct{}

func (findTasksTestCases) testWhenFindTasksIsSuccessfullyReturned(t *testing.T) {
	taskBuilder := task.NewBuilder()
	_task, validationErr := taskBuilder.WithID(md.TaskID).WithTitle(md.TaskTitle).WithDescription(md.TaskDescription).
		WithStatus(md.TaskStatus).WithCreatedAt(md.TaskCreatedAt).WithUpdatedAt(md.TaskUpdatedAt).Build()
	assert.Nil(t, validationErr)

	tasks := []task.Task{*_task}

	ctrl := gomock.NewController(t)
	taskServicesMock := mocks.NewMockTaskManager(ctrl)
	taskServicesMock.EXPECT().FindTasks().Return(tasks, nil)

	clientRequest := httptest.NewRequest(http.MethodGet, routes.ListTask, nil)
	recorder := httptest.NewRecorder()
	e := echo.New()
	context := e.NewContext(clientRequest, recorder)

	taskHandlers := NewTaskHandlers(taskServicesMock)
	_ = taskHandlers.FindTasks(context)

	expectedHTTPStatusCode := http.StatusOK
	expectedJSON := functions.ReadJSON(filenames.GetTasks)

	assert.Equal(t, expectedHTTPStatusCode, recorder.Code, md.TheStatusCodesDoesNotMatch)
	assert.JSONEq(t, expectedJSON, recorder.Body.String(), md.TheJSONsDoesNotMatch)
}

func (findTasksTestCases) testWhenFindTasksReturnsAnError(t *testing.T) {
	internalError := nativeErrors.New(md.ConnectionError)
	returnedError := errors.NewUnexpectedError(messages.UnexpectedErrorMessage, internalError)

	ctrl := gomock.NewController(t)
	taskServicesMock := mocks.NewMockTaskManager(ctrl)
	taskServicesMock.EXPECT().FindTasks().Return(nil, returnedError)

	clientRequest := httptest.NewRequest(http.MethodGet, routes.ListTask, nil)
	recorder := httptest.NewRecorder()
	e := echo.New()
	context := e.NewContext(clientRequest, recorder)

	taskHandlers := NewTaskHandlers(taskServicesMock)
	_ = taskHandlers.FindTasks(context)

	expectedHTTPStatusCode := http.StatusInternalServerError
	expectedJSON := functions.ReadJSON(filenames.AnUnexpectedErrorOccurred)

	assert.Equal(t, expectedHTTPStatusCode, recorder.Code, md.TheStatusCodesDoesNotMatch)
	assert.JSONEq(t, expectedJSON, recorder.Body.String(), md.TheJSONsDoesNotMatch)
}

//TODO: Criar testes de getByID
//TODO: Criar testes de update
//TODO: Criar testes de delete
