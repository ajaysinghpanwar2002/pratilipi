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

	fmt.Println("Order Input: ", orderInput)

	// Step 1: Check if user exists
	user, err := h.userService.GetUserByID(c.Context(), orderInput.UserID)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "User not found")
	}

	fmt.Println("User that we found after search on our local table ", user)

	// Step 2: Check if product exists and has enough stock
	product, err := h.productService.GetProductByID(orderInput.ProductID)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "Product not found")
	}

	fmt.Println("Product that we found after search on our local table ", product)

	if product.Stock < orderInput.Quantity {
		return fiber.NewError(fiber.StatusBadRequest, "Not enough stock")
	}

	fmt.Println("Product has enough stock")

	// Step 3: Create the order
	order, err := h.orderService.CreateOrder(user.ID, product.ID, orderInput.Quantity, product.Price)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to create order")
	}

	fmt.Println("Order that we created after search on our local table ", order)

	// Step 4: Emit OrderPlaced event
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
