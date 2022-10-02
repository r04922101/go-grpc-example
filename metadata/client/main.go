package main

import (
	"context"
	"flag"
	"log"
	"time"

	pb "github.com/r04922101/go-grpc-example/metadata"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
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
	// send metadata to server
	ctx = metadata.AppendToOutgoingContext(ctx, "client-meta-key1", "client-meta", "client-meta-key2-bin", string([]byte{72, 105}))
	defer cancel()

	var header, trailer metadata.MD
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: "Tony Huang"}, grpc.Header(&header), grpc.Trailer(&trailer))
	if err != nil {
		log.Fatalf("failed to send request: %v", err)
	}

	log.Printf("server response: %s\n", r.GetMessage())
	log.Println("with metadata header:")
	for k, v := range header {
		log.Printf("key: %s, value: %s\n", k, v)
	}
	log.Println("and metadata trailer:")
	for k, v := range trailer {
		log.Printf("key: %s, value: %s\n", k, v)
	}
}
