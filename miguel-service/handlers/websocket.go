package handlers

import (
	"log/slog"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"miguel-service/middleware"
	"miguel-service/redis"
)

type Message struct {
	Sender string `json:"sender"`
	Text   string `json:"text"`
}

// slog.Info("JWT Secret", "value", jwtSecret)

var (
	clients = make(map[*websocket.Conn]string)
	broadcast = make(chan string)
)

func SetupWebSocket(app *fiber.App) {
	// subscribe redis only once
	go func() {
		redis.Subscribe("chat", func(msg string) {
			broadcast <- msg
		})
	}()

	app.Use("/ws", func(c *fiber.Ctx) error {
		token := c.Query("token")
		if token == "" {
			return fiber.ErrUnauthorized
		}
		claims, err := middleware.ParseJWT(token)
		slog.Info("claims", "value", claims)
		if err != nil {
			return fiber.ErrUnauthorized
		}
		slog.Info("claims", "value", claims["user_id"])
		c.Locals("userID", claims["user_id"])
		return c.Next()
	})

	app.Get("/ws", websocket.New(func(c *websocket.Conn) {
		slog.Info("c.Locals('userID')", "value", c.Locals("userID"))
		userIDRaw, ok := c.Locals("userID").(float64) // ถ้ามาจาก JWT แบบ numeric
		if !ok {
			c.Close()
			return 
		}
		slog.Info("userIDRaw", "value", userIDRaw)
		userID := fmt.Sprintf("%.0f", userIDRaw)
		slog.Info("userID", "value", userID)
		if userID == "" {
			c.Close()
			return 
		}
		slog.Info("userID", "value", userID)
		clients[c] = userID
		defer func() {
			delete(clients, c)
			c.Close()
		}()

		// Start a goroutine to listen on broadcast channel and send to this client
		go func() {
			for {
				msg := <-broadcast
				c.WriteMessage(websocket.TextMessage, []byte(msg))
			}
		}()

		for {
			_, msg, err := c.ReadMessage()
			if err != nil {
				break
			}
			data := Message{
				Sender: userID,
				Text:   string(msg),
			}
			payload, _ := json.Marshal(data)
			redis.Publish("chat", string(payload))
		}
	}))
}
