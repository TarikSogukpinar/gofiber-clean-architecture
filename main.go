package main

import (
	"gofiber-clean-architecture/configuration"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {
	config := configuration.New()
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World! Go Fiber")
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "6060" // Default port
	}

	log.Fatal(app.Listen(":" + port))
}
