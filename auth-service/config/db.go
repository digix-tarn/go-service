package config

import (
	"fmt"
	"log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"my-microservice/models"
)

var DB *gorm.DB

func ConnectDatabase() {
	// แก้ไขค่าเหล่านี้ตามของคุณ
	dsn := "host=localhost user=postgres password=1234 dbname=go_system port=5432 sslmode=disable TimeZone=Asia/Bangkok"

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	fmt.Println("Connected to PostgreSQL!")

	// Auto migrate model
	err = database.AutoMigrate(&models.User{},&models.Profile{})
	if err != nil {
        log.Fatal("AutoMigrate failed:", err)
    }

	DB = database
}

type User struct {
	ID    uint   `gorm:"primaryKey"`
	Name  string
	Email string
}