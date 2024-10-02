package handlers

import (
	"github.com/ajaysinghpanwar2002/pratilipi/cmd/product_service/internal/models"
	services "github.com/ajaysinghpanwar2002/pratilipi/cmd/product_service/internal/service"
	"github.com/gofiber/fiber/v2"
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
	return c.Status(fiber.StatusCreated).JSON(product)
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

func errorResponse(c *fiber.Ctx, statusCode int, message string) error {
	return c.Status(statusCode).JSON(fiber.Map{"error": message})
}
