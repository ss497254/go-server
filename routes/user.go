package routes

import (
	"github.com/gofiber/fiber/v2"
)

func GetUser(c *fiber.Ctx) error {
	session := c.Locals("session")

	c.JSON(fiber.Map{
		"name":    "Saurabh Singh",
		"session": session,
	})

	return nil
}
