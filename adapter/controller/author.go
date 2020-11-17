package controller

import (
	"context"

	"github.com/iwanjunaid/basesvc/domain/model"
	"github.com/iwanjunaid/basesvc/usecase/author/interactor"
)

type AuthorController interface {
	GetAuthors(c context.Context) ([]*model.Author, error)
	InsertAuthor(c context.Context, author *model.Author) (*model.Author, error)
	InsertDocument(c context.Context, author *model.Author) error
}

type AuthorControllerImpl struct {
	AuthorInteractor interactor.AuthorInteractor
}

func NewAuthorController(interactor interactor.AuthorInteractor) AuthorController {
	return &AuthorControllerImpl{
		AuthorInteractor: interactor,
	}
}

// GetAuthors godoc
// @Summary Get all authors
// @Description list of authors
// @Tags authors
// @Produce json
// @Success 200 {array} model.Author
// @Router /authors [get]
func (a *AuthorControllerImpl) GetAuthors(ctx context.Context) ([]*model.Author, error) {
	authors, err := a.AuthorInteractor.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return authors, nil
}

func (a *AuthorControllerImpl) InsertAuthor(ctx context.Context, author *model.Author) (*model.Author, error) {
	author, err := a.AuthorInteractor.Create(ctx, author)
	if err != nil {
		return nil, err
	}
	return author, nil
}

func (a *AuthorControllerImpl) InsertDocument(ctx context.Context, author *model.Author) error {
	err := a.AuthorInteractor.CreateDocument(ctx, author)
	if err != nil {
		return err
	}
	return nil
}
