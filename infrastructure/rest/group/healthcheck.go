package group

import (
	"github.com/iwanjunaid/basesvc/infrastructure/rest/handler/healthcheck"
	"github.com/iwanjunaid/basesvc/internal/interfaces"
)

func InitHealthCheck(rest interfaces.Rest) {
	router := rest.GetRouter()

	healthCheckGroup := router.Group("/health")
	healthCheckGroup.Get("/", healthcheck.Check(rest))
}
