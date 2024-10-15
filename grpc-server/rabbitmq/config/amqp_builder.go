package config

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

var connection *amqp.Connection
var channel *amqp.Channel

func DialRabbitMQ() (*amqp.Connection, error) {
	var err error
	connection, err = amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		return nil, fmt.Errorf("failed to connect to RabbitMQ: %w", err)
	}
	return connection, nil
}

func CreateChannel() (*amqp.Channel, error) {
	if connection == nil {
		return nil, fmt.Errorf("no RabbitMQ connection available")
	}

	var err error
	channel, err = connection.Channel()
	if err != nil {
		return nil, fmt.Errorf("failed to create channel: %w", err)
	}
	return channel, nil
}

func CreateQueue() {
	if channel == nil {
		fmt.Println("no RabbitMQ channel available")
	}

	queue, err := channel.QueueDeclare(
		"QueueIago",
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		panic(err)
	}

	fmt.Println(queue)
}

func Publish() {
	message := `{"hello": "world"}`
	err := channel.Publish(
		"",
		"QueueIago",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        []byte(message),
		},
	)
	if err != nil {
		log.Fatalf("Failed to publish message: %v", err)
	}
}

func Close() {
	if channel != nil {
		err := channel.Close()
		if err != nil {
			log.Printf("error closing channel: %v", err)
		}
	}

	if connection != nil {
		err := connection.Close()
		if err != nil {
			log.Printf("error closing connection: %v", err)
		}
	}
}
