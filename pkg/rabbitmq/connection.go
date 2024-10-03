package rabbitmq

import (
	"fmt"
	"log"
	"os"

	"github.com/streadway/amqp"
)

var Conn *amqp.Connection
var Ch *amqp.Channel

// ConnectRabbitMQ establishes a connection to RabbitMQ
func ConnectRabbitMQ() error {
	var err error
	rabbitMQURL := os.Getenv("RABBITMQ_URL")
	if rabbitMQURL == "" {
		return fmt.Errorf("RABBITMQ_URL not set in environment")
	}

	// Connect to RabbitMQ
	Conn, err = amqp.Dial(rabbitMQURL)
	if err != nil {
		return fmt.Errorf("failed to connect to RabbitMQ: %w", err)
	}

	// Open a channel
	Ch, err = Conn.Channel()
	if err != nil {
		return fmt.Errorf("failed to open a channel: %w", err)
	}

	log.Println("Connected to RabbitMQ successfully!")
	return nil
}

// CloseRabbitMQ closes the RabbitMQ connection and channel
func CloseRabbitMQ() {
	if Ch != nil {
		Ch.Close()
	}
	if Conn != nil {
		Conn.Close()
	}
	log.Println("RabbitMQ connection closed.")
}
