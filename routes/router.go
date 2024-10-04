package routes

import (
	paymentmethods "github.com/PICH-IO/admin-api/internal/payment"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func SetupRoutes(app *fiber.App, db *sqlx.DB) {
	root := app.Group("/api/v1")
	paymentmethods.GetPaymentMethod(root, db)
}
