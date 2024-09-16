package users

import (
	"e-commerce-synapsis/atom/users"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

func UserLogin(c *fiber.Ctx) error {
	var data users.UserLoginRequest

	inputError := c.BodyParser(&data)
	if inputError != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "status": 400,
			"data": fiber.Map{},
			"message": "Invalid input",
        })
	}

	validate := validator.New()

	if err := validate.Struct(data); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  401,
			"data":    fiber.Map{},
			"message": "Invalid Credentials",
		})
	}

	res, err := users.UserLoginUseCase(data)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
            "status": 500,
			"data": fiber.Map{},
			"message": "Failed to login",
        })
	}

	return c.JSON(fiber.Map{
        "status": 200,
        "data": res,
        "message": "Login successful",
    })
}