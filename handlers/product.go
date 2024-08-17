package handlers

import "github.com/gofiber/fiber/v2"

func GetProduct(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"hi": "hello"})
}

func GetReport(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"hi": "hello"})
}
