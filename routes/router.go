package routes

import (
	"github.com/PICH-IO/admin-api/internal/auth"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func SetupRoutes(app *fiber.App, db *sqlx.DB) {

	root := app.Group("/api/v1")

	auth.SetupAuth(root, db)
}
