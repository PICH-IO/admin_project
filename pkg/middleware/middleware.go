package pkg_middleware

import (
	"errors"
	"fmt"

	"github.com/PICH-IO/admin-api/configs"
	pkg_jwt "github.com/PICH-IO/admin-api/pkg/jwt"
	util_error "github.com/PICH-IO/admin-api/pkg/utils/errors"
	util_response "github.com/PICH-IO/admin-api/pkg/utils/response"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt/v4"
)

func JWTAuthMiddleware(secret string) fiber.Handler {

	return func(c *fiber.Ctx) error {

		if secret == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(util_error.ErrorResponse{
				ErrorMessage: "Warning: Empty secret provided for JWT authentication",
			})
		}

		config := jwtware.Config{
			SigningKey: []byte(secret),
			ContextKey: configs.USER_CONTEXT,
			AuthScheme: "Bearer",
			ErrorHandler: func(c *fiber.Ctx, err error) error {
				_, ok := pkg_jwt.ExtractToken(c)

				if ok != nil {
					return c.Status(fiber.StatusUnauthorized).JSON(util_error.ErrorResponse{
						ErrorMessage: ok.Error(),
					})
				}

				if errors.Is(err, jwt.ErrTokenExpired) {
					// var errorMessage = util_common.Translate(c, "TokenExpired")
					errToken := util_response.HttpResponse(
						false,
						"TokenExpired",
						401, //constants.TokenExpired,
						fiber.Map{
							"Error": "error",
						},
					)
					return c.Status(fiber.StatusUnauthorized).JSON(errToken)
				} else {
					// var errorMessage = util_common.Translate(c, "TokenInvalide")
					errToken := util_response.HttpResponse(
						false,
						"TokenInvalide",
						401, //constants.TokenInvalide,
						fiber.Map{
							"Error": "TokenInvalide",
						},
					)
					return c.Status(fiber.StatusUnauthorized).JSON(errToken)
				}
			},
			KeyFunc: func(token *jwt.Token) (interface{}, error) {
				if secret == "" {
					return nil, fmt.Errorf("empty secret provided")
				}
				// Verify signing method
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
				}
				// Ensure the key is correct
				if []byte(secret) == nil {
					return nil, errors.New("token is invalid")
				}
				return []byte(secret), nil
			},
			SuccessHandler: func(c *fiber.Ctx) error {
				var claims, ok = pkg_jwt.ExtractTokenMetadata(c)
				if ok != nil {
					return c.Status(fiber.StatusUnauthorized).JSON(util_error.ErrorResponse{
						ErrorMessage: ok.Error(),
					})
				}
				c.Locals(configs.USER_CONTEXT, claims)
				return c.Next()
			},
		}

		// Create the JWT middleware with the configured options
		jwtMiddleware := jwtware.New(config)

		return jwtMiddleware(c)
	}
}
