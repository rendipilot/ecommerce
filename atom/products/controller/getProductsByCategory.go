package products

import (
	"e-commerce-synapsis/atom/products"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

func GetProductByCategory(c *fiber.Ctx) error {
	var data products.ProductRequest

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

	res, err := products.GetProductByCategoryUseCase(data.Category)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "status": 500,
			"data": fiber.Map{},
			"message": "Failed to get products list",
        })
	}

	return c.JSON(fiber.Map{
        "status": 200,
        "data": res,
        "message": "Get product successful",
    })
}