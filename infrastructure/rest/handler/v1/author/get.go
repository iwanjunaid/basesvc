package author

import (
	"github.com/gofiber/fiber/v2"
	"github.com/iwanjunaid/basesvc/internal/interfaces"
	"github.com/iwanjunaid/basesvc/shared/logger"
)

// Get handles HTTP GET request for retrieving an author
func Get(rest interfaces.Rest) func(*fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		appController := rest.GetAppController()
		logger.WithFields(logger.Fields{"component": "get"}).Infof("get author")
		return appController.Author.GetAuthors(ctx)
	}
}
