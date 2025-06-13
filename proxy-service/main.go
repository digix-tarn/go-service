package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"proxy-service/routes"
	"proxy-service/config"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	config.LoadConfig()

	app := fiber.New()

	api := app.Group("/api")
	for _, group := range routes.RouteGroups {
		group.Handler(api.Group(group.Path))
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	log.Fatal(app.Listen(":" + port))
}