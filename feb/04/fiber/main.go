package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	app.Get("/status", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": true})
	})
	app.Listen(":8080")
}
