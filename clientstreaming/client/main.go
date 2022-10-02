package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	pb "github.com/r04922101/go-grpc-example/clientstreaming"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	serverAddress = flag.String("serverAddress", "localhost:8080", "the server address to connect to")
)

func main() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*serverAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewHelloServiceClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	stream, err := c.SayHello(ctx)
	if err != nil {
		log.Fatalf("failed to get stream: %v", err)
	}

	names := []string{"Tony", "Alice", "Bob"}
	for _, name := range names {
		if err := stream.Send(&pb.HelloRequest{
			Name: name,
		}); err != nil {
			log.Fatalf("failed to send request: %v", err)
		}
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("failed to close stream: %v", err)
	}
	fmt.Printf("server response: %s", res.GetMessage())
}
