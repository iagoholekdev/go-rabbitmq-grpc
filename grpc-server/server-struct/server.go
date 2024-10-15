package serverstruct

import (
	pb "github.com/iagoholekdev/go-rabbitmq-grpc/grpc-server/publisher"
	"github.com/streadway/amqp"
)

type server struct {
	pb.UnimplementedPublisherServiceServer
	connection *amqp.Connection
	channel    *amqp.Channel
}
