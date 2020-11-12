package group

import (
	"github.com/gofiber/fiber/v2"
	"github.com/iwanjunaid/basesvc/infrastructure/rest/handler/healthcheck"
	"github.com/iwanjunaid/basesvc/internal/interfaces"
)

func InitHealthCheck(rest interfaces.Rest, root fiber.Router) {

	healthCheckGroup := root.Group("/health")
	healthCheckGroup.Get("/", healthcheck.Check(rest))
}
