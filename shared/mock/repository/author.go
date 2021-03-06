// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/iwanjunaid/basesvc/usecase/author/repository (interfaces: AuthorSQLRepository,AuthorDocumentRepository,AuthorEventRepository)

// Package repository is a generated GoMock package.
package repository

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	model "github.com/iwanjunaid/basesvc/domain/model"
	reflect "reflect"
)

// MockAuthorSQLRepository is a mock of AuthorSQLRepository interface
type MockAuthorSQLRepository struct {
	ctrl     *gomock.Controller
	recorder *MockAuthorSQLRepositoryMockRecorder
}

// MockAuthorSQLRepositoryMockRecorder is the mock recorder for MockAuthorSQLRepository
type MockAuthorSQLRepositoryMockRecorder struct {
	mock *MockAuthorSQLRepository
}

// NewMockAuthorSQLRepository creates a new mock instance
func NewMockAuthorSQLRepository(ctrl *gomock.Controller) *MockAuthorSQLRepository {
	mock := &MockAuthorSQLRepository{ctrl: ctrl}
	mock.recorder = &MockAuthorSQLRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAuthorSQLRepository) EXPECT() *MockAuthorSQLRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockAuthorSQLRepository) Create(arg0 context.Context, arg1 *model.Author) (*model.Author, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(*model.Author)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockAuthorSQLRepositoryMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockAuthorSQLRepository)(nil).Create), arg0, arg1)
}

// Find mocks base method
func (m *MockAuthorSQLRepository) Find(arg0 context.Context, arg1 string) (*model.Author, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", arg0, arg1)
	ret0, _ := ret[0].(*model.Author)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find
func (mr *MockAuthorSQLRepositoryMockRecorder) Find(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockAuthorSQLRepository)(nil).Find), arg0, arg1)
}

// FindAll mocks base method
func (m *MockAuthorSQLRepository) FindAll(arg0 context.Context) ([]*model.Author, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAll", arg0)
	ret0, _ := ret[0].([]*model.Author)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAll indicates an expected call of FindAll
func (mr *MockAuthorSQLRepositoryMockRecorder) FindAll(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAll", reflect.TypeOf((*MockAuthorSQLRepository)(nil).FindAll), arg0)
}

// MockAuthorDocumentRepository is a mock of AuthorDocumentRepository interface
type MockAuthorDocumentRepository struct {
	ctrl     *gomock.Controller
	recorder *MockAuthorDocumentRepositoryMockRecorder
}

// MockAuthorDocumentRepositoryMockRecorder is the mock recorder for MockAuthorDocumentRepository
type MockAuthorDocumentRepositoryMockRecorder struct {
	mock *MockAuthorDocumentRepository
}

// NewMockAuthorDocumentRepository creates a new mock instance
func NewMockAuthorDocumentRepository(ctrl *gomock.Controller) *MockAuthorDocumentRepository {
	mock := &MockAuthorDocumentRepository{ctrl: ctrl}
	mock.recorder = &MockAuthorDocumentRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAuthorDocumentRepository) EXPECT() *MockAuthorDocumentRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockAuthorDocumentRepository) Create(arg0 context.Context, arg1 *model.Author) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create
func (mr *MockAuthorDocumentRepositoryMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockAuthorDocumentRepository)(nil).Create), arg0, arg1)
}

// FindAll mocks base method
func (m *MockAuthorDocumentRepository) FindAll(arg0 context.Context) ([]*model.Author, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAll", arg0)
	ret0, _ := ret[0].([]*model.Author)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAll indicates an expected call of FindAll
func (mr *MockAuthorDocumentRepositoryMockRecorder) FindAll(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAll", reflect.TypeOf((*MockAuthorDocumentRepository)(nil).FindAll), arg0)
}

// MockAuthorEventRepository is a mock of AuthorEventRepository interface
type MockAuthorEventRepository struct {
	ctrl     *gomock.Controller
	recorder *MockAuthorEventRepositoryMockRecorder
}

// MockAuthorEventRepositoryMockRecorder is the mock recorder for MockAuthorEventRepository
type MockAuthorEventRepositoryMockRecorder struct {
	mock *MockAuthorEventRepository
}

// NewMockAuthorEventRepository creates a new mock instance
func NewMockAuthorEventRepository(ctrl *gomock.Controller) *MockAuthorEventRepository {
	mock := &MockAuthorEventRepository{ctrl: ctrl}
	mock.recorder = &MockAuthorEventRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAuthorEventRepository) EXPECT() *MockAuthorEventRepositoryMockRecorder {
	return m.recorder
}

// Publish mocks base method
func (m *MockAuthorEventRepository) Publish(arg0 context.Context, arg1, arg2 []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Publish", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Publish indicates an expected call of Publish
func (mr *MockAuthorEventRepositoryMockRecorder) Publish(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Publish", reflect.TypeOf((*MockAuthorEventRepository)(nil).Publish), arg0, arg1, arg2)
}
