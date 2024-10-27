package root

import (
	"github.com/gofiber/fiber/v2"
)

func GetRoot(c *fiber.Ctx) error {
	return c.Status(200).JSON(fiber.Map{
		"message": "Hello, world!",
	})
}
