package healthcheck

import (
	"github.com/gofiber/fiber/v2"
	"github.com/iwanjunaid/basesvc/internal/interfaces"
)

func Check(rest interfaces.Rest) func(*fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		return ctx.SendStatus(200)
	}
}
