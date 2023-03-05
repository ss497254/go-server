package utils

import (
	"server/config"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func SendAccessToken(payload jwt.MapClaims, c *fiber.Ctx) {
	jwt, err := GenerateJWT(payload, config.Config.COOKIE_SECRET, config.Config.COOKIE_MAXAGE)

	if err == nil {
		c.Cookie(&fiber.Cookie{
			Name:     config.Config.COOKIE_NAME,
			Value:    jwt,
			Domain:   config.Config.COOKIE_DOMAIN,
			MaxAge:   config.Config.COOKIE_MAXAGE,
			SameSite: config.Config.COOKIE_SAMESITE,
		})
	}
}

func ClearAccessToken(c *fiber.Ctx) {
	c.ClearCookie(config.Config.COOKIE_NAME)
}
