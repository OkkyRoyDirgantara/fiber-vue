package routes

import (
	"github.com/gofiber/fiber/v2"
)

func SetUpRoutes(app *fiber.App) {
	api := app.Group("api/v1")
	AuthRoute(api)
	ProductRoute(api)
}
