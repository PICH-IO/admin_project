package configs

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	DATABASE_URL string
	PORT         string
	JWT_SECRET   string
	USER_CONTEXT string
	TIME_ZONE    string
)

func InitConfig() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	DATABASE_URL = os.Getenv("DATABASE_URL")
	PORT = os.Getenv("PORT")
	JWT_SECRET = os.Getenv("JWT_SECRET")
	USER_CONTEXT = os.Getenv("USER_CONTEXT")
	TIME_ZONE = os.Getenv("TIME_ZONE")

	if PORT == "" && JWT_SECRET == "" {
		_ = fmt.Errorf("environment variable is not set")
	}
}
