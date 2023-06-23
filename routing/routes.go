package routing

import (
	"web_crud/controller"

	"github.com/gofiber/fiber/v2"
)

func WebRoute(app *fiber.App) {

	app.Get("/", controller.ViewLogin)
	// USER
	endpointUser := app.Group("/user")
	endpointUser.Get("/register", controller.ViewRegistration)
	endpointUser.Post("/create", controller.CreateUser)
	endpointUser.Post("/verify", controller.VerifyAccount)

	endpointDashboard := app.Group("/dashboard")
	endpointDashboard.Get("/", controller.ViewDashboard)
	endpointDashboard.Get("/get/:id", controller.GetUser)
	endpointDashboard.Get("/getUsers", controller.GetUsers)
}
