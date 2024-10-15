package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"

	pb "github.com/iagoholekdev/go-rabbitmq-grpc/grpc-server/publisher"
	"google.golang.org/grpc"
)

func main() {

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewPublisherServiceClient(conn)

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter message to publish: ")

	fmt.Print("Enter message to publish: ")
	if scanner.Scan() {
		message := scanner.Text()

		resp, err := client.PublishMessage(context.Background(), &pb.PublishRequest{Message: message})
		if err != nil {
			log.Fatalf("could not publish message: %v", err)
		}

		fmt.Printf("Message published successfully: %v\n", resp.Success)
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading input: %v", err)
	}

}
