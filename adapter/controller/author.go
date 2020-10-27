package controller

import (
	"net/http"

	"github.com/RoseRocket/xerrs"
	"github.com/gofiber/fiber/v2"
	"github.com/iwanjunaid/basesvc/internal/logger"
	"github.com/iwanjunaid/basesvc/usecase/author/interactor"

	log "github.com/sirupsen/logrus"
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
		logger.LogEntrySetFields(c, log.Fields{
			"stack_trace": xerrs.Details(err, logger.ErrMaxStack),
			"context":     "GetAuthors",
			"resp_status": http.StatusInternalServerError,
		})
		return err
	}

	return c.JSON(authors)
}
