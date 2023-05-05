package handlers

import (
	"todo_api/database"
	"todo_api/dto"
	"todo_api/models"

	"github.com/gofiber/fiber/v2"
)

func CreateTodo(c *fiber.Ctx) error {
	todoDTO := new(dto.CreateTodoDTO)

	if err := c.BodyParser(todoDTO); err != nil {
		return c.Status(400).JSON(fiber.Map{"message": err.Error()})
	}

	todo := models.Todo{
		Title:       todoDTO.Title,
		Description: todoDTO.Description,
	}

	database.DBConn.Create(&todo)
	return c.JSON(fiber.Map{"todo": todo, "status": "success"})
}

func GetTodo(c *fiber.Ctx) error {
	todo := new(models.Todo)

	result := database.DBConn.First(&todo, c.Params("id"))

	if result.Error != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "Todo not found",
			"status":  "error"})
	}

	return c.JSON(fiber.Map{"todo": todo, "status": "success"})
}

func GetTodos(c *fiber.Ctx) error {
	todos := []models.Todo{}
	database.DBConn.Find(&todos)
	return c.JSON(fiber.Map{"todos": todos, "status": "success"})
}

func DeleteTodo(c *fiber.Ctx) error {
	todo := new(models.Todo)

	result := database.DBConn.First(&todo, c.Params("id"))

	if result.Error != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "Todo not found",
			"status":  "error"})
	}

	database.DBConn.Delete(&todo)

	return c.JSON(fiber.Map{"deleted": todo.ID, "status": "success"})
}

func UpdateTodo(c *fiber.Ctx) error {
	todoDTO := new(dto.UpdateTodoDTO)
	todo := new(models.Todo)

	if database.DBConn.First(&todo, c.Params("id")).Error != nil {
		return c.Status(404).JSON(fiber.Map{"message": "Todo not found"})
	}

	if err := c.BodyParser(todoDTO); err != nil {
		return c.Status(400).JSON(fiber.Map{"message": err.Error()})
	}

	database.DBConn.Model(&todo).Updates(todoDTO)
	return c.JSON(fiber.Map{"todo": todo, "status": "success"})
}

func UpdateTodoStatus(c *fiber.Ctx) error {
	todoDTO := new(dto.UpdateTodoStatusDTO)
	todo := new(models.Todo)

	if database.DBConn.First(&todo, c.Params("id")).Error != nil {
		return c.Status(404).JSON(fiber.Map{"message": "Todo not found"})
	}

	if err := c.BodyParser(todoDTO); err != nil {
		return c.Status(400).JSON(fiber.Map{"message": err.Error()})
	}

	database.DBConn.Model(&todo).Update("completed", todoDTO.Completed)
	return c.JSON(fiber.Map{"todo": todo, "status": "success"})
}
