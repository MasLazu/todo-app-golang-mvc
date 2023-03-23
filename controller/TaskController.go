package controller

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"todoApps/database"
	"todoApps/models"
)

type CreateTask struct {
	Data string `json:"data" xml:"data" query:"data" validate:"required"`
}

type MarkTask struct {
	ID uint `json:"id" query:"id"`
}

func IndexTaskController(c *fiber.Ctx) error {
	var allTask []models.Task

	if database.DB.Limit(10).Find(&allTask).Error != nil {
		return c.Render("index", fiber.Map{
			"baseUrl": c.BaseURL(),
			"error":   "error to get tasks data",
		})
	}

	return c.Render("index", fiber.Map{
		"baseUrl": c.BaseURL(),
		"tasks":   allTask,
	})
}

func CreateTaskController(c *fiber.Ctx) error {
	task := new(CreateTask)

	if err := c.BodyParser(task); err != nil {
		return c.Render("index", fiber.Map{
			"baseUrl": c.BaseURL(),
			"error":   "error internal server error",
		})
	}

	validate := validator.New()
	if err := validate.Struct(task); err != nil {
		return c.Render("index", fiber.Map{
			"baseUrl": c.BaseURL(),
			"error":   "error bad request",
		})
	}

	newTask := models.Task{
		Data: task.Data,
		Done: false,
	}

	result := database.DB.Create(&newTask)
	if result.Error != nil {
		return c.Render("index", fiber.Map{
			"baseUrl": c.BaseURL(),
			"error":   "error internal server error",
		})
	}

	return c.Redirect("/")
}

func MarkTaskController(c *fiber.Ctx) error {
	task := new(MarkTask)
	var markedTask models.Task

	if err := c.BodyParser(task); err != nil {
		return c.Redirect("/")
	}

	database.DB.Find(&markedTask, "id = ?", task.ID)
	markedTask.Done = !markedTask.Done
	database.DB.Save(&markedTask)
	return c.Redirect("/")
}

func DeleteTaskController(c *fiber.Ctx) error {
	task := new(MarkTask)
	c.BodyParser(task)
	database.DB.Delete(&models.Task{}, task.ID)
	return c.Redirect("/")
}
