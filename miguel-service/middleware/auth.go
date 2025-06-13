package middleware

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"os"
)

func ParseJWT(tokenStr string) (jwt.MapClaims, error) {
	secret := os.Getenv("JWT_SECRET")
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil || !token.Valid {
		return nil, errors.New("invalid JWT token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid JWT claims")
	}

	return claims, nil
}

func JWTWebSocketGuard() fiber.Handler {
	return func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			token := c.Query("token")
			if token == "" {
				return fiber.ErrUnauthorized
			}
			claims, err := ParseJWT(token)
			if err != nil {
				return fiber.ErrUnauthorized
			}
			c.Locals("userID", claims["sub"])
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	}
}