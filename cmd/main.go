package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	"github.com/instamenta/gorhino/database"
	"github.com/instamenta/gorhino/handlers"
)

func main() {
	database.ConnectDb()

	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views:       engine,
		ViewsLayout: "layouts/main",
	})

	setupRoutes(app)

	app.Static("/", "./public")

	app.Use(handlers.NotFound)

	app.Listen(":3051")
}
