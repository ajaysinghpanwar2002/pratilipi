package repositories

import (
	"github.com/ajaysinghpanwar2002/pratilipi/cmd/order_service/internal/models"
	"github.com/ajaysinghpanwar2002/pratilipi/pkg/db"
)

type ProductRepository struct {
}

func NewProductRepository() *ProductRepository {
	return &ProductRepository{}
}

func (r *ProductRepository) CreateProduct(product models.Product) error {
	query := `INSERT INTO products (id, name, price, stock, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := db.DB.Exec(query, product.ID, product.Name, product.Price, product.Stock, product.CreatedAt, product.UpdatedAt)
	return err
}

func (r *ProductRepository) UpdateStock(productID string, stock int) error {
	query := `UPDATE products SET stock = $2, updated_at = NOW() WHERE id = $1`
	_, err := db.DB.Exec(query, productID, stock)
	return err
}

func (r *ProductRepository) GetProductByID(productID string) (*models.Product, error) {
	var product models.Product
	query := `SELECT * FROM products WHERE id = $1`
	err := db.DB.Get(&product, query, productID)
	return &product, err
}
