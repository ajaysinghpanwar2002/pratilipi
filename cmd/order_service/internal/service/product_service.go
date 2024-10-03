package services

import (
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
	productID := data["product_id"].(string)
	name := data["name"].(string)
	price := data["price"].(float64)
	stock := data["stock"].(int)

	product := models.Product{
		ID:    productID,
		Name:  name,
		Price: price,
		Stock: stock,
	}

	return s.productRepo.CreateProduct(product)
}

func (s *ProductService) HandleInventoryUpdatedEvent(data map[string]interface{}) error {
	productID := data["product_id"].(string)
	stock := data["stock"].(int)

	return s.productRepo.UpdateStock(productID, stock)
}

func (s *ProductService) GetProductByID(productID string) (*models.Product, error) {
	product, err := s.productRepo.GetProductByID(productID)
	if err != nil {
		return nil, err
	}
	return product, nil
}
