package users

import (
	"e-commerce-synapsis/atom/users"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

func UserRegister(c *fiber.Ctx) error {
	var data users.UserRegisterRequest

	inputError := c.BodyParser(&data)
	if inputError != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "status": 400,
			"data": fiber.Map{},
			"message": "Invalid Input",
        })
	}

	validate := validator.New()

	if err := validate.Struct(data); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  401,
			"data":    fiber.Map{},
			"message": "Invalid credentials",
		})
	}

	res, err := users.UserRegisterUseCase(data)

	if res == "exist" {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
            "status": 409,
            "data": fiber.Map{},
            "message": "Email already registered",
        })
	}

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "status": 500,
			"data": fiber.Map{},
			"message": "Failed to register",
        })
	}

	return c.JSON(fiber.Map{
        "status": 200,
        "data": res,
        "message": "Register successful",
    })
}