package pkg_jwt

import (
	"fmt"
	"strings"
	"time"

	"github.com/PICH-IO/admin-api/configs"
	pkg_models "github.com/PICH-IO/admin-api/pkg/models"
	util_response "github.com/PICH-IO/admin-api/pkg/utils/response"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

// Gernerate Jwt
func GenerateJWT(user *pkg_models.Token) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":          user.Id,
		"user_name":   user.Username,
		"role_id":     user.RoleId,
		"permissions": user.Permissions,
		"exp":         time.Now().Add(24 * time.Hour).Unix(),
	})

	tokenString, err := token.SignedString([]byte(configs.JWT_SECRET))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ExtractToken(c *fiber.Ctx) (string, error) {

	bearerToken := c.Get("Authorization")

	if bearerToken == "" {
		// var errorMessage = util_common.Translate(c, "MissingHeader")
		errHeader := util_response.HttpResponse(
			false,
			"missingHeader",
			400, //pkg_constants.MissingHeader,
			fiber.Map{
				"Errors": bearerToken,
			},
		)
		return "", c.JSON(errHeader)
	}

	strArr := strings.Split(bearerToken, " ")

	if len(strArr) != 2 || strings.ToLower(strArr[0]) != "bearer" {
		// var errorMessage = util_common.Translate(c, "InvalidAuthHeaderFormat")
		errHeader := util_response.HttpResponse(
			false,
			"InvalidAuthHeaderFormat",
			400, //constants.InvalidAuthHeaderFormat,
			fiber.Map{
				"Error": "InvalidAuthHeaderFormat",
			},
		)
		return "", c.JSON(errHeader)
	}
	return strArr[1], nil
}

func ExtractTokenMetadata(c *fiber.Ctx) (*pkg_models.Token, error) {

	var tokenString, errExtract = ExtractToken(c)
	if errExtract != nil {
		return nil, errExtract
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//** Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(configs.JWT_SECRET), nil
	})

	if err != nil {
		return nil, err
	}

	//**  Validate Token is expire
	claims, ok := token.Claims.(jwt.MapClaims)

	if ok && token.Valid {
		permissions, ok := claims["permissions"].([]interface{})
		if !ok {
			return nil, fmt.Errorf("invalid permissions format")
		}
		var permString []string
		for _, p := range permissions {
			if perm, ok := p.(string); ok {
				permString = append(permString, perm)
			}
		}

		return &pkg_models.Token{
			Id:          claims["id"].(float64),
			Username:    claims["user_name"].(string),
			RoleId:      claims["role_id"].(float64),
			Permissions: claims["permissions"].([]string),
		}, nil
	}
	return nil, fmt.Errorf("invalid token")
}
