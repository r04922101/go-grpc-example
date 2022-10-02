package main

import (
	"context"
	"flag"
	"io"
	"log"
	"time"

	pb "github.com/r04922101/go-grpc-example/serverstreaming"

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
	stream, err := c.SayHello(ctx, &pb.HelloRequest{Name: "Tony Huang"})
	if err != nil {
		log.Fatalf("failed to send request and get stream: %v", err)
	}

	for {
		r, err := stream.Recv()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatalf("failed to receive response: %v", err)
		}
		log.Printf("server response %s", r.GetMessage())
	}
}
