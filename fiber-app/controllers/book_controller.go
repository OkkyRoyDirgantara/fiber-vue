package controllers

import "github.com/gofiber/fiber/v2"

func AllBooks(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON("OK")
}
