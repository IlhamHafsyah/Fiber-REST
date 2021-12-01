package controllers

import (
	"fiber-api/config"
	"fiber-api/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm/clause"
)

func GetProduct(c *fiber.Ctx) error {
	var product []models.Product

	config.DB.Preload(clause.Associations).Find(&product)

	return c.JSON(&product)
}

func GetProductById(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	product := models.Product{ProductID: uint(id)}

	config.DB.Preload(clause.Associations).Find(&product)

	return c.JSON(&product)
}

func GetProductWihCategory(c *fiber.Ctx) error {
	type Result struct {
		Name         string
		CategoryDesc string
	}

	var result []Result
	var productwihtCategory models.Product

	config.DB.Model(&productwihtCategory).Select("product.name, category.category_desc").Joins("left join hasycoffee.category on product.id_category = category.category_id").Scan(&result)

	return c.JSON(&result)
}

func CreateProduct(c *fiber.Ctx) error {
	var newProduct models.Product

	err := c.BodyParser(&newProduct)
	if err != nil {
		return nil
	}

	res := config.DB.Create(&newProduct)
	if res.Error != nil || res.RowsAffected <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "error",
			"message": res.Error,
		})
	}

	return c.JSON(&fiber.Map{
		"status":  "success",
		"message": "Product Created",
		"data":    newProduct,
	})
}

func UpdateProduct(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	product := models.Product{ProductID: uint(id)}

	err := c.BodyParser(&product)
	if err != nil {
		return err
	}

	config.DB.Preload(clause.Associations).Updates(&product).Find(&product)

	return c.JSON(&fiber.Map{
		"message": "Product Updated",
		"data":    product,
	})
}

func DeleteProduct(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	product := models.Product{ProductID: uint(id)}

	config.DB.Delete(&product)

	return c.Status(fiber.StatusNoContent).JSON(&fiber.Map{
		"message": "Users Deleted",
	})
}
