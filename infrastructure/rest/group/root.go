package group

import (
	"github.com/iwanjunaid/basesvc/infrastructure/rest/middleware"
	"github.com/iwanjunaid/basesvc/internal/interfaces"
)

func InitRoot(rest interfaces.Rest) {
	router := rest.GetRouter()

	router.Group("/", middleware.NewAuthentication(rest))
}
