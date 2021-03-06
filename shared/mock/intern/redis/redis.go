// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/iwanjunaid/basesvc/intern/redis (interfaces: InternalRedis)

// Package redis is a generated GoMock package.
package redis

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockInternalRedis is a mock of InternalRedis interface
type MockInternalRedis struct {
	ctrl     *gomock.Controller
	recorder *MockInternalRedisMockRecorder
}

// MockInternalRedisMockRecorder is the mock recorder for MockInternalRedis
type MockInternalRedisMockRecorder struct {
	mock *MockInternalRedis
}

// NewMockInternalRedis creates a new mock instance
func NewMockInternalRedis(ctrl *gomock.Controller) *MockInternalRedis {
	mock := &MockInternalRedis{ctrl: ctrl}
	mock.recorder = &MockInternalRedisMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockInternalRedis) EXPECT() *MockInternalRedisMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockInternalRedis) Create(arg0 context.Context, arg1 string, arg2 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create
func (mr *MockInternalRedisMockRecorder) Create(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockInternalRedis)(nil).Create), arg0, arg1, arg2)
}

// Get mocks base method
func (m *MockInternalRedis) Get(arg0 context.Context, arg1 string, arg2 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Get indicates an expected call of Get
func (mr *MockInternalRedisMockRecorder) Get(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockInternalRedis)(nil).Get), arg0, arg1, arg2)
}
