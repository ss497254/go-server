package middleware

import "github.com/gofiber/fiber/v2"

func Authenticated(c *fiber.Ctx) error {
	session := c.Locals("session")

	if session == nil {
		return c.JSON(&fiber.Map{
			"message": "Please Login to continue",
		})
	}

	return c.Next()
}
