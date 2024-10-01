package handlers

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/ajaysinghpanwar2002/pratilipi/cmd/user_service/internal/models"
	"github.com/ajaysinghpanwar2002/pratilipi/cmd/user_service/internal/services"
	"github.com/ajaysinghpanwar2002/pratilipi/pkg/rabbitmq"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

type UserHandler struct {
	service *services.UserService
}

func NewUserHandler(service *services.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) RegisterUser(c *fiber.Ctx) error {
	user := new(models.User)
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	if user.Password == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Password is required"})
	}

	if err := h.service.RegisterUser(user); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to register user"})
	}

	err := rabbitmq.EmitEvent("user_events", "UserRegistered", map[string]interface{}{
		"user_id":  user.ID,
		"username": user.Username,
		"email":    user.Email,
	})

	if err != nil {
		log.Printf("Failed to emit event: %v", err)
	}

	return c.Status(201).JSON(fiber.Map{"message": "User registered successfully", "user_id": user.ID})
}

func (h *UserHandler) LoginUser(c *fiber.Ctx) error {
	loginData := struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}{}

	if err := c.BodyParser(&loginData); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	// Authenticate user
	user, err := h.service.Authenticate(loginData.Username, loginData.Password)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{"error": err.Error()})
	}

	// Create JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to generate token"})
	}

	return c.Status(200).JSON(fiber.Map{"token": tokenString})
}

func (h *UserHandler) UpdateProfile(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint) // Extract user_id from JWT
	fmt.Println("userId", userID)

	var updateData map[string]interface{}
	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	if err := h.service.UpdateProfile(userID, updateData); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to update profile"})
	}

	err := rabbitmq.EmitEvent("user_events", "UserProfileUpdated", map[string]interface{}{
		"id": userID,
	})

	if err != nil {
		log.Printf("Failed to emit event: %v", err)
	}

	return c.JSON(fiber.Map{"message": "Profile updated successfully"})
}
