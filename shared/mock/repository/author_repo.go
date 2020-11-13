package repository

import (
	"context"

	"github.com/iwanjunaid/basesvc/domain/model"
	"github.com/stretchr/testify/mock"
)

type MockAuthorRepository struct {
	mock.Mock
}

func (m *MockAuthorRepository) Create(ctx context.Context, author *model.Author) (*model.Author, error) {
	call := m.Called(ctx, author)
	res := call.Get(0)
	if res == nil {
		return nil, call.Error(1)
	}
	return nil, nil
}

func (m *MockAuthorRepository) FindAll(ctx context.Context) ([]*model.Author, error) {
	call := m.Called(ctx)
	res := call.Get(0)
	if res == nil {
		return nil, call.Error(1)
	}
	return nil, nil
}
