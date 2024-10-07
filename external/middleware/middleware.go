// pkg_middleware/auth_middleware.go
package pkg_middleware

import (
	"fmt"

	"github.com/PICH-IO/admin-api/configs"
	pkg_jwt "github.com/PICH-IO/admin-api/pkg/jwt"
	pkg_models "github.com/PICH-IO/admin-api/pkg/models"
	"github.com/gofiber/fiber/v2"
)

// JWTAuthMiddleware authenticates the user by validating the JWT token
func JWTAuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		token, err := pkg_jwt.ExtractTokenMetadata(c)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized: " + err.Error(),
			})
		}

		// Store user information in context for later use
		c.Locals(configs.USER_CONTEXT, token)
		return c.Next()
	}
}

// PermissionMiddleware checks if the user has the required permissions
func PermissionMiddleware(requiredPermissions ...string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Retrieve user from context (set by JWTAuthMiddleware)
		user, ok := c.Locals(configs.USER_CONTEXT).(*pkg_models.Token)
		if !ok || user == nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized: User not authenticated",
			})
		}

		// Check if user has all required permissions
		for _, perm := range requiredPermissions {
			if !hasPermission(user.Permissions, perm) {
				return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
					"error": fmt.Sprintf("Forbidden: Missing permission '%s'", perm),
				})
			}
		}

		return c.Next()
	}
}

// hasPermission checks if the user's permissions include the required permission
func hasPermission(userPermissions []string, required string) bool {
	for _, perm := range userPermissions {
		if perm == required {
			return true
		}
	}
	return false
}
