package controller

import (
	"context"
	"net/http"

	"github.com/iwanjunaid/basesvc/internal/respond"

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
	InsertAuthor(c *fiber.Ctx) error
	InsertDocument(author *model.Author) error
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
		return respond.Fail(c, http.StatusInternalServerError, http.StatusInternalServerError, err)

	}
	return respond.Success(c, http.StatusOK, authors)
}

func (a *AuthorControllerImpl) InsertAuthor(c *fiber.Ctx) error {
	var author *model.Author
	if err := c.BodyParser(&author); err != nil {
		return err
	}
	// Validate Author
	err := ValidateAuthors(author)
	if err != nil {
		logger.LogEntrySetFields(c, log.Fields{
			"stack_trace": xerrs.Details(err, logger.ErrMaxStack),
			"context":     "InsertAuthor",
			"resp_status": http.StatusBadRequest,
		})
		return respond.Fail(c, http.StatusBadRequest, http.StatusInternalServerError, err)
	}
	authorResult, err := a.AuthorInteractor.Create(c.Context(), author)
	if err != nil {
		logger.LogEntrySetFields(c, log.Fields{
			"stack_trace": xerrs.Details(err, logger.ErrMaxStack),
			"context":     "InsertAuthor",
			"resp_status": http.StatusInternalServerError,
		})
		return respond.Fail(c, http.StatusInternalServerError, http.StatusInternalServerError, err)

	}
	return respond.Success(c, http.StatusOK, authorResult)
}

func (a *AuthorControllerImpl) InsertDocument(author *model.Author) error {
	err := a.AuthorInteractor.CreateDocument(context.Background(), author)
	if err != nil {
		return err
	}
	return nil
}

// ValidateAuthors validates Author struct
func ValidateAuthors(author *model.Author) error {
	var err error

	err = validation.ValidateStruct(&author,
		validation.Field(&author.Name, validation.Required),
		validation.Field(&author.Email, validation.Required, is.Email),
	)
	return err
}
