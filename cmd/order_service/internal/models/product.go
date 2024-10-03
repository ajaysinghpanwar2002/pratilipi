package models

import "time"

type Product struct {
	ID        string    `db:"id"`
	Name      string    `db:"name"`
	Price     float64   `db:"price"`
	Stock     int       `db:"stock"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
