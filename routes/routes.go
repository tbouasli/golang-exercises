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

func CreateUser(c *fiber.Ctx) error {
	var people models.People

	user := c.BodyParser(&people)
	if user != nil {
		return c.SendString("Erro ao processar os dados")
	} else if people.Name == "" || people.Age == 0 {
		return c.SendString("Nome e idade são obrigatórios")
	}

	result := database.DB.Create(&people)
	if result.Error != nil {
		return c.SendString("Erro ao criar registro")
	}

	return c.JSON(people)
}

func UpdateUser(c *fiber.Ctx) error {
	var people models.People
	// Recebendo o id pela URL
	id := c.Params("id")

	// Buscando no banco de dados por ID
	result := database.DB.First(&people, id)
	if result.Error != nil {
		return c.SendString("Registro não encontrado")
	}

	// Recebendo os dados do corpo da requisição
	user := c.BodyParser(&people)
	if user != nil {
		return c.SendString("Erro ao processar os dados")
	} else if people.Name == "" || people.Age == 0 {
		return c.SendString("Nome e idade são obrigatórios")
	}

	// Atualizando o registro
	result = database.DB.Model(&people).Updates(people)
	if result.Error != nil {
		return c.SendString("Erro ao atualizar registro")
	}

	// Retornando o registro atualizado
	return c.JSON(people)
}

func DeleteUser(c *fiber.Ctx) error {
	var people models.People
	id := c.Params("id")

	result := database.DB.Delete(&people, id)
	if result.Error != nil {
		return c.SendString("Erro ao deletar registro!")
	} else if result.RowsAffected == 0 {
		return c.SendString("Registro não encontrado!")
	}
	return c.SendString("Registro deletado com sucesso!")
}
