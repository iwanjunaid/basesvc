package author

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/iwanjunaid/basesvc/internal/interfaces"
	"github.com/iwanjunaid/basesvc/internal/respond"
)

// Get handles HTTP GET request for retrieving an author
func Get(rest interfaces.Rest) func(*fiber.Ctx) error {
	return func(ctx *fiber.Ctx) (err error) {
		appController := rest.GetAppController()
		authorRes, err := appController.Author.GetAuthor(ctx.Context(), ctx.Params("id"))

		if err != nil {
			return respond.Fail(ctx, http.StatusInternalServerError, http.StatusInternalServerError, err)
		}
		return respond.Success(ctx, http.StatusOK, authorRes)
	}
}
