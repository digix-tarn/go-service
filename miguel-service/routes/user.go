package routes

import (
	"my-microservice/handlers"
	"my-microservice/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupUserRoutes(api fiber.Router) {
	user := api.Group("/user")
	user.Use(middleware.JWTProtected())
	user.Get("/list", handlers.GetUsers)
	user.Get("/:id", handlers.GetUserByID)
	user.Post("/create", handlers.CreateUser)
}
