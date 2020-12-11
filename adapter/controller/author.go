package controller

import (
	"context"
	"crypto/sha256"
	"fmt"

	"github.com/iwanjunaid/basesvc/domain/model"
	athri "github.com/iwanjunaid/basesvc/usecase/author/interactor"
	gvtri "github.com/iwanjunaid/basesvc/usecase/gravatar/interactor"
)

type AuthorController interface {
	GetAuthor(c context.Context, id string) (*model.Author, error)
	GetAuthors(c context.Context) ([]*model.Author, error)
	InsertAuthor(c context.Context, author *model.Author) (*model.Author, error)
	InsertDocument(c context.Context, author *model.Author) error
}

type AuthorControllerImpl struct {
	AuthorInteractor   athri.AuthorInteractor
	GravatarInteractor gvtri.GravatarInteractor
}

func NewAuthorController(ai athri.AuthorInteractor, gi gvtri.GravatarInteractor) AuthorController {
	return &AuthorControllerImpl{
		AuthorInteractor:   ai,
		GravatarInteractor: gi,
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
	authors, err := a.AuthorInteractor.GetAll(ctx, "all_authors")
	if err != nil {
		return nil, err
	}

	for _, author := range authors {
		author, err = a.setAvatar(ctx, author)
	}
	if err != nil {
		return authors, err
	}

	return authors, nil
}

// GetAuthor godoc
// @Summary Get author by id
// @Description get author by id
// @Tags authors
// @Produce json
// @Param id path string true "Author ID"
// @Success 200 {array} model.Author
// @Router /authors/{id} [get]
func (a *AuthorControllerImpl) GetAuthor(ctx context.Context, id string) (*model.Author, error) {
	h := sha256.New()
	h.Write([]byte(fmt.Sprintf("author_id=%s", id)))
	key := fmt.Sprintf("%x", h.Sum(nil))

	author, err := a.AuthorInteractor.Get(ctx, key, id)
	if err != nil {
		return nil, err
	}

	author, err = a.setAvatar(ctx, author)
	if err != nil {
		return author, err
	}
	return author, nil
}

// InsertAuthor godoc
// @Summary Create a new author
// @Description create a new author
// @Tags authors
// @Accept json
// @Produce json
// @Param author body model.Author true "author"
// @Success 200
// @Router /authors [post]
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

func (a *AuthorControllerImpl) setAvatar(ctx context.Context, author *model.Author) (*model.Author, error) {
	// Set Gravatar Profile
	var avatar string

	profile, err := a.GravatarInteractor.Get(ctx, author.Email)
	if err != nil {
		return author, err
	}

	if profile != nil && len(profile.Entry) > 0 {
		avatar = profile.Entry[0].ThumbnailUrl
	}

	author.Avatar = avatar

	return author, nil
}
