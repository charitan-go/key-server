package auth

import (
	"fmt"
	"log"
	"os"

	amqp "github.com/rabbitmq/amqp091-go"
)

type AuthRabbitmqProducer interface {
	NotiGetPrivateKey() error
}

type authRabbitmqProducerImpl struct{}

func NewAuthRabbitmqProducer() AuthRabbitmqProducer {
	return &authRabbitmqProducerImpl{}
}

func (p *authRabbitmqProducerImpl) NotiGetPrivateKey() error {
	// Connect
	amqpConnectionStr := fmt.Sprintf("amqp://%s:%s@message-broker:5672",
		os.Getenv("MESSAGE_BROKER_USER"),
		os.Getenv("MESSAGE_BROKER_PASSWORD"))
	conn, err := amqp.Dial(amqpConnectionStr)
	if err != nil {
		log.Fatalf("Failed to connect to Rabbitmq: %v", err)
		return err
	}
	defer conn.Close()

	// Open channel
	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open channel: %v", err)
		return err
	}
	defer ch.Close()

	// Declare a topic exchange (will create it if it doesn't exist)
	exchangeName := "GET_PRIVATE_KEY"
	err = ch.ExchangeDeclare(
		exchangeName, // name
		"topic",      // type
		true,         // durable
		false,        // auto-deleted
		false,        // internal
		false,        // no-wait
		nil,          // arguments
	)
	if err != nil {
		log.Fatalf("Failed to open channel: %v", err)
	}

	// Publish a message indicating the key generation is complete
	routingKey := "key.get.private.key"
	body := "Private key generated successfully"
	err = ch.Publish(
		exchangeName, // exchange
		routingKey,   // routing key
		false,        // mandatory
		false,        // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	if err != nil {
		log.Fatalf("Failed to open channel: %v", err)
	} else {
		log.Printf("Published message: %s", body)
	}

	return nil

}
