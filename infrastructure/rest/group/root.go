package group

import (
	"github.com/gofiber/fiber/v2"
	"github.com/iwanjunaid/basesvc/infrastructure/rest/middleware"
	"github.com/iwanjunaid/basesvc/internal/interfaces"
)

func InitRoot(rest interfaces.Rest) fiber.Router {
	router := rest.GetRouter()

	return router.Group("/", middleware.NewAuthentication(rest))
}
