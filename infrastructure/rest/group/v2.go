package group

import (
	"github.com/iwanjunaid/basesvc/internal/interfaces"
)

// @title BaseSVC API
// @version 1.0
// @description This is a sample basesvc server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /v2
func InitV2(rest interfaces.Rest) {
	router := rest.GetRouter()

	router.Group("/v2")
}
