package shopping_cart

import (
	shopping_cart "e-commerce-synapsis/atom/shopping-cart"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

func GetCartListByUserId(c *fiber.Ctx) error {
	var data shopping_cart.GetCartRequest

	inputError := c.BodyParser(&data)
	if inputError != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  400,
			"data":    nil,
			"message": "Invalid input",
		})
	}

	validate := validator.New()

	if err := validate.Struct(data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  400,
			"data":    nil,
			"message": "Invalid input",
		})
	}

	res, err := shopping_cart.GetCartUseCase(data.UserID)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  401,
			"data":    nil,
			"message": "Failed to get products list",
		})
	}

	return c.JSON(fiber.Map{
		"status":  200,
		"data":    res,
		"message": "Get cart successful",
	})

}