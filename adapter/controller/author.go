package controller

import (
	"net/http"

	"github.com/RoseRocket/xerrs"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/gofiber/fiber/v2"
	"github.com/iwanjunaid/basesvc/domain/model"
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
		c.Status(500)
		return err

	}

	// Validate Authors
	err = ValidateAuthors(authors)
	if err != nil {
		logger.LogEntrySetFields(c, log.Fields{
			"stack_trace": xerrs.Details(err, logger.ErrMaxStack),
			"context":     "GetAuthors",
			"resp_status": http.StatusInternalServerError,
		})
		c.Status(500)
		return err

	}

	c.JSON(authors)
	return nil

}

// ValidateAuthors validates Author struct
func ValidateAuthors(authors []*model.Author) error {
	var err error
	for _, a := range authors {
		if err = validation.ValidateStruct(&a,
			validation.Field(&a.ID, validation.Required),
			validation.Field(&a.Name, validation.Required),
			validation.Field(&a.Email, validation.Required, is.Email),
		); err != nil {
		}
	}
	return err
}
