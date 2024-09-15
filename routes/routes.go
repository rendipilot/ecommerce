package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func SetupRoutes() *fiber.App {

	app := fiber.New();

	app.Use(cors.New(cors.Config{
		AllowOrigins:  "http://localhost:3001,*",
		AllowMethods:  "POST, PUT, PATCH, DELETE, GET, OPTIONS, TRACE, CONNECT",
		AllowHeaders:  "Access-Control-Allow-Origin, Access-Control-Allow-Headers, Origin, Content-Type, Content-Length, Date, origin, Origins, x-requested-with, access-control-allow-methods, apikey, Authorization",
		ExposeHeaders: "Content-Length",
	}))


	return app
}