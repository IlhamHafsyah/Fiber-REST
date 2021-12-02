package routes

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func RoutesNavigation(base string, app *fiber.App) {
	CategoryRoute(fmt.Sprintf("%v/category", base), app)
	ProductRoute(fmt.Sprintf("%v/product", base), app)
	CouponRoute(fmt.Sprintf("%v/coupon", base), app)
}
