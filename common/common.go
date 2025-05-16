package common

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// LoadEnv loads the environment variables from the .env file
func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️ .env file not found — assuming environment variables are passed by Docker Compose")
	}
}

var JwtSecret = []byte(os.Getenv("JWT_SECRET"))
