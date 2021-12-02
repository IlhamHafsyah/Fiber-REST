package controllers

import (
	"fiber-api/config"
	"fiber-api/models"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm/clause"
)

func GetCoupon(c *fiber.Ctx) error {
	var coupon []models.Coupon

	config.DB.Preload(clause.Associations).Find(&coupon)

	return c.JSON(&fiber.Map{
		"status": "success",
		"code":   200,
		"data":   &coupon,
	})
}

func CreateCoupon(c *fiber.Ctx) error {
	var newCoupon models.Coupon

	err := c.BodyParser(&newCoupon)
	if err != nil {
		return nil
	}

	res := config.DB.Create(&newCoupon)
	if res.Error != nil || res.RowsAffected <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "error",
			"message": res.Error,
		})
	}

	return c.JSON(&fiber.Map{
		"status":  "success",
		"message": "Coupon Created",
		"data":    newCoupon,
	})
}

func GetCouponById(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	coupon := models.Coupon{CouponId: int(id)}
	config.DB.Preload(clause.Associations).Find(&coupon)
	var result = [1]models.Coupon{coupon}
	if result[0].CouponName == "" {
		return c.JSON((&fiber.Map{
			"status":  "seccess",
			"message": fmt.Sprintf("Coupon with id: %v not found", id),
			"code":    404,
		}))
	} else {
		return c.JSON(&fiber.Map{
			"status":  "success",
			"message": fmt.Sprintf("Success get coupon with id: %v", id),
			"code":    200,
			"data":    result,
		})
	}
}

func UpdateCoupon(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	coupon := models.Coupon{CouponId: int(id)}

	err := c.BodyParser(&coupon)
	if err != nil {
		return err
	}

	config.DB.Preload(clause.Associations).Updates(&coupon).Find(&coupon)

	return c.JSON(&fiber.Map{
		"status":  "success",
		"message": fmt.Sprintf("Coupon with id: %v id updated", id),
		"code":    200,
		"data":    coupon,
	})
}

func DeleteCoupon(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	coupon := models.Coupon{CouponId: int(id)}

	if coupon.CouponName == "" {
		return c.JSON(&fiber.Map{
			"status":  "success",
			"message": fmt.Sprintf("Coupon with id: %v does not exist", id),
			"code":    404,
		})
	} else {
		config.DB.Delete(&coupon)

		return c.Status(fiber.StatusNoContent).JSON(&fiber.Map{
			"message": fmt.Sprintf("Coupon with id: %v id deleted", id),
		})
	}

}
