package main

import (
	"crud-echo/config"
	"crud-echo/handlers"
	"crud-echo/models"

	"github.com/gofiber/fiber/v2"
)

func main() {
	config.ConnectDatabase()
	models.Migrate(config.DB)
	app := fiber.New()
	app.Post("/users", handlers.CreateUser)
	app.Get("/users", handlers.GetUsers)
	app.Get("/users/:id", handlers.GetUser)
	app.Put("/users/:id", handlers.UpdateUser)
	app.Delete("/users/:id", handlers.DeleteUser)

	app.Listen(":8080")
}
