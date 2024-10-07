package auth

import (
	"github.com/gofiber/fiber/v2"
)

type AuthHandler interface {
	Login(c *fiber.Ctx) error
}

type authHandler struct {
	service AuthService
}

func NewAuthHandler(service AuthService) AuthHandler {
	return &authHandler{service: service}
}

func (h *authHandler) Login(c *fiber.Ctx) error {
	var req UserReq
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	// Call the service layer for login
	userRes, err := h.service.Login(req)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": err.ErrorMessage})
	}

	// Return the JWT token on successful login
	return c.Status(fiber.StatusOK).JSON(userRes)
}
