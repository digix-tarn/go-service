package routes

import (
	"my-microservice/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupAuthRoutes(api fiber.Router) {
	auth := api.Group("/auth")
	auth.Post("/login", handlers.Login)
}
