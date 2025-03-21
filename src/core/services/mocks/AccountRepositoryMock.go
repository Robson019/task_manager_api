// Code generated by MockGen. DO NOT EDIT.
// Source: ./src/core/interfaces/secondary/AccountLoader.go
//
// Generated by this command:
//
//	mockgen -source=./src/core/interfaces/secondary/AccountLoader.go -destination=./src/core/services/mocks/AccountRepositoryMock.go
//

// Package mock_secondary is a generated GoMock package.
package mock_secondary

import (
	reflect "reflect"
	account "task_manager/src/core/domain/account"
	errors "task_manager/src/core/errors"

	uuid "github.com/google/uuid"
	gomock "go.uber.org/mock/gomock"
)

// MockAccountLoader is a mock of AccountLoader interface.
type MockAccountLoader struct {
	ctrl     *gomock.Controller
	recorder *MockAccountLoaderMockRecorder
}

// MockAccountLoaderMockRecorder is the mock recorder for MockAccountLoader.
type MockAccountLoaderMockRecorder struct {
	mock *MockAccountLoader
}

// NewMockAccountLoader creates a new mock instance.
func NewMockAccountLoader(ctrl *gomock.Controller) *MockAccountLoader {
	mock := &MockAccountLoader{ctrl: ctrl}
	mock.recorder = &MockAccountLoaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAccountLoader) EXPECT() *MockAccountLoaderMockRecorder {
	return m.recorder
}

// FindProfileByID mocks base method.
func (m *MockAccountLoader) FindProfileByID(accountID uuid.UUID) (*account.Account, errors.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindProfileByID", accountID)
	ret0, _ := ret[0].(*account.Account)
	ret1, _ := ret[1].(errors.Error)
	return ret0, ret1
}

// FindProfileByID indicates an expected call of FindProfileByID.
func (mr *MockAccountLoaderMockRecorder) FindProfileByID(accountID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindProfileByID", reflect.TypeOf((*MockAccountLoader)(nil).FindProfileByID), accountID)
}
