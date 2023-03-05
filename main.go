package main

import (
	"server/config"
	"server/database"
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
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*", // comma string format e.g. "localhost, nikschaefer.tech"
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	routes.Initalize(app)
	// log.Fatal(app.Listen(":" + config.Config.PORT.GetValue()))
}
