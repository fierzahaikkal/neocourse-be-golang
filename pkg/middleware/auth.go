// pkg/middleware/auth_middleware.go
package middleware

import (
	"strings"

	"github.com/fierzahaikkal/neocourse-be-golang/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

func AuthMiddleware(jwtSecret string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		bearerToken := c.Get("Authorization")
		if bearerToken == "" {
			return utils.ErrorResponse(c, "Missing token", fiber.StatusUnauthorized)
		}

		tokenParts := strings.Split(bearerToken, " ")
		if len(tokenParts) != 2 || strings.ToLower(tokenParts[0]) != "bearer" {
			return utils.ErrorResponse(c, "Invalid token format", fiber.StatusUnauthorized)
		}

		tokenString := tokenParts[1]

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtSecret), nil
		})

		if err != nil || !token.Valid {
			return utils.ErrorResponse(c, "Invalid token", fiber.StatusUnauthorized)
		}

		// If you need to pass the token claims to the next handler, you can do:
		c.Locals("user", token.Claims)

		return c.Next()
	}
}