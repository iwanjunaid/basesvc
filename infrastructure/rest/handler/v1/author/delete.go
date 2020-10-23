package author

import (
	"github.com/gofiber/fiber/v2"
	"github.com/iwanjunaid/basesvc/internal/interfaces"
	"github.com/iwanjunaid/basesvc/shared/logger"
)

// Delete handles HTTP DELETE request for deleting an author
func Delete(rest interfaces.Rest) func(*fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		appController := rest.GetAppController()
		logger.WithFields(logger.Fields{"component": "delete"}).Infof("delete author")
		return appController.Author.GetAuthors(ctx)
	}
}
