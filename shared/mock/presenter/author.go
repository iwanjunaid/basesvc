// Code generated by MockGen. DO NOT EDIT.
// Source: ./usecase/author/presenter/author.go

// Package presenter is a generated GoMock package.
package presenter

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	model "github.com/iwanjunaid/basesvc/domain/model"
	reflect "reflect"
)

// MockAuthorPresenter is a mock of AuthorPresenter interface
type MockAuthorPresenter struct {
	ctrl     *gomock.Controller
	recorder *MockAuthorPresenterMockRecorder
}

// MockAuthorPresenterMockRecorder is the mock recorder for MockAuthorPresenter
type MockAuthorPresenterMockRecorder struct {
	mock *MockAuthorPresenter
}

// NewMockAuthorPresenter creates a new mock instance
func NewMockAuthorPresenter(ctrl *gomock.Controller) *MockAuthorPresenter {
	mock := &MockAuthorPresenter{ctrl: ctrl}
	mock.recorder = &MockAuthorPresenterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAuthorPresenter) EXPECT() *MockAuthorPresenterMockRecorder {
	return m.recorder
}

// ResponseUsers mocks base method
func (m *MockAuthorPresenter) ResponseUsers(arg0 context.Context, arg1 []*model.Author) ([]*model.Author, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResponseUsers", arg0, arg1)
	ret0, _ := ret[0].([]*model.Author)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ResponseUsers indicates an expected call of ResponseUsers
func (mr *MockAuthorPresenterMockRecorder) ResponseUsers(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResponseUsers", reflect.TypeOf((*MockAuthorPresenter)(nil).ResponseUsers), arg0, arg1)
}
