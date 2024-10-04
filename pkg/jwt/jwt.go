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

// var tokeyKey = []byte("MT_SECRET")

// Gernerate Jwt
func GernerateJWT(player *pkg_models.Token) (string, error) {
	var token = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":        player.Id,
		"user_name": player.Username,
		// "membership_id":   player.MembershipId,
		// "membership_role": player.MembershipRole,
		// "role_id":         player.RoleId,
		"exp": time.Now().Add(24 * time.Hour).Unix(),
	})

	tokenString, err := token.SignedString([]byte(configs.JWT_SECRET))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ExtractTokenMetadata(c *fiber.Ctx) (*pkg_models.Token, error) {
	var tokenString, _ = ExtractToken(c)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//** Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("secret_mt_api"), nil
	})

	//**  Validate Token is expire
	claims, ok := token.Claims.(jwt.MapClaims)
	// fmt.Println(claims)

	if ok && token.Valid {
		return &pkg_models.Token{
			Id:       claims["id"].(float64),
			Username: claims["user_name"].(string),
			// MembershipId:   claims["membership_id"].(float64),
			// MembershipRole: claims["membership_role"].(string),
			// RoleId:         claims["role_id"].(float64),
		}, nil
	}
	return nil, err
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
