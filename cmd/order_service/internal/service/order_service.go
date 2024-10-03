package services

import (
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
	return s.orderRepo.CreateOrder(order)
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

	err := s.PlaceOrder(order)
	if err != nil {
		return nil, err
	}

	return order, nil
}

func (s *OrderService) GetAllOrders() ([]*models.Order, error) {
	return s.orderRepo.GetAllOrders()
}
