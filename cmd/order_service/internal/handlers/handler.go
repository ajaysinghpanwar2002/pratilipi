package handlers

import (
	"log"

	"github.com/ajaysinghpanwar2002/pratilipi/cmd/order_service/internal/models"
	services "github.com/ajaysinghpanwar2002/pratilipi/cmd/order_service/internal/service"
	"github.com/ajaysinghpanwar2002/pratilipi/pkg/rabbitmq"
	"github.com/gofiber/fiber/v2"
)

type OrderHandler struct {
	orderService   *services.OrderService
	userService    *services.UserService
	productService *services.ProductService
}

func NewHandler(orderService *services.OrderService, userService *services.UserService, productService *services.ProductService) *OrderHandler {
	return &OrderHandler{orderService, userService, productService}
}

func (h *OrderHandler) PlaceOrder(c *fiber.Ctx) error {
	var orderInput models.Order
	if err := c.BodyParser(&orderInput); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid input")
	}

	// Step 1: Check if user exists
	user, err := h.userService.GetUserByID(c.Context(), orderInput.UserID)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "User not found")
	}

	// Step 2: Check if product exists and has enough stock
	product, err := h.productService.GetProductByID(orderInput.ProductID)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "Product not found")
	}

	if product.Stock < orderInput.Quantity {
		return fiber.NewError(fiber.StatusBadRequest, "Not enough stock")
	}

	// Step 3: Create the order
	order, err := h.orderService.CreateOrder(user.ID, product.ID, orderInput.Quantity)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to create order")
	}

	// Step 4: Emit OrderPlaced event
	err = rabbitmq.EmitEvent("order_events", "OrderPlaced", map[string]interface{}{
		"order_id":   order.ID,
		"user_id":    user.ID,
		"product_id": product.ID,
		"quantity":   orderInput.Quantity,
		"status":     order.Status,
	})
	if err != nil {
		log.Printf("Failed to emit order placed event: %v", err)
	}

	return c.Status(fiber.StatusCreated).JSON(order)
}

func (h *OrderHandler) GetOrders(c *fiber.Ctx) error {
	orders, err := h.orderService.GetAllOrders()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve orders",
		})
	}
	return c.JSON(orders)
}
