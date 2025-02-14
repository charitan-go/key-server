package auth

import (
	"fmt"
	"log"
	"os"

	amqp "github.com/rabbitmq/amqp091-go"
)

type AuthRabbitMQProducer interface {
	NotiGetPrivateKey() error
}

type authRabbitMQProducerImpl struct{}

func NewAuthRabbitMQProducer() AuthRabbitMQProducer {
	return &authRabbitMQProducerImpl{}
}

func (p *authRabbitMQProducerImpl) NotiGetPrivateKey() error {
	// Connect
	amqpConnectionStr := fmt.Sprintf("amqp://%s:%s@message-broker:5672",
		os.Getenv("MESSAGE_BROKER_USER"),
		os.Getenv("MESSAGE_BROKER_PASSWORD"))
	conn, err := amqp.Dial(amqpConnectionStr)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
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
	exchangeName := "AUTH_NOTI_GET_PRIVATE_KEY"
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
	routingKey := "key.generated"
	body := "Key generation completed"
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
	}

	log.Printf("Published message: %s", body)
	return nil

}
