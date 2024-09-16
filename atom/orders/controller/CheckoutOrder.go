package orders

import (
	"e-commerce-synapsis/atom/orders"
	"log"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

func CheckoutOrder(c *fiber.Ctx) error {
	var data orders.OrderRequest

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
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  401,
			"data":    nil,
			"message": "Invalid credentials",
		})
	}

	res, err := orders.CheckoutOrderUseCase(data.UserID)

	if err == nil && res == 0 {
		log.Println("no contect")
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
            "status":  200,
            "data":    nil,
            "message": "No order found",
        })
	}

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  500,
			"data":    nil,
			"message": "Failed to add order",
		})
	}

	return c.JSON(fiber.Map{
		"status":  200,
		"data":    res,
		"message": "successful add order",
	})
}