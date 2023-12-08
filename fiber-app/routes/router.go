package routes

import (
	"fiber-vue/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetUpRoutes(app *fiber.App) {
	app.Get("/allbooks", controllers.AllBooks)
}
