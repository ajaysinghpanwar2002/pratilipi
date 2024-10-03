package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	handlers "github.com/ajaysinghpanwar2002/pratilipi/cmd/product_service/internal/handler"
	"github.com/ajaysinghpanwar2002/pratilipi/cmd/product_service/internal/repositories"
	services "github.com/ajaysinghpanwar2002/pratilipi/cmd/product_service/internal/service"
	"github.com/ajaysinghpanwar2002/pratilipi/pkg/db"
	"github.com/ajaysinghpanwar2002/pratilipi/pkg/rabbitmq"
	"github.com/gofiber/fiber/v2"
	"github.com/streadway/amqp"
)

const (
	port          = ":8080"
	queueName     = "product_events"
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

	consumeRabbitmqOrderEvents(ctx, services.NewProductService(repositories.NewProductRepository()))

	log.Printf("Service started successfully on port %s", port)
	if err := app.Listen(port); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}

func initializeDatabase(ctx context.Context) error {
	if err := db.Connect(ctx, "PRODUCT_DB"); err != nil {
		return fmt.Errorf("database connection failed: %w", err)
	}
	if err := db.RunMigrations(ctx, db.DB.DB, migrationPath); err != nil {
		return fmt.Errorf("database migration failed: %w", err)
	}
	return nil
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
	if err != nil {
		return fmt.Errorf("queue declaration failed: %w", err)
	}

	return nil
}

func setupRoutes(app *fiber.App) {
	productRepo := repositories.NewProductRepository()
	productService := services.NewProductService(productRepo)
	productHandler := handlers.NewProductHandler(productService)

	app.Post("/products", productHandler.CreateProduct)
	app.Get("/products/:id", productHandler.GetProduct)
	app.Put("/products/:id", productHandler.UpdateProduct)
	app.Delete("/products/:id", productHandler.DeleteProduct)

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("Product Service is running")
	})
}

func handleShutdown(cancel context.CancelFunc) {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	<-signalChan
	cancel()
	log.Println("Shutting down gracefully...")
}

func consumeRabbitmqOrderEvents(ctx context.Context, productService *services.ProductService) {
	queueName := "order_events"
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
		case "OrderPlaced":
			if err := productService.HandleOrderPlacedEvent(ctx, event.Data); err != nil {
				log.Printf("Failed to handle OrderPlaced event: %v", err)
			}
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
