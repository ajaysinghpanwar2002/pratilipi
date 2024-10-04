package handlers

import (
	"fmt"
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

	user, err := h.userService.GetUserByID(c.Context(), orderInput.UserID)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "User not found")
	}

	product, err := h.productService.GetProductByID(orderInput.ProductID)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "Product not found")
	}

	if product.Stock < orderInput.Quantity {
		return fiber.NewError(fiber.StatusBadRequest, "Not enough stock")
	}

	order, err := h.orderService.CreateOrder(user.ID, product.ID, orderInput.Quantity, product.Price)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to create order")
	}

	// Update the product stock
	product.Stock -= orderInput.Quantity
	if err := h.productService.UpdateProductStock(product.ID, product.Stock); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to update product stock")
	}

	err = rabbitmq.EmitEvent("order_events", "OrderPlaced", map[string]interface{}{
		"order_id":   order.ID,
		"user_id":    user.ID,
		"product_id": product.ID,
		"quantity":   orderInput.Quantity,
		"status":     order.Status,
	})

	fmt.Println("OrderPlaced event emitted")

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

func (h *OrderHandler) GetOrder(c *fiber.Ctx) error {
	orderID := c.Params("id")
	order, err := h.orderService.GetOrderByID(orderID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Order not found",
		})
	}
	return c.JSON(order)
}
