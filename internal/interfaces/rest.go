package interfaces

import (
	"github.com/gofiber/fiber/v2"
	"github.com/iwanjunaid/basesvc/adapter/controller"
)

type Rest interface {
	GetRouter() *fiber.App
	GetAppController() *controller.AppController
}
