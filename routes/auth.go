package routes

import (
	"server/config"
	"server/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func Login(c *fiber.Ctx) error {

	jwt, err := utils.GenerateJWT(jwt.MapClaims{"name": "saurabh"}, config.Config.COOKIE_SECRET, config.Config.COOKIE_MAXAGE)

	if err == nil {
		c.Cookie(&fiber.Cookie{
			Name:     config.Config.COOKIE_NAME,
			Value:    jwt,
			Domain:   config.Config.COOKIE_DOMAIN,
			MaxAge:   config.Config.COOKIE_MAXAGE,
			SameSite: config.Config.COOKIE_SAMESITE,
		})
	}

	return nil
}

func Register(c *fiber.Ctx) error {
	return nil
}
