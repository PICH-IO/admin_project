package paymentmethods

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func GetPaymentMethod(app fiber.Router, db *sqlx.DB) {
	var dataResponse = NewPaymentHandler(db)
	app.Get("/payment-method", dataResponse.Show)
}
