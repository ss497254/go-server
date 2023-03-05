package routes

import (
	"server/middleware"

	"github.com/gofiber/fiber/v2"
)

func Initalize(app *fiber.App) {
	app.Post("/login", Login)
	app.Post("/register", Register)
	app.Get("/me", middleware.Authenticated, GetUser)
}
