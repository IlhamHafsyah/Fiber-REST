package routes

import (
	"fiber-api/controllers"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func CouponRoute(route string, app *fiber.App) {
	app.Get(route, controllers.GetCoupon)
	app.Post(fmt.Sprintf("%v/insert", route), controllers.CreateCoupon)
	app.Get(fmt.Sprintf("%v/:id", route), controllers.GetCouponById)
	app.Put(fmt.Sprintf("%v/update/:id", route), controllers.UpdateCoupon)
	app.Delete(fmt.Sprintf("%v/delete/:id", route), controllers.DeleteCoupon)
}
