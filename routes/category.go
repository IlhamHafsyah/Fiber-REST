package routes

import (
	"fiber-api/controllers"

	"github.com/gofiber/fiber/v2"
)

func CategoryRoute(route string, app *fiber.App) {
	app.Get(route, controllers.GetCategory)
}
