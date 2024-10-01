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

func main() {
	db.Connect()
	db.RunMigrations(db.DB.DB)

	defer db.DB.Close()

	err := rabbitmq.ConnectRabbitMQ()
	if err != nil {
		log.Fatalf("RabbitMQ connection failed: %v", err)
	}
	defer rabbitmq.CloseRabbitMQ()

	// queues for the user events
	_, err = rabbitmq.Ch.QueueDeclare(
		"user_events", // name
		true,          // durable
		false,         // delete when unused
		false,         // exclusive
		false,         // no-wait
		nil,           // arguments
	)

	if err != nil {
		log.Fatalf("Failed to declare queue: %v", err)
	}

	userRepo := &repositories.UserRepository{}
	UserService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(UserService)

	app := fiber.New()

	// Routes
	app.Post("/register", userHandler.RegisterUser)
	app.Post("/login", userHandler.LoginUser)
	app.Put("/profile", middlewares.AuthMiddleware, userHandler.UpdateProfile)

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("User Service is running")
	})

	// Start server
	log.Println("Service started successfully on port 8080")
	if err := app.Listen(":8080"); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
