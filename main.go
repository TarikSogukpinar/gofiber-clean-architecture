package main

import (
	"context"
	config "gofiber-clean-architecture/configuration"
	"gofiber-clean-architecture/database"
	"gofiber-clean-architecture/handler"
	"gofiber-clean-architecture/repository"
	"gofiber-clean-architecture/service"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {
	config.LoadConfig()

	if err := database.Connect(); err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := database.Mg.Client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	// Setup repositories, services, and handlers
	userRepo := repository.NewUserRepository(database.UserCollection.Database())
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	app := fiber.New()

	userHandler.RegisterRoutes(app)

	port := os.Getenv("PORT")
	if port == "" {
		port = "6060" // Default port
	}

	log.Fatal(app.Listen(":" + port))
}
