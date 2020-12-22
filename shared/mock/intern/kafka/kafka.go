// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/iwanjunaid/basesvc/internal/kafka (interfaces: InternalKafka)

// Package kafka is a generated GoMock package.
package kafka

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockInternalKafka is a mock of InternalKafka interface
type MockInternalKafka struct {
	ctrl     *gomock.Controller
	recorder *MockInternalKafkaMockRecorder
}

// MockInternalKafkaMockRecorder is the mock recorder for MockInternalKafka
type MockInternalKafkaMockRecorder struct {
	mock *MockInternalKafka
}

// NewMockInternalKafka creates a new mock instance
func NewMockInternalKafka(ctrl *gomock.Controller) *MockInternalKafka {
	mock := &MockInternalKafka{ctrl: ctrl}
	mock.recorder = &MockInternalKafkaMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockInternalKafka) EXPECT() *MockInternalKafkaMockRecorder {
	return m.recorder
}

// Publish mocks base method
func (m *MockInternalKafka) Publish(arg0 context.Context, arg1, arg2 []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Publish", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Publish indicates an expected call of Publish
func (mr *MockInternalKafkaMockRecorder) Publish(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Publish", reflect.TypeOf((*MockInternalKafka)(nil).Publish), arg0, arg1, arg2)
}