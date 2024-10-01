package middlewares

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

const (
	StatusUnauthorized = fiber.StatusUnauthorized
	jwtSecretEnv       = "JWT_SECRET"
)

var jwtSecret = os.Getenv(jwtSecretEnv)

func AuthMiddleware(c *fiber.Ctx) error {
	tokenString := c.Get("Authorization")
	if tokenString == "" {
		return errorResponse(c, StatusUnauthorized, "No token provided")
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})

	if err != nil {
		return errorResponse(c, StatusUnauthorized, "Invalid token")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		c.Locals("user_id", uint(claims["user_id"].(float64)))
		return c.Next()
	}

	return errorResponse(c, StatusUnauthorized, "Invalid token")
}

func errorResponse(c *fiber.Ctx, statusCode int, message string) error {
	return c.Status(statusCode).JSON(fiber.Map{"error": message})
}
