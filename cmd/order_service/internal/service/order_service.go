package services

import (
	"fmt"
	"time"

	"github.com/ajaysinghpanwar2002/pratilipi/cmd/order_service/internal/models"
	repositories "github.com/ajaysinghpanwar2002/pratilipi/cmd/order_service/internal/repository"
)

type OrderService struct {
	orderRepo *repositories.OrderRepository
}

func NewOrderService(repo *repositories.OrderRepository) *OrderService {
	return &OrderService{repo}
}

func (s *OrderService) PlaceOrder(order *models.Order) error {
	order.Status = "Placed"
	if err := s.orderRepo.CreateOrder(order); err != nil {
		return fmt.Errorf("failed to place order: %w", err)
	}
	return nil
}

func (s *OrderService) CreateOrder(userID, productID string, quantity int, productPrice float64) (*models.Order, error) {
	order := &models.Order{
		UserID:     userID,
		ProductID:  productID,
		Quantity:   quantity,
		Status:     "Pending",
		TotalPrice: productPrice * float64(quantity),
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	if err := s.PlaceOrder(order); err != nil {
		return nil, fmt.Errorf("failed to create order: %w", err)
	}

	return order, nil
}

func (s *OrderService) GetAllOrders() ([]*models.Order, error) {
	orders, err := s.orderRepo.GetAllOrders()
	if err != nil {
		return nil, fmt.Errorf("failed to get all orders: %w", err)
	}
	return orders, nil
}
