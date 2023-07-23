package configuration

import (
	"github.com/gofiber/fiber/v2"
)

func NewFiberConfiguration() fiber.Config {
	return fiber.Config{
		ErrorHandler: ErrorHandler,
	}
}
