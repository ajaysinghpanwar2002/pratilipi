package services

import (
	"github.com/ajaysinghpanwar2002/pratilipi/cmd/user_service/internal/models"
	"github.com/ajaysinghpanwar2002/pratilipi/cmd/user_service/internal/repositories"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo *repositories.UserRepository
}

func NewUserService(repo *repositories.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) RegisterUser(user *models.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return s.repo.CreateUser(user)
}

func (s *UserService) Authenticate(username, password string) (bool, models.User, error) {
	user, err := s.repo.GetUserByUsername(username)
	if err != nil {
		return false, user, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err == nil, user, err
}

func (s *UserService) UpdateProfile(userId uint, updateData map[string]interface{}) error {
	return s.repo.UpdateUserProfile(userId, updateData)
}
