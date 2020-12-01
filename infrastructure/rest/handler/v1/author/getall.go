package author

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/iwanjunaid/basesvc/internal/interfaces"
	"github.com/iwanjunaid/basesvc/internal/respond"
)

// GetAll handles HTTP GET request for retrieving multiple authors
func GetAll(rest interfaces.Rest) func(c *fiber.Ctx) (err error) {
	return func(ctx *fiber.Ctx) error {
		appController := rest.GetAppController()
		authorRes, err := appController.Author.GetAuthors(ctx.Context())

		if err != nil {
			return respond.Fail(ctx, http.StatusInternalServerError, http.StatusInternalServerError, err)
		}
		return respond.Success(ctx, http.StatusOK, authorRes)
	}
}
