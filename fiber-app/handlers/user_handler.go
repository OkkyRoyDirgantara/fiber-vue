package handlers

import (
	"fiber-vue/database"
	"fiber-vue/models"
	"fiber-vue/utils"

	"github.com/gofiber/fiber/v2"
)

func UserRegister(ctx *fiber.Ctx) error {
	// Parse request body into NewUser struct
	User := new(models.User)
	if err := ctx.BodyParser(User); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Error on register request", "data": err})
	}

	// Validate the NewUser struct
	if err := utils.ValidationInput(User); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Validation error", "data": err.Error()})
	}

	// Hash the password
	hash, err := utils.HashPassword(User.Password)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Error on register request", "data": err})
	}

	// Update the password with the hashed one
	User.Password = hash

	// Save the user to the database
	db := database.DBConn
	if err := db.Create(&User).Error; err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Error on register request", "data": err})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "User successfully registered", "data": User})
}
