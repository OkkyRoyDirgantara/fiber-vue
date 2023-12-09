package routes

import (
	"fiber-vue/handlers"

	"github.com/gofiber/fiber/v2"
)

func AuthRoute(route fiber.Router) {
	// auth
	auth := route.Group("auth")
	auth.Post("/login", handlers.AuthLogin)

	// user
	user := route.Group("user")
	user.Post("/register", handlers.UserRegister)
}
