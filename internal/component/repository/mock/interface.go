// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockCreateUser is a mock of CreateUser interface.
type MockCreateUser struct {
	ctrl     *gomock.Controller
	recorder *MockCreateUserMockRecorder
}

// MockCreateUserMockRecorder is the mock recorder for MockCreateUser.
type MockCreateUserMockRecorder struct {
	mock *MockCreateUser
}

// NewMockCreateUser creates a new mock instance.
func NewMockCreateUser(ctrl *gomock.Controller) *MockCreateUser {
	mock := &MockCreateUser{ctrl: ctrl}
	mock.recorder = &MockCreateUserMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCreateUser) EXPECT() *MockCreateUserMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *MockCreateUser) CreateUser(ctx context.Context, username, email, password string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", ctx, username, email, password)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockCreateUserMockRecorder) CreateUser(ctx, username, email, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockCreateUser)(nil).CreateUser), ctx, username, email, password)
}
