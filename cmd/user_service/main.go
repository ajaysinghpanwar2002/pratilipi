package main

import (
	"log"

	"github.com/ajaysinghpanwar2002/pratilipi/cmd/user_service/db"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// Initialize Fiber app
	app := fiber.New()

	// Connect to the database
	db.Connect()

	// Run database migrations
	db.RunMigrations(db.DB.DB)

	// Define a simple GET endpoint
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("User Service is running!")
	})

	// Start the server on port 8080
	log.Println("Service started successfully on port 8080")
	if err := app.Listen(":8080"); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
