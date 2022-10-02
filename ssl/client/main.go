package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	pb "github.com/r04922101/go-grpc-example/ssl"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var (
	serverAddress = flag.String("serverAddress", "localhost:8080", "the server address to connect to")
	tlsCert       = flag.String("tlsCert", "ssl/certs/cert.pem", "server certificate path")
)

func grpcAuthDial() (*grpc.ClientConn, error) {
	opts := make([]grpc.DialOption, 0, 1)
	if *tlsCert != "" {
		creds, err := credentials.NewClientTLSFromFile(*tlsCert, *serverAddress)
		if err != nil {
			return nil, fmt.Errorf("failed to create TLS credentials: %v", err)
		}
		opts = append(opts, grpc.WithTransportCredentials(creds))
	} else {
		opts = append(opts, grpc.WithInsecure())
	}
	return grpc.Dial(*serverAddress, opts...)
}

func main() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpcAuthDial()
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewHelloServiceClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: "Tony Huang"})
	if err != nil {
		log.Fatalf("failed to send request: %v", err)
	}
	log.Printf("server response: %s", r.GetMessage())
}
