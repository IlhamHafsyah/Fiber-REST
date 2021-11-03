package routes

import (
	"fiber-api/controllers"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func Setup(base string, app *fiber.App) {
	route := app.Group(fmt.Sprintf("%v/products", base))

	route.Get("/category", controllers.GetCategory)
	route.Get("/:id", controllers.GetProductById)
	route.Get("/", controllers.GetProduct)
	route.Get("/productwithcategory/", controllers.GetProductWihCategory)
	// app.Post("/users", controllers.CreateUsers)
	// app.Put("/users/:id", controllers.UpdateUsers)
	// app.Delete("/users/:id", controllers.DeleteUsers)
}
