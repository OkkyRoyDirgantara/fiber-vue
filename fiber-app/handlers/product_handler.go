package handlers

import (
	"fiber-vue/database"
	"fiber-vue/models"
	"fiber-vue/utils"

	"github.com/gofiber/fiber/v2"
)

type Product struct {
	Code        string `validate:"required,min=3,max=10"`
	Name        string `validate:"required,min=3,max=20"`
	Stock       int    `validate:"required,min=1"`
	Description string `validate:"required,min=3,max=255"`
	Status      bool
}

func GetAllProduct(ctx *fiber.Ctx) error {
	db := database.DBConn
	var Products []Product
	db.Find(&Products)
	return ctx.JSON(fiber.Map{
		"status":  "success",
		"message": "All Products",
		"data":    Products,
	})
}

func CreateProduct(ctx *fiber.Ctx) error {
	db := database.DBConn

	Product := new(Product)
	if err := ctx.BodyParser(Product); err != nil {
		return ctx.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	// Validate the  Product struct
	if err := utils.ValidationInput(Product); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Validation error", "data": err.Error()})
	}

	if err := db.Create(&Product).Error; err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "Could not create product",
			"data":    err,
		})
	}

	return ctx.JSON(fiber.Map{
		"status":  "success",
		"message": "Product created",
		"data":    Product,
	})
}

func UpdateProductQuantity(ctx *fiber.Ctx) error {
	db := database.DBConn

	// Parse request body
	var updateData struct {
		Code   string `json:"code" validate:"required,min=3,max=10"`
		Action string `json:"action" validate:"required,oneof=add subtract"`
		Amount int    `json:"amount" validate:"required,min=1"`
	}

	if err := ctx.BodyParser(&updateData); err != nil {
		return ctx.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	// Validate request body
	if err := utils.ValidationInput(&updateData); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Validation error", "data": err.Error()})
	}

	var product Product
	if db.Where("code = ?", updateData.Code).First(&product).RowsAffected == 0 {
		return ctx.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "Product not found",
			"data":    nil,
		})
	}

	// Update product quantity based on the action
	switch updateData.Action {
	case "add":
		product.Stock += updateData.Amount
	case "subtract":
		if product.Stock < updateData.Amount {
			return ctx.Status(400).JSON(fiber.Map{
				"status":  "error",
				"message": "Insufficient stock",
				"data":    nil,
			})
		}
		product.Stock -= updateData.Amount
	}

	modelProduct := new(models.Product)
	modelProduct.Code = product.Code
	modelProduct.Name = product.Name
	modelProduct.Stock = product.Stock

	db.Model(&models.Product{}).Where("code = ?", product.Code).Update("stock", product.Stock)

	return ctx.JSON(fiber.Map{
		"status":  "success",
		"message": "Product quantity updated",
		"data":    product,
	})
}
