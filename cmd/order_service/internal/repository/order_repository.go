package repositories

import (
	"fmt"

	"github.com/ajaysinghpanwar2002/pratilipi/cmd/order_service/internal/models"
	"github.com/ajaysinghpanwar2002/pratilipi/pkg/db"
)

type OrderRepository struct {
}

func NewOrderRepository() *OrderRepository {
	return &OrderRepository{}
}

func (r *OrderRepository) CreateOrder(order *models.Order) error {
	query := `INSERT INTO orders (user_id, product_id, quantity, status, total_price, created_at, updated_at) 
              VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`
	err := db.DB.QueryRow(query, order.UserID, order.ProductID, order.Quantity, order.Status, order.TotalPrice, order.CreatedAt, order.UpdatedAt).Scan(&order.ID)
	if err != nil {
		return fmt.Errorf("error inserting order: %w", err)
	}
	return nil
}

func (r *OrderRepository) GetOrderByID(orderID string) (*models.Order, error) {
	var order models.Order
	query := `SELECT * FROM orders WHERE id = $1`
	err := db.DB.Get(&order, query, orderID)
	if err != nil {
		return nil, fmt.Errorf("error fetching order by ID: %w", err)
	}
	return &order, nil
}

func (r *OrderRepository) GetAllOrders() ([]*models.Order, error) {
	var orders []*models.Order
	query := `SELECT * FROM orders`
	err := db.DB.Select(&orders, query)
	if err != nil {
		return nil, fmt.Errorf("error fetching all orders: %w", err)
	}
	return orders, nil
}
