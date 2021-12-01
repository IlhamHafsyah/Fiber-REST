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
	app.Post(fmt.Sprintf("%v/insert", route), controllers.CreateProduct)
	app.Put(fmt.Sprintf("%v/update/:id", route), controllers.UpdateProduct)
	app.Delete(fmt.Sprintf("%v/delete/:id", route), controllers.DeleteProduct)
}
