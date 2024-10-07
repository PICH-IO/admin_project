package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func SetupAuth(app fiber.Router, db *sqlx.DB) {
	repo := NewAuthReposity(db)
	service := NewAuthService(repo)
	handler := NewAuthHandler(service)

	app.Post("/auth/login", handler.Login)
}
