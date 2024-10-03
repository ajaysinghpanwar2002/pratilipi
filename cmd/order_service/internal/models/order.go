package models

import "time"

type Order struct {
	ID         string    `db:"id" json:"id"`
	UserID     string    `db:"user_id" json:"user_id"`
	ProductID  string    `db:"product_id" json:"product_id"`
	Quantity   int       `db:"quantity" json:"quantity"`
	Status     string    `db:"status" json:"status"`
	TotalPrice float64   `db:"total_price" json:"total_price"`
	CreatedAt  time.Time `db:"created_at" json:"created_at"`
	UpdatedAt  time.Time `db:"updated_at" json:"updated_at"`
}
