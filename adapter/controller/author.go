package controller

import (
	"context"
	"net/http"

	"github.com/iwanjunaid/basesvc/internal/respond"

	"github.com/RoseRocket/xerrs"
	"github.com/gofiber/fiber/v2"
	"github.com/iwanjunaid/basesvc/domain/model"
	"github.com/iwanjunaid/basesvc/internal/logger"
	"github.com/iwanjunaid/basesvc/usecase/author/interactor"

	log "github.com/sirupsen/logrus"
)

type AuthorController interface {
	GetAuthors(c *fiber.Ctx) error
	InsertAuthor(c context.Context, author *model.Author) (*model.Author, error)
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

func (a *AuthorControllerImpl) InsertAuthor(ctx context.Context, author *model.Author) (*model.Author, error) {
	author, err := a.AuthorInteractor.Create(ctx, author)
	if err != nil {
		return nil, err
	}
	return author, nil
}

func (a *AuthorControllerImpl) InsertDocument(author *model.Author) error {
	err := a.AuthorInteractor.CreateDocument(context.Background(), author)
	if err != nil {
		return err
	}
	return nil
}
