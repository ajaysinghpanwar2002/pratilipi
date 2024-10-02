package main

import (
	"context"
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
)

const (
	port          = ":8080"
	queueName     = "product_events"
	migrationPath = "file://./db/migrations"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Handle graceful shutdown
	go handleShutdown(cancel)

	initializeDatabase(ctx)
	defer db.DB.Close()

	initializeRabbitMQ()
	defer rabbitmq.CloseRabbitMQ()

	app := fiber.New()
	setupRoutes(app)

	log.Printf("Service started successfully on port %s", port)
	if err := app.Listen(port); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}

func initializeDatabase(ctx context.Context) {
	db.Connect(ctx, "PRODUCT_DB")
	db.RunMigrations(ctx, db.DB.DB, migrationPath)
}

func initializeRabbitMQ() {
	if err := rabbitmq.ConnectRabbitMQ(); err != nil {
		log.Fatalf("RabbitMQ connection failed: %v", err)
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
		log.Fatalf("Failed to declare queue: %v", err)
	}
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
