package routes

import (
	"go-project/database"
	"go-project/models"

	"github.com/gofiber/fiber/v2"
)

func ListByIdOrAll(c *fiber.Ctx) error {
	id := c.Query("id")
	if id != "" {
		var people models.People
		result := database.DB.First(&people, id)
		if result.Error != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Registro não encontrado",
			})
		}

		return c.JSON(people)
	}

	var people []models.People
	result := database.DB.Find(&people)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Erro ao buscar dados",
		})
	}

	return c.JSON(people)
}
