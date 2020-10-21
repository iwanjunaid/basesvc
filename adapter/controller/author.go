package controller

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/iwanjunaid/basesvc/usecase/author/interactor"
)

type AuthorController interface {
	GetAuthors(c *fiber.Ctx) error
}

type AuthorControllerImpl struct {
	AuthorInteractor interactor.AuthorInteractor
}

func NewAuthorController(interactor interactor.AuthorInteractor) AuthorController {
	return &AuthorControllerImpl{
		AuthorInteractor: interactor,
	}
}

func (a *AuthorControllerImpl) GetAuthors(c *fiber.Ctx) error {
	ctx := c.Context()
	authors, err := a.AuthorInteractor.GetAll(ctx)

	if err != nil {
		return err
	}

	if authors != nil {
		c.Response().SetStatusCode(http.StatusNotFound)
	}

	return c.JSON(authors)
}
