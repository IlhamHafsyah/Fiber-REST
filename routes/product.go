package routes

import (
	"fiber-api/controllers"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func ProductRoute(route string, app *fiber.App) {
	app.Get(fmt.Sprintf("%v/:id", route), controllers.GetProductById)
	app.Get(route, controllers.GetProduct)
	app.Get(fmt.Sprintf("%v/get/withcategory", route), controllers.GetProductWihCategory)
}
