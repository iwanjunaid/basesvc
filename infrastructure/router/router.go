package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/iwanjunaid/basesvc/adapter/controller"
)

// NewRouter create new router
func NewRouter(app *fiber.App, c controller.AppController) {
	app.Use(cors.New())
	app.Use(logger.New())
	// app.Use(middleware.Recover()) // ?

	app.Get("/authors", func(ctx *fiber.Ctx) error {
		return c.Author.GetAuthors(ctx)
	})
}
