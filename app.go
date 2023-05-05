package main

import (
	"todo_api/database"
	"todo_api/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	database.ConnectDB()

	app := fiber.New()
	app.Use(cors.New())
	api := app.Group("/api")
	todos := api.Group("/todos")
	routes.TodoRouter(todos)
	app.Listen(":3000")
}
