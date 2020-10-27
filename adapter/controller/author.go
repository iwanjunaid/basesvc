package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/iwanjunaid/basesvc/domain/model"
	"github.com/iwanjunaid/basesvc/usecase/author/interactor"
)

type AuthorController interface {
	GetAuthors(c *fiber.Ctx) error
	InsertAuthorDocument(author *model.Author) error
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
func (a *AuthorControllerImpl) GetAuthors(c *fiber.Ctx) error {
	ctx := c.Context()
	authors, err := a.AuthorInteractor.GetAll(ctx)

	if err != nil {
		return err
	}

	return c.JSON(authors)
}

func (a *AuthorControllerImpl) InsertAuthorDocument(author *model.Author) error {
	return nil
}
