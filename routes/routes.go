package routes

import "github.com/gofiber/fiber/v2"

func ListAll(c *fiber.Ctx) error {
	return c.SendString("Teste")
}
