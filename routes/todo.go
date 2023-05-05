package routes

import (
	"todo_api/handlers"

	"github.com/gofiber/fiber/v2"
)

func TodoRouter(app fiber.Router) {
	app.Get("/", handlers.GetTodos)
	app.Post("/", handlers.CreateTodo)
	app.Get("/:id<int>", handlers.GetTodo)
	app.Patch("/:id<int>", handlers.UpdateTodo)
	app.Delete("/:id<int>", handlers.DeleteTodo)
	app.Put("/:id<int>/status", handlers.UpdateTodoStatus)
}
