package middlewares

import (
	utils_token "e-commerce-synapsis/utils"

	"github.com/gofiber/fiber/v2"
)

func JwtAuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		err := utils_token.TokenValid(c)

		if err != nil {
			c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized",
			})
		}
		return c.Next()
	}
}