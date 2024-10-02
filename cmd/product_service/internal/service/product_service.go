package services

import (
	"context"
	"fmt"
	"log"

	"github.com/ajaysinghpanwar2002/pratilipi/cmd/product_service/internal/models"
	"github.com/ajaysinghpanwar2002/pratilipi/cmd/product_service/internal/repositories"
)

type ProductService struct {
	repo *repositories.ProductRepository
}

func NewProductService(repo *repositories.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) CreateProduct(ctx context.Context, product *models.Product) error {
	if err := s.repo.CreateProduct(ctx, product); err != nil {
		return s.logAndReturnError("Failed to create product", err)
	}
	log.Printf("Product created successfully with ID: %d", product.ID)
	return nil
}

func (s *ProductService) GetProductByID(ctx context.Context, id int64) (*models.Product, error) {
	product, err := s.repo.GetProductByID(ctx, id)
	if err != nil {
		return nil, s.logAndReturnError("Failed to get product by ID", err)
	}
	log.Printf("Product retrieved successfully with ID: %d", product.ID)
	return product, nil
}

func (s *ProductService) UpdateProduct(ctx context.Context, product *models.Product) error {
	if err := s.repo.UpdateProduct(ctx, product); err != nil {
		return s.logAndReturnError("Failed to update product", err)
	}
	log.Printf("Product updated successfully with ID: %d", product.ID)
	return nil
}

func (s *ProductService) DeleteProduct(ctx context.Context, id int64) error {
	if err := s.repo.DeleteProduct(ctx, id); err != nil {
		return s.logAndReturnError("Failed to delete product", err)
	}
	log.Printf("Product deleted successfully with ID: %d", id)
	return nil
}

func (s *ProductService) logAndReturnError(message string, err error) error {
	log.Printf("%s: %v", message, err)
	return fmt.Errorf("%s: %w", message, err)
}
