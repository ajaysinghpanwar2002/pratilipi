package models

import "time"

type Product struct {
	ID        string    `db:"id" json:"id"`
	Name      string    `db:"name" json:"name"`
	Price     float64   `db:"price" json:"price"`
	Stock     int       `db:"stock" json:"stock"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}
