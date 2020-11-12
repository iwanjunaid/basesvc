package group

import (
	"github.com/gofiber/fiber/v2"
	"github.com/iwanjunaid/basesvc/infrastructure/rest/handler/v1/author"
	"github.com/iwanjunaid/basesvc/internal/interfaces"
)

func InitAuthorV1(rest interfaces.Rest, v1 fiber.Router) {
	authorGroup := v1.Group("/authors")
	authorGroup.Get("/", author.GetAll(rest))
	authorGroup.Get("/:id", author.Get(rest))
	authorGroup.Post("/", author.Create(rest))
	authorGroup.Patch("/:id", author.Update(rest))
	authorGroup.Delete("/:id", author.Delete(rest))
}
