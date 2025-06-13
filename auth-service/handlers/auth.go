package handlers

import (
	"time"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
    // "github.com/joho/godotenv"
    "os"
    "log/slog"

	"my-microservice/requests"
	"my-microservice/models"
	"my-microservice/config"
	"my-microservice/utils"
)

// jwtSecret := os.Getenv("JWT_SECRET")

// var jwtSecret = os.Getenv("JWT_SECRET") // ควรเก็บใน env vars

func Login(c *fiber.Ctx) error {
    var input requests.LoginInput
    if err := c.BodyParser(&input); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Cannot parse JSON",
        })
    }

    var user models.User
    if err := config.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
            "error": "Invalid email or password",
        })
    }

    // ตรวจสอบ password
    if err := utils.CheckPasswordHash(input.Password, user.Password); err != nil {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
            "error": "Invalid email or password",
        })
    }

    // ถ้าอยากทำ JWT token เพื่อใช้ session ต่อได้ (แนะนำทำ)
    claims := jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 72).Unix(), // token หมดอายุ 72 ชั่วโมง
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    jwtSecret := os.Getenv("JWT_SECRET")
    slog.Info("JWT Secret", "value", jwtSecret)
	t, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not login",
		})
	}

    // return c.JSON(fiber.Map{
    //     "message": "Login success",
    //     "user":    user,
    // })

	return c.JSON(fiber.Map{
		"token": t,
	})
}