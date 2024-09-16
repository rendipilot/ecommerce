package shopping_cart

import (
	shopping_cart "e-commerce-synapsis/atom/shopping-cart"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

func DeleteCartById(c *fiber.Ctx) error {
	var data shopping_cart.DeleteCartRequest

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
			"message": "Invalid credentials",
		})
	}

	res, err := shopping_cart.DeleteCartByIdUseCase(data)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "status": 500,
			"data": fiber.Map{},
			"message": "Failed to delete cart list",
        })
	}

	return c.JSON(fiber.Map{
        "status": 200,
        "data": res,
        "message": "Delete Cart successful",
    })

}