package services

import (
	"fmt"
	"time"

	"github.com/ajaysinghpanwar2002/pratilipi/cmd/order_service/internal/models"
	repositories "github.com/ajaysinghpanwar2002/pratilipi/cmd/order_service/internal/repository"
)

type ProductService struct {
	productRepo *repositories.ProductRepository
}

func NewProductService(productRepo *repositories.ProductRepository) *ProductService {
	return &ProductService{productRepo}
}

func (s *ProductService) HandleProductCreatedEvent(data map[string]interface{}) error {
	productID, ok := data["product_id"].(string)
	if !ok {
		return fmt.Errorf("invalid product_id in event data")
	}
	name, ok := data["name"].(string)
	if !ok {
		return fmt.Errorf("invalid name in event data")
	}
	price, ok := data["price"].(float64)
	if !ok {
		return fmt.Errorf("invalid price in event data")
	}
	stock, ok := data["stock"].(float64)
	if !ok {
		return fmt.Errorf("invalid stock in event data")
	}

	now := time.Now()

	product := models.Product{
		ID:        productID,
		Name:      name,
		Price:     price,
		Stock:     int(stock),
		CreatedAt: now,
		UpdatedAt: now,
	}

	if err := s.productRepo.CreateProduct(&product); err != nil {
		return fmt.Errorf("failed to create product: %w", err)
	}
	return nil
}

func (s *ProductService) HandleInventoryUpdatedEvent(data map[string]interface{}) error {
	productID, ok := data["product_id"].(string)
	if !ok {
		return fmt.Errorf("invalid product_id in event data")
	}
	stock, ok := data["stock"].(float64)
	if !ok {
		return fmt.Errorf("invalid stock in event data")
	}

	if err := s.productRepo.UpdateStock(productID, int(stock)); err != nil {
		return fmt.Errorf("failed to update stock for product ID %s: %w", productID, err)
	}
	return nil
}

func (s *ProductService) GetProductByID(productID string) (*models.Product, error) {
	product, err := s.productRepo.GetProductByID(productID)
	if err != nil {
		return nil, fmt.Errorf("failed to get product by ID %s: %w", productID, err)
	}
	return product, nil
}

func (s *ProductService) UpdateProductStock(productID string, stock int) error {
	if err := s.productRepo.UpdateStock(productID, stock); err != nil {
		return fmt.Errorf("failed to update stock for product ID %s: %w", productID, err)
	}
	return nil
}
