package main

import (
	"log"

	"github.com/ajaysinghpanwar2002/pratilipi/cmd/user_service/db"
	"github.com/ajaysinghpanwar2002/pratilipi/cmd/user_service/internal/handlers"
	"github.com/ajaysinghpanwar2002/pratilipi/cmd/user_service/internal/middlewares"
	"github.com/ajaysinghpanwar2002/pratilipi/cmd/user_service/internal/repositories"
	"github.com/ajaysinghpanwar2002/pratilipi/cmd/user_service/internal/services"
	"github.com/ajaysinghpanwar2002/pratilipi/pkg/rabbitmq"
	"github.com/gofiber/fiber/v2"
)

const (
	port      = ":8080"
	queueName = "user_events"
)

func main() {
	initializeDatabase()
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

func initializeDatabase() {
	db.Connect()
	db.RunMigrations(db.DB.DB)
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
	userRepo := &repositories.UserRepository{}
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	app.Post("/register", userHandler.RegisterUser)
	app.Post("/login", userHandler.LoginUser)
	app.Put("/profile", middlewares.AuthMiddleware, userHandler.UpdateProfile)

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("User Service is running")
	})
}
