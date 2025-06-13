package middleware

import (
	"os"
	"log/slog"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)

func JWTProtected() fiber.Handler {
	slog.Info("JWT Secret mid", "value", os.Getenv("JWT_SECRET"))
	return jwtware.New(jwtware.Config{
		// SigningKey: []byte("secret123"),
		SigningKey: []byte(os.Getenv("JWT_SECRET")),
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized",
			})
		},
	})
}
