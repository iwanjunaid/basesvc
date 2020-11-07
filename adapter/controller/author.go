package controller

import (
	"net/http"

	"github.com/RoseRocket/xerrs"
	"github.com/gofiber/fiber/v2"
	"github.com/iwanjunaid/basesvc/domain/model"
	"github.com/iwanjunaid/basesvc/internal/logger"
	"github.com/iwanjunaid/basesvc/usecase/author/interactor"

	log "github.com/sirupsen/logrus"
)

type AuthorController interface {
	GetAuthors(c *fiber.Ctx) error
	CreateAuthor(c *fiber.Ctx, a *model.Author) error
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

	c.JSON(authors)
	return nil

}

// CreateAuthor godoc
// @Summary Create a new author
// @Description create a new author
// @Tags authors
// @Accept json
// @Produce json
// @Param author body model.Author true "author"
// @Success 200
// @Router /authors [post]
func (a *AuthorControllerImpl) CreateAuthor(c *fiber.Ctx, author *model.Author) error {
	ctx := c.Context()

	err := a.AuthorInteractor.Create(ctx, author)

	if err != nil {
		logger.LogEntrySetFields(c, log.Fields{
			"stack_trace": xerrs.Details(err, logger.ErrMaxStack),
			"context":     "CreateAuthor",
			"resp_status": http.StatusInternalServerError,
		})
		c.Status(500)
		return err

	}

	c.JSON(author)

	return nil

}
