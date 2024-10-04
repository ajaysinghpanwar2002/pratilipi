package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/ajaysinghpanwar2002/pratilipi/cmd/user_service/internal/handlers"
	"github.com/ajaysinghpanwar2002/pratilipi/cmd/user_service/internal/middlewares"
	"github.com/ajaysinghpanwar2002/pratilipi/cmd/user_service/internal/repositories"
	"github.com/ajaysinghpanwar2002/pratilipi/cmd/user_service/internal/services"
	"github.com/ajaysinghpanwar2002/pratilipi/pkg/db"
	"github.com/ajaysinghpanwar2002/pratilipi/pkg/rabbitmq"
	"github.com/gofiber/fiber/v2"
)

const (
	port          = ":8080"
	queueName     = "user_events"
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
	db.Connect(ctx, "USER_DB")
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
	userRepo := repositories.NewUserRepository()
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	app.Post("/register", userHandler.RegisterUser)
	app.Post("/login", userHandler.LoginUser)
	app.Put("/profile", middlewares.AuthMiddleware, userHandler.UpdateProfile)
	app.Get("/users", userHandler.GetAllUsers)
	app.Get("/users/:id", userHandler.GetUserById)

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("User Service is running")
	})
}

func handleShutdown(cancel context.CancelFunc) {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	<-signalChan
	cancel()
	log.Println("Shutting down gracefully...")
}
