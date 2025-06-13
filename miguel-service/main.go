package main

import (
	"log"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	// "github.com/gofiber/websocket/v2"
	"miguel-service/handlers"
	"miguel-service/middleware"
	"miguel-service/redis"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app := fiber.New()

	redis.InitRedis()

	app.Use("/ws", middleware.JWTWebSocketGuard())
	handlers.SetupWebSocket(app)

	app.Listen(":5003")
}


// package main

// import (
// 	"log"
// 	"os"

// 	"github.com/gofiber/fiber/v2"
// 	"github.com/gofiber/websocket/v2"
// 	"github.com/joho/godotenv"
// 	"miguel-service/handlers"
// )

// func main() {
// 	_ = godotenv.Load()
// 	jwtSecret := os.Getenv("JWT_SECRET")
// 	if jwtSecret == "" {
// 		jwtSecret = "mysecretkey123"
// 	}

// 	app := fiber.New()

// 	app.Use("/ws", func(c *fiber.Ctx) error {
// 		token := c.Query("token")
// 		if token == "" {
// 			return fiber.ErrUnauthorized
// 		}

// 		claims, err := handlers.ParseJWT(token, jwtSecret)
// 		if err != nil {
// 			return fiber.NewError(fiber.StatusUnauthorized, "Invalid token")
// 		}

// 		c.Locals("userID", claims["sub"]) // สมมุติว่า JWT มี `sub` เป็น userID
// 		return c.Next()
// 	})

// 	app.Get("/ws", websocket.New(handlers.WebSocketHandler))

// 	log.Println("Miguel Service running on :5003")
// 	log.Fatal(app.Listen(":5003"))
// }
