package repositories

import (
	"fmt"
	"log"
	"time"

	"github.com/ajaysinghpanwar2002/pratilipi/cmd/user_service/db"
	"github.com/ajaysinghpanwar2002/pratilipi/cmd/user_service/internal/models"
)

type UserRepository struct{}

func (r *UserRepository) CreateUser(user *models.User) error {
	query := `INSERT INTO users (username, password, email, created_at, updated_at) 
              VALUES ($1, $2, $3, $4, $5) RETURNING id`
	err := db.DB.QueryRow(query, user.Username, user.Password, user.Email, time.Now(), time.Now()).Scan(&user.ID)
	if err != nil {
		log.Printf("Failed to register user: %v", err)
		return err
	}
	return nil
}

func (r *UserRepository) GetUserByUsername(username string) (models.User, error) {
	var user models.User
	err := db.DB.Get(&user, "SELECT * FROM users WHERE username = $1", username)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (r *UserRepository) UpdateUserProfile(userID uint, updateData map[string]interface{}) error {
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
	values = append(values, updateData["updated_at"], int(userID))

	_, err := db.DB.Exec(query, values...)
	return err
}
