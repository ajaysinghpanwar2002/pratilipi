package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/ajaysinghpanwar2002/pratilipi/cmd/user_service/db"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/streadway/amqp"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        uint      `json:"id" db:"id"`
	Username  string    `json:"username" db:"username"`
	Password  string    `json:"-" db:"password"`
	Email     string    `json:"email" db:"email"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

var (
	rabbitmqConn *amqp.Connection
	rabbitmqCh   *amqp.Channel
)

func main() {
	db.Connect()

	var err error
	rabbitmqConn, err = amqp.Dial(os.Getenv("RABBITMQ_URL"))
	if err != nil {
		log.Fatal("Failed to connect to RabbitMQ:", err)
	}
	defer rabbitmqConn.Close()

	rabbitmqCh, err = rabbitmqConn.Channel()
	if err != nil {
		log.Fatal("Failed to open a channel:", err)
	}
	defer rabbitmqCh.Close()

	// Declare queues
	_, err = rabbitmqCh.QueueDeclare(
		"user_events", // name
		true,          // durable
		false,         // delete when unused
		false,         // exclusive
		false,         // no-wait
		nil,           // arguments
	)
	if err != nil {
		log.Fatal("Failed to declare a queue:", err)
	}

	go consumeMessages()

	app := fiber.New()

	// Routes
	app.Post("/register", registerUser)
	app.Post("/login", loginUser)
	app.Put("/profile", authMiddleware, updateProfile)
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("User Service is running!")
	})

	// Start server
	log.Println("Service started successfully on port 8080")
	if err := app.Listen(":8080"); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}

func authMiddleware(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	tokenString := authHeader[len("Bearer "):]
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil || !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	// Set user ID in the context for further use
	c.Locals("user_id", claims["user_id"])
	return c.Next()
}

func registerUser(c *fiber.Ctx) error {
	user := new(User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to hash password"})
	}
	user.Password = string(hashedPassword)

	query := `INSERT INTO users (username, password, email, created_at, updated_at) 
              VALUES ($1, $2, $3, $4, $5) RETURNING id`
	err = db.DB.QueryRow(query, user.Username, user.Password, user.Email, time.Now(), time.Now()).Scan(&user.ID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to register user"})
	}

	// Emit "User Registered" event
	event := map[string]interface{}{
		"type": "UserRegistered",
		"data": map[string]interface{}{
			"id":       user.ID,
			"username": user.Username,
			"email":    user.Email,
		},
	}
	emitEvent(event)

	return c.Status(201).JSON(fiber.Map{"message": "User registered successfully", "user_id": user.ID})
}

func loginUser(c *fiber.Ctx) error {
	loginData := struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}{}

	if err := c.BodyParser(&loginData); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	var user User
	err := db.DB.Get(&user, "SELECT * FROM users WHERE username = $1", loginData.Username)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{"error": "Invalid credentials"})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginData.Password)); err != nil {
		return c.Status(401).JSON(fiber.Map{"error": "Invalid credentials"})
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to generate token"})
	}

	return c.JSON(fiber.Map{"token": tokenString})
}

func updateProfile(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(float64) // Extract user_id from JWT

	var updateData map[string]interface{}
	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	// Build the SQL query dynamically
	query := "UPDATE users SET "
	values := []interface{}{}
	i := 1

	for key, value := range updateData {
		if i > 1 {
			query += ", "
		}
		query += key + " = $" + fmt.Sprint(i)
		values = append(values, value)
		i++
	}

	query += ", updated_at = $" + fmt.Sprint(i) + " WHERE id = $" + fmt.Sprint(i+1)
	values = append(values, time.Now(), int(userID))

	_, err := db.DB.Exec(query, values...)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to update profile"})
	}

	// Emit "UserProfileUpdated" event
	event := map[string]interface{}{
		"type": "UserProfileUpdated",
		"data": map[string]interface{}{
			"id": userID,
		},
	}
	emitEvent(event)

	return c.JSON(fiber.Map{"message": "Profile updated successfully"})
}

func emitEvent(event map[string]interface{}) {
	eventJSON, err := json.Marshal(event)
	if err != nil {
		log.Println("Failed to marshal event:", err)
		return
	}

	err = rabbitmqCh.Publish(
		"",            // exchange
		"user_events", // routing key
		false,         // mandatory
		false,         // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        eventJSON,
		})
	if err != nil {
		log.Println("Failed to publish event:", err)
	}
}

func consumeMessages() {
	msgs, err := rabbitmqCh.Consume(
		"user_events", // queue
		"",            // consumer
		true,          // auto-ack
		false,         // exclusive
		false,         // no-local
		false,         // no-wait
		nil,           // args
	)
	if err != nil {
		log.Fatal("Failed to register a consumer:", err)
	}

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			var event map[string]interface{}
			if err := json.Unmarshal(d.Body, &event); err != nil {
				log.Println("Error parsing event:", err)
				continue
			}

			// Process the event
			switch event["type"] {
			case "UserRegistered":
				log.Println("User registered:", event["data"])
			case "UserProfileUpdated":
				log.Println("User profile updated:", event["data"])
			// Add more cases for other event types as needed
			default:
				log.Println("Unknown event type:", event["type"])
			}
		}
	}()

	log.Println("Waiting for messages. To exit press CTRL+C")
	<-forever
}
