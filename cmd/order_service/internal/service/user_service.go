package services

import (
	"context"
	"log"

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
	// Extract the data
	userID := data["user_id"].(string)
	username := data["username"].(string)
	email := data["email"].(string)

	// Add user to the local users table
	user := models.User{
		ID:       userID,
		Username: username,
		Email:    email,
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
