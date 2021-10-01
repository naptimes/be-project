package main

import (
	"be-project/router"

	"github.com/gofiber/fiber/middleware/cors"
)

func main() {
	router.Router()

	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))
}
