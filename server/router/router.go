package router

import (
	"github.com/Gilgammesh/go-react-crud/controllers"
	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App) {
	// Default route
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Server run with GO")
	})

	// Base path
	apiGroup := app.Group("/api")

	// Users
	apiGroup.Get("/users", controllers.GetUsers)
	apiGroup.Get("/users/:id", controllers.GetUser)
	apiGroup.Post("/users", controllers.CreateUser)
	apiGroup.Put("/users/:id", controllers.UpdateUser)
	apiGroup.Delete("/users/:id", controllers.DeleteUser)
}
