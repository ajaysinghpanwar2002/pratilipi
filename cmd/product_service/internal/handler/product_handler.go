package handlers

import (
	"log"

	"github.com/ajaysinghpanwar2002/pratilipi/cmd/product_service/internal/models"
	services "github.com/ajaysinghpanwar2002/pratilipi/cmd/product_service/internal/service"
	"github.com/ajaysinghpanwar2002/pratilipi/pkg/rabbitmq"
	"github.com/gofiber/fiber/v2"
)

const (
	productEventsQueue = "product_events"
)

type ProductHandler struct {
	service *services.ProductService
}

func NewProductHandler(service *services.ProductService) *ProductHandler {
	return &ProductHandler{service: service}
}

func (h *ProductHandler) CreateProduct(c *fiber.Ctx) error {
	ctx := c.Context()
	var product models.Product
	if err := c.BodyParser(&product); err != nil {
		return errorResponse(c, fiber.StatusBadRequest, "Invalid input")
	}
	if err := h.service.CreateProduct(ctx, &product); err != nil {
		return errorResponse(c, fiber.StatusInternalServerError, "Could not create product")
	}

	if err := emitProductEvent("ProductCreated", product); err != nil {
		log.Printf("Failed to emit event: %v", err)
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Product created successfully", "product_id": product.ID})
}

func (h *ProductHandler) GetProduct(c *fiber.Ctx) error {
	ctx := c.Context()
	id := c.Params("id")
	if id == "" {
		return errorResponse(c, fiber.StatusBadRequest, "Invalid product ID")
	}
	product, err := h.service.GetProductByID(ctx, id)
	if err != nil {
		return errorResponse(c, fiber.StatusNotFound, "Product not found")
	}
	return c.JSON(product)
}

func (h *ProductHandler) UpdateProduct(c *fiber.Ctx) error {
	ctx := c.Context()
	id := c.Params("id")
	if id == "" {
		return errorResponse(c, fiber.StatusBadRequest, "Invalid product ID")
	}
	var product models.Product
	if err := c.BodyParser(&product); err != nil {
		return errorResponse(c, fiber.StatusBadRequest, "Invalid input")
	}

	product.ID = id

	if err := h.service.UpdateProduct(ctx, &product); err != nil {
		return errorResponse(c, fiber.StatusInternalServerError, "Could not update product")
	}

	if err := emitProductEvent("InventoryUpdated", product); err != nil {
		log.Printf("Failed to emit event: %v", err)
	}

	return c.Status(fiber.StatusNoContent).JSON(fiber.Map{"message": "Product updated successfully"})
}

func (h *ProductHandler) DeleteProduct(c *fiber.Ctx) error {
	ctx := c.Context()
	id := c.Params("id")
	if id == "" {
		return errorResponse(c, fiber.StatusBadRequest, "Invalid product ID")
	}
	if err := h.service.DeleteProduct(ctx, id); err != nil {
		return errorResponse(c, fiber.StatusInternalServerError, "Could not delete product")
	}
	return c.SendStatus(fiber.StatusNoContent)
}

func (h *ProductHandler) GetAllProducts(c *fiber.Ctx) error {
	ctx := c.Context()
	products, err := h.service.GetAllProducts(ctx)
	if err != nil {
		return errorResponse(c, fiber.StatusInternalServerError, "Failed to get products")
	}
	return c.JSON(products)
}

func emitProductEvent(eventType string, product models.Product) error {
	return rabbitmq.EmitEvent(productEventsQueue, eventType, map[string]interface{}{
		"product_id": product.ID,
		"name":       product.Name,
		"price":      product.Price,
		"stock":      product.Stock,
	})
}

func errorResponse(c *fiber.Ctx, statusCode int, message string) error {
	return c.Status(statusCode).JSON(fiber.Map{"error": message})
}
