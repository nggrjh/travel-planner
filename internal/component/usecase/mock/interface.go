// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockRegisterUser is a mock of RegisterUser interface.
type MockRegisterUser struct {
	ctrl     *gomock.Controller
	recorder *MockRegisterUserMockRecorder
}

// MockRegisterUserMockRecorder is the mock recorder for MockRegisterUser.
type MockRegisterUserMockRecorder struct {
	mock *MockRegisterUser
}

// NewMockRegisterUser creates a new mock instance.
func NewMockRegisterUser(ctrl *gomock.Controller) *MockRegisterUser {
	mock := &MockRegisterUser{ctrl: ctrl}
	mock.recorder = &MockRegisterUserMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRegisterUser) EXPECT() *MockRegisterUserMockRecorder {
	return m.recorder
}

// RegisterUser mocks base method.
func (m *MockRegisterUser) RegisterUser(ctx context.Context, email, password string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterUser", ctx, email, password)
	ret0, _ := ret[0].(error)
	return ret0
}

// RegisterUser indicates an expected call of RegisterUser.
func (mr *MockRegisterUserMockRecorder) RegisterUser(ctx, email, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterUser", reflect.TypeOf((*MockRegisterUser)(nil).RegisterUser), ctx, email, password)
}
