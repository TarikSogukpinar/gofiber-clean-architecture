package controller

import (
	"gofiber-clean-architecture/model"
	"gofiber-clean-architecture/service"

	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (controller *UserController) Route(app *fiber.App) {
	app.Post("/register", controller.Register)
	app.Post("/login", controller.Login)
}

func (controller *UserController) Register(c *fiber.Ctx) error {
	var request model.User
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	err := controller.userService.RegisterUser(c.Context(), request.Email, request.Password, "")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User registered successfully",
	})
}

func (controller *UserController) Login(c *fiber.Ctx) error {
	var request model.User
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	user, err := controller.userService.LoginUser(c.Context(), request.Email, request.Password)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid username or password",
		})
	}

	return c.Status(fiber.StatusOK).JSON(user)
}
