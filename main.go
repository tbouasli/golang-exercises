package main

import (
	"go-project/database"
	"go-project/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database.ConnectDatabase()
	app := fiber.New()
	app.Get("/", routes.ListAll)

	app.Listen(":3000")
}
