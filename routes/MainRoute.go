package routes

import (
	"github.com/gofiber/fiber/v2"
	"todoApps/controller"
)

func MainRoute(app *fiber.App) {
	app.Static("/public", "./public")

	app.Get("/", controller.IndexTaskController).Name("index")
	app.Post("/", controller.CreateTaskController).Name("create")
	app.Post("/mark", controller.MarkTaskController).Name("mark")
	app.Post("/delete", controller.DeleteTaskController).Name("delete")
}
