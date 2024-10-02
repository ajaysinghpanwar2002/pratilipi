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

	err := rabbitmq.EmitEvent(productEventsQueue, "ProductCreated", map[string]interface{}{
		"product_id": product.ID,
		"name":       product.Name,
		"price":      product.Price,
		"Stock":      product.Stock,
	})

	if err != nil {
		log.Printf("Failed to emit event: %v", err)
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "product created successfully", "user_id": product})
}

func (h *ProductHandler) GetProduct(c *fiber.Ctx) error {
	ctx := c.Context()
	id, err := c.ParamsInt("id")
	if err != nil {
		return errorResponse(c, fiber.StatusBadRequest, "Invalid product ID")
	}
	product, err := h.service.GetProductByID(ctx, int64(id))
	if err != nil {
		return errorResponse(c, fiber.StatusNotFound, "Product not found")
	}
	return c.JSON(product)
}

func (h *ProductHandler) UpdateProduct(c *fiber.Ctx) error {
	ctx := c.Context()
	id, err := c.ParamsInt("id")
	if err != nil {
		return errorResponse(c, fiber.StatusBadRequest, "Invalid product ID")
	}
	var product models.Product
	if err := c.BodyParser(&product); err != nil {
		return errorResponse(c, fiber.StatusBadRequest, "Invalid input")
	}
	product.ID = int64(id)
	if err := h.service.UpdateProduct(ctx, &product); err != nil {
		return errorResponse(c, fiber.StatusInternalServerError, "Could not update product")
	}

	err1 := rabbitmq.EmitEvent(productEventsQueue, "InventoryUpdated", map[string]interface{}{
		"product_id": product.ID,
		"stock":      product.Stock,
	})

	if err1 != nil {
		log.Printf("Failed to emit event: %v", err)
	}

	return c.JSON(product)
}

func (h *ProductHandler) DeleteProduct(c *fiber.Ctx) error {
	ctx := c.Context()
	id, err := c.ParamsInt("id")
	if err != nil {
		return errorResponse(c, fiber.StatusBadRequest, "Invalid product ID")
	}
	if err := h.service.DeleteProduct(ctx, int64(id)); err != nil {
		return errorResponse(c, fiber.StatusInternalServerError, "Could not delete product")
	}
	return c.SendStatus(fiber.StatusNoContent)
}

func errorResponse(c *fiber.Ctx, statusCode int, message string) error {
	return c.Status(statusCode).JSON(fiber.Map{"error": message})
}
