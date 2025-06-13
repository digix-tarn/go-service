package routes

import (
	"github.com/gofiber/fiber/v2"
	"proxy-service/config"
)

func UserRoutes(app fiber.Router) {
	app.All("/*", ProxyHandler(config.UserServiceURL, "/api/user"))
}
