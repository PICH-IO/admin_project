package configs

import (
	"cmp"
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

	DATABASE_URL = cmp.Or(os.Getenv("DATABASE_URL"), "DATABASE_URL")
	PORT = cmp.Or(os.Getenv("PORT"), "PORT")
	JWT_SECRET = cmp.Or(os.Getenv("JWT_SECRET"), "THEISIS_SECRET")
	TIME_ZONE = cmp.Or(os.Getenv("TIME_ZONE"), "Asia/Phnom_Penh")
	USER_CONTEXT = cmp.Or(os.Getenv("USER_CONTEXT"), "userContext")

	if PORT == "" && JWT_SECRET == "" {
		_ = fmt.Errorf("environment variable is not set")
	}
}
