package repositories

import (
	"fmt"

	"github.com/ajaysinghpanwar2002/pratilipi/cmd/order_service/internal/models"
	"github.com/ajaysinghpanwar2002/pratilipi/pkg/db"
)

type UserRepository struct {
}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (r *UserRepository) CreateUser(user models.User) error {
	query := `INSERT INTO users (id, username, email, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)`
	result, err := db.DB.Exec(query, user.ID, user.Username, user.Email, user.CreatedAt, user.UpdatedAt)
	if err != nil {
		return fmt.Errorf("failed to insert user: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to retrieve rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no rows affected, user not inserted")
	}

	return nil
}

func (r *UserRepository) UpdateUser(userID string, updates map[string]interface{}) error {
	query := `UPDATE users SET username = COALESCE($2, username), email = COALESCE($3, email), updated_at = NOW() WHERE id = $1`
	_, err := db.DB.Exec(query, userID, updates["username"], updates["email"])
	if err != nil {
		return fmt.Errorf("failed to update user with ID %s: %w", userID, err)
	}
	return nil
}

func (r *UserRepository) GetUserByID(userID string) (*models.User, error) {
	var user models.User
	query := `SELECT * FROM users WHERE id = $1`
	err := db.DB.Get(&user, query, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch user by ID %s: %w", userID, err)
	}
	return &user, nil
}
