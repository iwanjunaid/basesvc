package author

import (
	"github.com/gofiber/fiber/v2"
	"github.com/iwanjunaid/basesvc/internal/interfaces"
	"github.com/iwanjunaid/basesvc/shared/logger"
)

// Update handles HTTP PATCH request for updating an author
func Update(rest interfaces.Rest) func(*fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		appController := rest.GetAppController()
		logger.WithFields(logger.Fields{"component": "update"}).Infof("update author")
		return appController.Author.GetAuthors(ctx)
	}
}
