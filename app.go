package main

import (
	"os"
	"todo_api/database"
	"todo_api/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	database.ConnectDB()

	app := fiber.New()
	app.Use(cors.New())
	api := app.Group("/api")
	todos := api.Group("/todos")

	port, exists := os.LookupEnv("PORT")
	if exists == false {
		port = ":3000"
	}

	routes.TodoRouter(todos)
	app.Listen(port)
}
