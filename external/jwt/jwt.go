// pkg_jwt/jwt.go
package pkg_jwt

import (
	"fmt"
	"time"

	"strings"

	"github.com/PICH-IO/admin-api/configs"
	pkg_models "github.com/PICH-IO/admin-api/pkg/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

// GenerateJWT generates a JWT token with user claims (ID, username, role, and permissions)
func GenerateJWT(user *pkg_models.Token) (string, error) {
	// JWT claims
	claims := jwt.MapClaims{
		"id":          user.ID,
		"username":    user.Username,
		"role_id":     user.RoleID,
		"permissions": user.Permissions,
		"exp":         time.Now().Add(24 * time.Hour).Unix(),
	}

	// Generate JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(configs.JWT_SECRET))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ExtractToken extracts the token from the Authorization header
func ExtractToken(c *fiber.Ctx) (string, error) {
	bearerToken := c.Get("Authorization")
	if bearerToken == "" {
		return "", fmt.Errorf("missing Authorization header")
	}

	parts := strings.Split(bearerToken, " ")
	if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
		return "", fmt.Errorf("invalid Authorization header format")
	}

	return parts[1], nil
}

// ExtractTokenMetadata parses and validates the JWT token, returning user claims
func ExtractTokenMetadata(c *fiber.Ctx) (*pkg_models.Token, error) {
	tokenString, err := ExtractToken(c)
	if err != nil {
		return nil, err
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Ensure that the token method is "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(configs.JWT_SECRET), nil
	})

	if err != nil {
		return nil, err
	}

	// Validate and extract claims from the token
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		permissions, ok := claims["permissions"].([]interface{})
		if !ok {
			return nil, fmt.Errorf("invalid permissions format")
		}

		var permStrings []string
		for _, p := range permissions {
			if perm, ok := p.(string); ok {
				permStrings = append(permStrings, perm)
			}
		}

		return &pkg_models.Token{
			ID:          int(claims["id"].(float64)),
			Username:    claims["username"].(string),
			RoleID:      int(claims["role_id"].(float64)),
			Permissions: permStrings,
		}, nil
	}
	return nil, fmt.Errorf("invalid token")
}
