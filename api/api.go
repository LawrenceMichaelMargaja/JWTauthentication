package api

import "github.com/gofiber/fiber/v2"

func NewApi(s services.Services) *fiber.App {
	api := fiber.New()

	api.Get("/report", func(c *fiber.Ctx) error {
		report := s.GetReport()
		return c.JSON(report)
	})

	api.Get("/product", func(c *fiber.Ctx) error {
		product := s.GetProduct()
		return c.JSON(product)
	})

	return api
}
