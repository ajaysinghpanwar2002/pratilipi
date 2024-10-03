package models

import "time"

type Order struct {
	ID         string    `db:"id"`
	UserID     string    `db:"user_id"`
	ProductID  string    `db:"product_id"`
	Quantity   int       `db:"quantity"`
	Status     string    `db:"status"`
	TotalPrice float64   `db:"total_price"`
	CreatedAt  time.Time `db:"created_at"`
	UpdatedAt  time.Time `db:"updated_at"`
}
