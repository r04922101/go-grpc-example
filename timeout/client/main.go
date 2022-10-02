package main

import (
	"context"
	"flag"
	"log"
	"time"

	pb "github.com/r04922101/go-grpc-example/timeout"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	serverAddress = flag.String("serverAddress", "localhost:8080", "the server address to connect to")
)

func sendRequest(c pb.HelloServiceClient, timeout bool) {
	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	name := "Tony Huang"
	if timeout {
		name = "timeout"
	}
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("failed to send request: %v", err)
	}
	log.Printf("server response: %s", r.GetMessage())
}

func main() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*serverAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewHelloServiceClient(conn)
	sendRequest(c, false)
	sendRequest(c, true)
}
