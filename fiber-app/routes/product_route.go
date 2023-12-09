package routes

import (
	"fiber-vue/handlers"
	"fiber-vue/middleware"

	"github.com/gofiber/fiber/v2"
)

func ProductRoute(route fiber.Router) {
	// product
	product := route.Group("product")
	product.Get("/", middleware.Protected(), handlers.GetAllProduct)
	product.Post("/", middleware.Protected(), handlers.CreateProduct)
	product.Patch("/", middleware.Protected(), handlers.UpdateProductQuantity)
}
