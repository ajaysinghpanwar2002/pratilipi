package services

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/ajaysinghpanwar2002/pratilipi/cmd/order_service/internal/models"
	repositories "github.com/ajaysinghpanwar2002/pratilipi/cmd/order_service/internal/repository"
)

type UserService struct {
	userRepo *repositories.UserRepository
}

func NewUserService(userRepo *repositories.UserRepository) *UserService {
	return &UserService{userRepo}
}

func (s *UserService) HandleUserRegisteredEvent(data map[string]interface{}) error {
	// Extract the data and safely cast each field
	userID, ok := data["user_id"].(string)
	if !ok {
		if idFloat, ok := data["user_id"].(float64); ok {
			userID = fmt.Sprintf("%.0f", idFloat)
		} else {
			return fmt.Errorf("invalid data type for user_id")
		}
	}

	username, ok := data["username"].(string)
	if !ok {
		return fmt.Errorf("invalid data type for username")
	}

	email, ok := data["email"].(string)
	if !ok {
		return fmt.Errorf("invalid data type for email")
	}

	// Add user to the local users table
	now := time.Now()

	user := models.User{
		ID:        userID,
		Username:  username,
		Email:     email,
		CreatedAt: now,
		UpdatedAt: now,
	}

	if err := s.userRepo.CreateUser(user); err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}
	return nil
}

func (s *UserService) HandleUserProfileUpdatedEvent(data map[string]interface{}) error {
	userID, ok := data["id"].(string)
	if !ok {
		return fmt.Errorf("invalid data type for user_id")
	}

	// Assuming the event sends updated fields (e.g., username, email)
	updates := map[string]interface{}{}

	if username, ok := data["username"]; ok {
		updates["username"] = username
	}
	if email, ok := data["email"]; ok {
		updates["email"] = email
	}

	if err := s.userRepo.UpdateUser(userID, updates); err != nil {
		return fmt.Errorf("failed to update user profile: %w", err)
	}
	return nil
}

func (s *UserService) GetUserByID(ctx context.Context, userID string) (*models.User, error) {
	user, err := s.userRepo.GetUserByID(userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get user by ID %s: %w", userID, err)
	}
	log.Printf("user retrieved successfully with ID: %s", user.ID)
	return user, nil
}
