// pkg_handlers/super_admin_handlers.go
package pkg_handlers

import (
	"github.com/PICH-IO/admin-api/pkg/models"
	"github.com/PICH-IO/admin-api/pkg/services"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

// AddSubAdmin handles adding a new Sub Admin
func AddSubAdmin(db *sqlx.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req struct {
			Username string `json:"username" validate:"required"`
			Email    string `json:"email" validate:"required,email"`
			Password string `json:"password" validate:"required,min=6"`
		}

		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid request body",
			})
		}

		// Hash password
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to hash password",
			})
		}

		// Get Sub Admin role ID
		roleID, err := services.GetRoleIDByName(db, "Sub Admin")
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to get role ID",
			})
		}

		// Create user
		user := &models.User{
			Username:     req.Username,
			Email:        req.Email,
			PasswordHash: string(hashedPassword),
			RoleID:       roleID,
		}

		userID, err := services.CreateUser(db, user)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to create Sub Admin",
			})
		}

		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"message": "Sub Admin created successfully",
			"user_id": userID,
		})
	}
}

// GetAllSubAdmins retrieves all Sub Admins
func GetAllSubAdmins(db *sqlx.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		subAdmins, err := services.GetUsersByRoleName(db, "Sub Admin")
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to retrieve Sub Admins",
			})
		}

		return c.JSON(subAdmins)
	}
}

// UpdateSubAdmin handles updating a Sub Admin's information
func UpdateSubAdmin(db *sqlx.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		subAdminID := c.Params("id")
		var req struct {
			Email    string `json:"email" validate:"email"`
			Password string `json:"password" validate:"min=6"`
		}

		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid request body",
			})
		}

		// Fetch existing user
		user, err := services.GetUserByID(db, subAdminID)
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Sub Admin not found",
			})
		}

		// Update fields if provided
		if req.Email != "" {
			user.Email = req.Email
		}
		if req.Password != "" {
			hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"error": "Failed to hash password",
				})
			}
			user.PasswordHash = string(hashedPassword)
		}

		// Update user in database
		if err := services.UpdateUser(db, user); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to update Sub Admin",
			})
		}

		return c.JSON(fiber.Map{
			"message": "Sub Admin updated successfully",
		})
	}
}

// DeleteSubAdmin handles deleting or blocking a Sub Admin
func DeleteSubAdmin(db *sqlx.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		subAdminID := c.Params("id")

		// Delete or block logic (soft delete example)
		if err := services.SoftDeleteUser(db, subAdminID); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to delete Sub Admin",
			})
		}

		return c.JSON(fiber.Map{
			"message": "Sub Admin deleted successfully",
		})
	}
}
