package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/ajaysinghpanwar2002/pratilipi/cmd/order_service/internal/handlers"
	repositories "github.com/ajaysinghpanwar2002/pratilipi/cmd/order_service/internal/repository"
	services "github.com/ajaysinghpanwar2002/pratilipi/cmd/order_service/internal/service"
	"github.com/ajaysinghpanwar2002/pratilipi/pkg/db"
	"github.com/ajaysinghpanwar2002/pratilipi/pkg/rabbitmq"
	"github.com/gofiber/fiber/v2"
	"github.com/streadway/amqp"
)

const (
	port          = ":8080"
	queueName     = "order_events"
	migrationPath = "file://./db/migrations"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go handleShutdown(cancel)

	if err := initializeDatabase(ctx); err != nil {
		log.Fatalf("Database initialization failed: %v", err)
	}
	defer db.DB.Close()

	if err := initializeRabbitMQ(); err != nil {
		log.Fatalf("RabbitMQ initialization failed: %v", err)
	}
	defer rabbitmq.CloseRabbitMQ()

	app := fiber.New()
	setupRoutes(app)

	consumeRabbitmqUserEvents(services.NewUserService(repositories.NewUserRepository()))
	consumeRabbitmqProductEvents(services.NewProductService(repositories.NewProductRepository()))

	log.Printf("Service started successfully on port %s", port)
	if err := app.Listen(port); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}

func initializeDatabase(ctx context.Context) error {
	if err := db.Connect(ctx, "ORDER_DB"); err != nil {
		return fmt.Errorf("database connection failed: %w", err)
	}
	return db.RunMigrations(ctx, db.DB.DB, migrationPath)
}

func initializeRabbitMQ() error {
	if err := rabbitmq.ConnectRabbitMQ(); err != nil {
		return fmt.Errorf("RabbitMQ connection failed: %w", err)
	}

	_, err := rabbitmq.Ch.QueueDeclare(
		queueName, // name
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)

	return err
}

func setupRoutes(app *fiber.App) {
	orderRepo := repositories.NewOrderRepository()
	userRepo := repositories.NewUserRepository()
	productRepo := repositories.NewProductRepository()

	orderService := services.NewOrderService(orderRepo)
	userService := services.NewUserService(userRepo)
	productService := services.NewProductService(productRepo)

	handler := handlers.NewHandler(orderService, userService, productService)

	app.Post("/orders", handler.PlaceOrder)
	app.Get("/orders", handler.GetOrders)

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("Order Service is running")
	})
}

func handleShutdown(cancel context.CancelFunc) {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	<-signalChan
	cancel()
	log.Println("Shutting down gracefully...")
}

func consumeRabbitmqUserEvents(userService *services.UserService) {
	queueName := "user_events"
	_, err := rabbitmq.Ch.QueueDeclare(
		queueName,
		true,  // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		log.Fatalf("Failed to declare queue %s: %v", queueName, err)
	}

	rabbitmq.ConsumeMessages(queueName, func(d amqp.Delivery) {
		event := parseEvent(d.Body)
		switch event.Type {
		case "UserRegistered":
			userService.HandleUserRegisteredEvent(event.Data)
		case "UserProfileUpdated":
			userService.HandleUserProfileUpdatedEvent(event.Data)
		default:
			log.Printf("Unhandled event type: %s", event.Type)
		}
	})
}

func consumeRabbitmqProductEvents(productService *services.ProductService) {
	queueName := "product_events"
	_, err := rabbitmq.Ch.QueueDeclare(
		queueName,
		true,  // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		log.Fatalf("Failed to declare queue %s: %v", queueName, err)
	}

	rabbitmq.ConsumeMessages(queueName, func(d amqp.Delivery) {
		event := parseEvent(d.Body)
		switch event.Type {
		case "ProductCreated":
			productService.HandleProductCreatedEvent(event.Data)
		case "InventoryUpdated":
			productService.HandleInventoryUpdatedEvent(event.Data)
		default:
			log.Printf("Unhandled event type: %s", event.Type)
		}
	})
}

func parseEvent(body []byte) (event struct {
	Type string
	Data map[string]interface{}
}) {
	if err := json.Unmarshal(body, &event); err != nil {
		log.Printf("Failed to parse event: %v", err)
	}
	return event
}
