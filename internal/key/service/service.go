package service

import (
	"log"

	"github.com/charitan-go/key-server/external/auth"
	"github.com/charitan-go/key-server/pkg/proto"
	amqp "github.com/rabbitmq/amqp091-go"
)

type KeyService interface {
	NotiGetPrivateKey()
	GetPrivateKey(*proto.GetPrivateKeyRequestDto) (*proto.GetPrivateKeyResponseDto, error)
}

type keyServiceImpl struct {
	authRabbitMQProducer auth.AuthRabbitMQProducer
}

func NewKeyService(authRabbitMQProducer auth.AuthRabbitMQProducer) KeyService {
	return &keyServiceImpl{authRabbitMQProducer}
}

func (svc *keyServiceImpl) GetPrivateKey(reqDto *proto.GetPrivateKeyRequestDto) (*proto.GetPrivateKeyResponseDto, error) {
	return nil, nil
}

func (svc *keyServiceImpl) NotiGetPrivateKey() {
	// Connect to RabbitMQ
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ %v", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel %v", err)
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
		log.Fatalf("Failed to declare an exchange %v", err)
	}

	// Publish a message indicating the key generation is complete
	routingKey := "key.generated"
	body := "Key generation completed: key version 1"
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
		log.Fatalf("Failed to declare an exchange %v", err)
	}

	log.Printf("Published message: %s", body)

}
