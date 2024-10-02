package repositories

import (
	"context"
	"fmt"
	"log"

	"github.com/ajaysinghpanwar2002/pratilipi/cmd/product_service/internal/models"
	"github.com/ajaysinghpanwar2002/pratilipi/pkg/db"
)

type ProductRepository struct{}

func NewProductRepository() *ProductRepository {
	return &ProductRepository{}
}

func (r *ProductRepository) CreateProduct(ctx context.Context, product *models.Product) error {
	query := `INSERT INTO products (name, description, price, stock) 
              VALUES (:name, :description, :price, :stock) 
              RETURNING id, created_at, updated_at`

	rows, err := db.DB.NamedQueryContext(ctx, query, product)
	if err != nil {
		log.Printf("Failed to create product: %v", err)
		return fmt.Errorf("failed to create product: %w", err)
	}
	defer rows.Close()

	if rows.Next() {
		if err := rows.StructScan(product); err != nil {
			log.Printf("Failed to scan product: %v", err)
			return fmt.Errorf("failed to scan product: %w", err)
		}
	}

	return nil
}

func (r *ProductRepository) GetProductByID(ctx context.Context, id int64) (*models.Product, error) {
	var product models.Product
	query := `SELECT * FROM products WHERE id = $1`
	if err := db.DB.GetContext(ctx, &product, query, id); err != nil {
		log.Printf("Failed to get product by ID: %v", err)
		return nil, fmt.Errorf("failed to get product by ID: %w", err)
	}
	return &product, nil
}

func (r *ProductRepository) UpdateProduct(ctx context.Context, product *models.Product) error {
	query := `UPDATE products SET name = :name, description = :description, price = :price, stock = :stock, updated_at = NOW() WHERE id = :id`
	_, err := db.DB.NamedExecContext(ctx, query, product)
	if err != nil {
		log.Printf("Failed to update product: %v", err)
		return fmt.Errorf("failed to update product: %w", err)
	}
	return nil
}

func (r *ProductRepository) DeleteProduct(ctx context.Context, id int64) error {
	query := `DELETE FROM products WHERE id = $1`
	_, err := db.DB.ExecContext(ctx, query, id)
	if err != nil {
		log.Printf("Failed to delete product: %v", err)
		return fmt.Errorf("failed to delete product: %w", err)
	}
	return nil
}
