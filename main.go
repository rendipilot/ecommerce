package main

import (
	"e-commerce-synapsis/routes"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	ri := godotenv.Load();

	if ri != nil {
		log.Fatal(ri)
	}

	app := routes.SetupRoutes()
	app.Listen(":8080")
}