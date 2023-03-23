package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/django"
	"todoApps/database"
	"todoApps/database/migration"
	"todoApps/routes"
)

func main() {
	engine := django.New("./views", ".django")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	database.DatabaseInit()
	migration.RunMigrations()

	routes.MainRoute(app)

	app.Listen(":3000")
}
