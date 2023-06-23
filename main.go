package main

import (
	"web_crud/middleware"
	"web_crud/model"
	"web_crud/routing"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/template/html/v2"
)

func main() {
	// Connect DB
	middleware.ConnectDB()
	// Create Table

	middleware.CreateTable(model.Users{})
	app := fiber.New(fiber.Config{
		Views: html.New("./template", ".html"),
	})

	// Configure application CORS
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	// Dito iloload ung path ng CSS,JS,IMAGES etc...
	app.Static("/", "./assets/css")
	app.Static("/", "./assets/js")
	app.Static("/", "./assets/images")

	routing.WebRoute(app)

	app.Listen(":3131")
}
