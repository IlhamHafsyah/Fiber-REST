package controllers

import (
	"fiber-api/config"
	"fiber-api/models"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm/clause"
)

func GetProduct(c *fiber.Ctx) error {
	fmt.Println("jalan")
	var product []models.Product

	config.DB.Preload(clause.Associations).Find(&product)

	return c.JSON(&product)
}

func GetProductById(c *fiber.Ctx) error {
	fmt.Println("jalan by id")

	id, _ := strconv.Atoi(c.Params("id"))
	fmt.Println(id)
	product := models.Product{ProductID: uint(id)}

	config.DB.Preload(clause.Associations).Find(&product)

	return c.JSON(&product)
}

func GetProductWihCategory(c *fiber.Ctx) error {
	fmt.Println("jalan")

	type Result struct {
		Name         string
		CategoryDesc string
	}

	var result Result
	var productwihtCategory models.Product

	config.DB.Find(&productwihtCategory).Select("product.name, category.category_desc").Joins("left join hasycoffee.category on product.id_category = category.category_id").Scan(&result)

	return c.JSON(&result)
}
