// pkg_handlers/auth_handlers.go
package pkg_handlers

import (
	"github.com/PICH-IO/admin-api/pkg/jwt"
	"github.com/PICH-IO/admin-api/pkg/models"
	"github.com/PICH-IO/admin-api/pkg/services"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

// RegisterUser handles user registration
func RegisterUser(db *sqlx.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req struct {
			Username string `json:"username" validate:"required"`
			Email    string `json:"email" validate:"required,email"`
			Password string `json:"password" validate:"required,min=6"`
			RoleName string `json:"role_name" validate:"required"` // e.g., "Super Admin", "Sub Admin", etc.
		}

		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid request body",
			})
		}

		// Validate request fields as needed (omitted for brevity)

		// Hash the password
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to hash password",
			})
		}

		// Get role ID from role name
		roleID, err := services.GetRoleIDByName(db, req.RoleName)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid role name",
			})
		}

		// Create the user in the database
		user := &models.User{
			Username:     req.Username,
			Email:        req.Email,
			PasswordHash: string(hashedPassword),
			RoleID:       roleID,
		}

		userID, err := services.CreateUser(db, user)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to create user",
			})
		}

		user.UserID = userID

		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"message": "User registered successfully",
		})
	}
}

// LoginUser handles user login and JWT generation
func LoginUser(db *sqlx.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req struct {
			Username string `json:"username" validate:"required"`
			Password string `json:"password" validate:"required"`
		}

		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid request body",
			})
		}

		// Fetch user by username
		user, err := services.GetUserByUsername(db, req.Username)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid username or password",
			})
		}

		// Compare password
		if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid username or password",
			})
		}

		// Fetch permissions for the user's role
		permissions, err := services.FetchPermissionsForRole(db, user.RoleID)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to fetch permissions",
			})
		}

		// Create token
		tokenPayload := &models.Token{
			ID:          user.UserID,
			Username:    user.Username,
			RoleID:      user.RoleID,
			Permissions: permissions,
		}

		token, err := jwt.GenerateJWT(tokenPayload)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to generate token",
			})
		}

		return c.JSON(fiber.Map{
			"token": token,
		})
	}
}
