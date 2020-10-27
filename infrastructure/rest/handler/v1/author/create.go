package author

import (
	"github.com/gofiber/fiber/v2"
	"github.com/iwanjunaid/basesvc/internal/interfaces"
	"github.com/iwanjunaid/basesvc/shared/logger"
)

// Create handles HTTP POST requst for creating a new author
func Create(rest interfaces.Rest) func(*fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		appController := rest.GetAppController()
		logger.WithFields(logger.Fields{"component": "create"}).Infof("create author")
		return appController.Author.GetAuthors(ctx)
	}
}
