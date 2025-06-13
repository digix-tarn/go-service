package config

import (
	"os"
)

var (
	AuthServiceURL string
	UserServiceURL string
)

func LoadConfig() {
	AuthServiceURL = os.Getenv("AUTH_SERVICE_URL")
	UserServiceURL = os.Getenv("USER_SERVICE_URL")
}
