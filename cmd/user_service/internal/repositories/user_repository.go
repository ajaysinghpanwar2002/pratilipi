package repositories

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/ajaysinghpanwar2002/pratilipi/cmd/user_service/internal/models"
	"github.com/ajaysinghpanwar2002/pratilipi/pkg/db"
)

const (
	insertUserQuery = `INSERT INTO users (username, password, email, created_at, updated_at) 
                       VALUES ($1, $2, $3, $4, $5) RETURNING id`
	selectUserQuery = `SELECT * FROM users WHERE username = $1`
)

type UserRepository struct{}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (r *UserRepository) CreateUser(ctx context.Context, user *models.User) error {
	now := currentTime()
	err := db.DB.QueryRowContext(ctx, insertUserQuery, user.Username, user.Password, user.Email, now, now).Scan(&user.ID)
	if err != nil {
		log.Printf("Failed to register user: %v", err)
		return fmt.Errorf("failed to register user: %w", err)
	}
	return nil
}

func (r *UserRepository) GetUserByUsername(ctx context.Context, username string) (models.User, error) {
	var user models.User
	err := db.DB.GetContext(ctx, &user, selectUserQuery, username)
	if err != nil {
		log.Printf("Error retrieving user by username: %v", err)
		return models.User{}, fmt.Errorf("error retrieving user by username: %w", err)
	}
	return user, nil
}

func (r *UserRepository) UpdateUserProfile(ctx context.Context, userID string, updateData map[string]interface{}) error {
	query, values := buildUpdateQuery(updateData, userID)
	_, err := db.DB.ExecContext(ctx, query, values...)
	if err != nil {
		log.Printf("Failed to update user profile: %v", err)
		return fmt.Errorf("failed to update user profile: %w", err)
	}
	return nil
}

func buildUpdateQuery(updateData map[string]interface{}, userID string) (string, []interface{}) {
	query := "UPDATE users SET "
	values := []interface{}{}
	i := 1

	for key, value := range updateData {
		if i > 1 {
			query += ", "
		}
		query += key + " = $" + fmt.Sprint(i)
		values = append(values, value)
		i++
	}

	query += ", updated_at = $" + fmt.Sprint(i) + " WHERE id = $" + fmt.Sprint(i+1)
	values = append(values, currentTime(), userID)

	return query, values
}
func currentTime() time.Time {
	return time.Now()
}
