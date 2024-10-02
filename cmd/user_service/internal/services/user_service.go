package services

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/ajaysinghpanwar2002/pratilipi/cmd/user_service/internal/models"
	"github.com/ajaysinghpanwar2002/pratilipi/cmd/user_service/internal/repositories"
	"golang.org/x/crypto/bcrypt"
)

const (
	errUserNotFound        = "user not found"
	errIncorrectPassword   = "username or password is incorrect"
	errHashingPassword     = "error hashing password"
	errCreatingUser        = "error creating user"
	errUpdatingUserProfile = "error updating user profile"
)

type UserService struct {
	repo *repositories.UserRepository
}

func NewUserService(repo *repositories.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) RegisterUser(ctx context.Context, user *models.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("%s: %v", errHashingPassword, err)
		return fmt.Errorf("%s: %w", errHashingPassword, err)
	}
	user.Password = string(hashedPassword)
	if err := s.repo.CreateUser(ctx, user); err != nil {
		log.Printf("%s: %v", errCreatingUser, err)
		return fmt.Errorf("%s: %w", errCreatingUser, err)
	}
	log.Printf("User registered successfully with ID: %d", user.ID)
	return nil
}

func (s *UserService) Authenticate(ctx context.Context, username, password string) (models.User, error) {
	user, err := s.repo.GetUserByUsername(ctx, username)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Printf("%s: %s", errUserNotFound, username)
			return user, errors.New(errUserNotFound)
		}
		log.Printf("Error retrieving user by username: %v", err)
		return user, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		log.Printf("%s: %s", errIncorrectPassword, username)
		return user, errors.New(errIncorrectPassword)
	}

	log.Printf("User authenticated successfully: %s", username)
	return user, nil
}

func (s *UserService) UpdateProfile(ctx context.Context, userId uint, updateData map[string]interface{}) error {
	if err := s.repo.UpdateUserProfile(ctx, userId, updateData); err != nil {
		log.Printf("%s: %v", errUpdatingUserProfile, err)
		return fmt.Errorf("%s: %w", errUpdatingUserProfile, err)
	}
	log.Printf("User profile updated successfully for user ID: %d", userId)
	return nil
}
