package author

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/iwanjunaid/basesvc/internal/interfaces"
	"github.com/iwanjunaid/basesvc/internal/respond"
	"github.com/iwanjunaid/basesvc/internal/telemetry"
)

// GetAll handles HTTP GET request for retrieving multiple authors
func GetAll(rest interfaces.Rest) func(c *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		txn := telemetry.GetTelemetry(ctx.Context())
		defer txn.End()
		appController := rest.GetAppController()
		authorRes, err := appController.Author.GetAuthors(ctx.Context())
		if err != nil {
			txn.NoticeError(err)
			return respond.Fail(ctx, http.StatusInternalServerError, http.StatusInternalServerError, err)
		}
		return respond.Success(ctx, http.StatusOK, authorRes)
	}
}
