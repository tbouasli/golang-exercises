package main

import (
	"go-project/database"
	"go-project/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database.ConnectDatabase()

	app := fiber.New()

	app.Get("/people", routes.ListByIdOrAll)
	app.Post("/people", routes.CreateUser)
	app.Put("/people/:id", routes.UpdateUser)
	app.Delete("/people/:id", routes.DeleteUser)

	app.Listen(":3000")
}
