package main

import (
	"log"
	"my-microservice/routes"
	"my-microservice/config"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"

)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app := fiber.New()
	config.ConnectDatabase()
	
	routes.SetupRoutes(app)

	app.Listen(":5001")
}
