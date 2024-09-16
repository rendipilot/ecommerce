package routes

import (
	"e-commerce-synapsis/atom/products/controller"
	"e-commerce-synapsis/atom/shopping-cart/controller"
	"e-commerce-synapsis/atom/orders/controller"
	"e-commerce-synapsis/atom/users/controller"
	middlewares "e-commerce-synapsis/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func SetupRoutes() *fiber.App {

	app := fiber.New();

	app.Use(cors.New(cors.Config{
		AllowOrigins:  "http://localhost:3001",
		AllowMethods:  "POST, PUT, PATCH, DELETE, GET, OPTIONS, TRACE, CONNECT",
		AllowHeaders:  "Access-Control-Allow-Origin, Access-Control-Allow-Headers, Origin, Content-Type, Content-Length, Date, origin, Origins, x-requested-with, access-control-allow-methods, apikey, Authorization",
		ExposeHeaders: "Content-Length",
	}))


	app.Post("/register", users.UserRegister)
	app.Post("/login", users.UserLogin)

	protected := app.Group("")
	protected.Use(middlewares.JwtAuthMiddleware())

	productsList := protected.Group("products")
	{
		productsList.Post("category", products.GetProductByCategory)
	}

	cart := protected.Group("cart")
	{
		cart.Post("/add", shopping_cart.AddCart)
		cart.Post("/get-cart", shopping_cart.GetCartListByUserId)
		cart.Put("/delete", shopping_cart.DeleteCartById)
	}

	order := protected.Group("order")
	{
		order.Post("/new-order", orders.CheckoutOrder)
		order.Put("/checkout", orders.PaymentOrder)
	}


	return app
}