package auth

import (
	"log"

	"github.com/charitan-go/key-server/rabbitmq"
	amqp "github.com/rabbitmq/amqp091-go"
)

type AuthRabbitmqProducer interface {
	NotiGetPrivateKey() error
}

type authRabbitmqProducerImpl struct {
	rabbitmqSvc rabbitmq.RabbitmqService
}

func NewAuthRabbitmqProducer(rabbitmqSvc rabbitmq.RabbitmqService) AuthRabbitmqProducer {
	return &authRabbitmqProducerImpl{rabbitmqSvc}
}

func (p *authRabbitmqProducerImpl) NotiGetPrivateKey() error {
	ch, err := p.rabbitmqSvc.ConnectRabbitmq()
	if err != nil {
		log.Fatalf("Failed to open channel")
	}

	// Declare a topic exchange (will create it if it doesn't exist)
	exchangeName := "GET_PRIVATE_KEY"
	err = p.rabbitmqSvc.DeclareExchange(ch, exchangeName)
	if err != nil {
		log.Fatalf("Failed to open channel: %v", err)
		return err
	}

	// Publish a message indicating the key generation is complete
	routingKey := "key.get.private.key"
	body := "Private key generated successfully"
	msg := amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(body),
	}
	err = p.rabbitmqSvc.Publish(ch, exchangeName, routingKey, msg)
	if err != nil {
		log.Fatalf("Failed to open channel: %v", err)
	} else {
		log.Printf("Published message: %s", body)
	}

	return nil
}
