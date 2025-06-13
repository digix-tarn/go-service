package routes

import (
	"bytes"
	"github.com/gofiber/fiber/v2"
	"io"
	"net/http"
	// "strings"
	"log/slog"
)

func ProxyHandler(targetBaseURL string, prefix string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		slog.Info("targetBaseURL => ", "value", targetBaseURL)
		slog.Info("prefix => ", "value", prefix)
		slog.Info("c.OriginalURL()[len(prefix):] => ", "value", c.OriginalURL()[len(prefix):])
		// path := strings.TrimPrefix(c.OriginalURL(), prefix)
		// slog.Info("trimmed path => ", "value", path)
		proxyURL := targetBaseURL + prefix + c.OriginalURL()[len(prefix):]
		slog.Info("proxyURL => ", "value", proxyURL)

		req, err := http.NewRequest(c.Method(), proxyURL, bytes.NewReader(c.Body()))
		if err != nil {
			return c.Status(500).SendString("Failed to create proxy request")
		}

		c.Request().Header.VisitAll(func(k, v []byte) {
			req.Header.Set(string(k), string(v))
		})

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return c.Status(500).SendString("Failed to reach backend")
		}
		defer resp.Body.Close()

		body, _ := io.ReadAll(resp.Body)
		return c.Status(resp.StatusCode).Send(body)
	}
}
