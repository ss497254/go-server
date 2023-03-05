package middleware

import (
	"server/config"
	"server/utils"

	"github.com/gofiber/fiber/v2"
)

func Init(c *fiber.Ctx) error {
	c.Accepts("application/json")

	cookie := c.Cookies(config.Config.COOKIE_NAME)

	if cookie != "" {
		session, err := utils.ParseJWT(cookie, config.Config.COOKIE_SECRET)

		if err == nil {
			c.Locals("session", session)
		}
	}

	return c.Next()
}
