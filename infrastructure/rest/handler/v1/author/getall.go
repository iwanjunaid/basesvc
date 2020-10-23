package author

import (
	"github.com/gofiber/fiber/v2"
	"github.com/iwanjunaid/basesvc/internal/interfaces"
	"github.com/iwanjunaid/basesvc/shared/logger"
)

// GetAll handles HTTP GET request for retrieving multiple authors
func GetAll(rest interfaces.Rest) func(*fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		appController := rest.GetAppController()
		logger.WithFields(logger.Fields{"component": "get-all"}).Infof("get all authors")
		return appController.Author.GetAuthors(ctx)
	}
}
