package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/iwanjunaid/basesvc/internal/interfaces"
	"github.com/iwanjunaid/basesvc/shared/logger"
)

// NewAuthentication creates middleware for handling api authentication
func NewAuthentication(rest interfaces.Rest) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		// Implement your authentication here
		logger.WithFields(logger.Fields{"component": "middleware"}).Infof("implementation of middleware")
		return c.Next()
	}
}
