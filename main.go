package main

import (
	"fiber-api/config"
	"fiber-api/routes"
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

//set base uri
const baseUrl string = "/api"

func main() {

	// load env
	err := godotenv.Load()
	if err != nil {
		panic(fmt.Errorf("fatal error locading .env: %w ", err))
	}

	// connect DB
	config.DbConnect()

	// create app
	app := fiber.New()

	// sample middleware
	app.Use(cors.New())
	app.Use(logger.New())

	// simple route
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(&fiber.Map{
			"info": "GoLang, Fiber, Postgres API",
		})
	})

	// routes navigation
	routes.RoutesNavigation(baseUrl, app)

	// listen
	err = app.Listen(":" + os.Getenv("PORT"))
	if err != nil {
		panic(fmt.Errorf("failed to listen on port %s: %w ", os.Getenv("PORT"), err))
	}
}
