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
		// Handle case where the user_id might be a float64 and convert it to a string
		userID = fmt.Sprintf("%.0f", data["user_id"].(float64))
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

	return s.userRepo.CreateUser(user)
}

func (s *UserService) HandleUserProfileUpdatedEvent(data map[string]interface{}) error {
	userID := data["id"].(string)
	// Assuming the event sends updated fields (e.g., username, email)
	updates := map[string]interface{}{}

	if username, ok := data["username"]; ok {
		updates["username"] = username
	}
	if email, ok := data["email"]; ok {
		updates["email"] = email
	}

	return s.userRepo.UpdateUser(userID, updates)
}

func (s *UserService) GetUserByID(ctx context.Context, userID string) (*models.User, error) {
	user, err := s.userRepo.GetUserByID(userID)
	if err != nil {
		return nil, err
	}
	log.Printf("user retrieved successfully with ID: %s", user.ID)
	return user, nil
}
