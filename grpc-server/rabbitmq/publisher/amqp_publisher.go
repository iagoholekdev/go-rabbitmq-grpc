package publisher

import (
	"log"

	"github.com/streadway/amqp"
)

func Publish(channel *amqp.Channel) {
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
