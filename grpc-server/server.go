package main

import (
	"context"
	"log"
	"net"

	pb "github.com/iagoholekdev/go-rabbitmq-grpc/grpc-server/publisher"
	"github.com/streadway/amqp"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedPublisherServiceServer
	connection *amqp.Connection
	channel    *amqp.Channel
}

func (s *server) initRabbitMQ() error {
	var err error
	s.connection, err = amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		return err
	}

	s.channel, err = s.connection.Channel()
	if err != nil {
		return err
	}

	return nil
}

func (s *server) PublishMessage(ctx context.Context, req *pb.PublishRequest) (*pb.PublishResponse, error) {
	err := s.channel.Publish(
		"",
		"QueueIago",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(req.Message),
		},
	)
	if err != nil {
		return nil, err
	}
	return &pb.PublishResponse{Success: true}, nil
}

func main() {
	s := &server{}
	if err := s.initRabbitMQ(); err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer s.connection.Close()
	defer s.channel.Close()

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterPublisherServiceServer(grpcServer, s)

	log.Println("gRPC server listening on :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
