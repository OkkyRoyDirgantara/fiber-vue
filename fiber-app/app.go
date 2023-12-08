package main

import (
	"fiber-vue/database"
	"fiber-vue/routes"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	database.ConnectDb()
	app := fiber.New()

	app.Use(cors.New())
	app.Use(logger.New())

	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404) // => 404 "Not Found"
	})

	routes.SetUpRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
