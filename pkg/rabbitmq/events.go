package rabbitmq

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

func EmitEvent(queueName string, eventType string, data map[string]interface{}) error {
	if Ch == nil {
		return fmt.Errorf("RabbitMQ channel is not initialized")
	}

	event := map[string]interface{}{
		"type": eventType,
		"data": data,
	}

	eventJSON, err := json.Marshal(event)
	if err != nil {
		return fmt.Errorf("failed to marshal event: %w", err)
	}

	err = Ch.Publish(
		"",        // exchange
		queueName, // routing key
		false,     // mandatory
		false,     // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        eventJSON,
		},
	)
	if err != nil {
		return fmt.Errorf("failed to publish event: %w", err)
	}

	log.Printf("Event of type %s emitted to queue %s", eventType, queueName)
	return nil
}

// ConsumeMessages starts consuming messages from the specified queue
func ConsumeMessages(queueName string, handler func(amqp.Delivery)) error {
	if Ch == nil {
		return fmt.Errorf("RabbitMQ channel is not initialized")
	}

	msgs, err := Ch.Consume(
		queueName, // queue
		"",        // consumer
		true,      // auto-ack
		false,     // exclusive
		false,     // no-local
		false,     // no-wait
		nil,       // args
	)
	if err != nil {
		return fmt.Errorf("failed to register consumer: %w", err)
	}

	// Goroutine to handle messages
	go func() {
		for d := range msgs {
			handler(d)
		}
	}()

	log.Printf("Started consuming messages from queue %s", queueName)
	return nil
}
