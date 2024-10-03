package middlewares

import (
	"os"
	"strconv"

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
		userID, err := getUserIDFromClaims(claims)
		if err != nil {
			return errorResponse(c, StatusUnauthorized, "Invalid user ID in token")
		}
		c.Locals("user_id", userID)
		return c.Next()
	}

	return errorResponse(c, StatusUnauthorized, "Invalid token")
}

func getUserIDFromClaims(claims jwt.MapClaims) (string, error) {
	switch id := claims["user_id"].(type) {
	case string:
		return id, nil
	case float64:
		return strconv.FormatFloat(id, 'f', 0, 64), nil
	case int:
		return strconv.Itoa(id), nil
	case uint:
		return strconv.FormatUint(uint64(id), 10), nil
	default:
		return "", fiber.NewError(fiber.StatusUnauthorized, "Invalid user ID type in token")
	}
}

func errorResponse(c *fiber.Ctx, statusCode int, message string) error {
	return c.Status(statusCode).JSON(fiber.Map{"error": message})
}
