package group

import (
	"github.com/gofiber/fiber/v2"
	"github.com/iwanjunaid/basesvc/infrastructure/rest/handler/v2/author"
	"github.com/iwanjunaid/basesvc/internal/interfaces"
)

func InitAuthorV2(rest interfaces.Rest, v2 fiber.Router) {
	authorGroup := v2.Group("/v2/authors")
	authorGroup.Get("/", author.GetAll(rest))
	authorGroup.Get("/:id", author.Get(rest))
	authorGroup.Post("/", author.Create(rest))
	authorGroup.Patch("/:id", author.Update(rest))
	authorGroup.Delete("/:id", author.Delete(rest))
}
