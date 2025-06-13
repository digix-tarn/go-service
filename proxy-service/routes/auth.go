package routes

import (
	"github.com/gofiber/fiber/v2"
	"proxy-service/config"
	"log/slog"
)

func AuthRoutes(app fiber.Router) {
	slog.Info("config => ", "value", config.AuthServiceURL)
	app.All("/*", ProxyHandler(config.AuthServiceURL, "/api/auth"))
}
