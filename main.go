package main

import (
	"log"
	"server/config"
	"server/database"
	"server/middleware"
	"server/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()

	config.SetUpConfig()
	database.Connect()

	app.Use(logger.New())
	app.Use(middleware.Init)
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders:     "Accept,Authorization,Content-Type,X-CSRF-TOKEN",
		ExposeHeaders:    "Link",
		AllowCredentials: true,
	}))

	routes.Initalize(app)
	log.Fatal(app.Listen("localhost:" + config.Config.PORT))
}
