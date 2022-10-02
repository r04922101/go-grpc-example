package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log"
	"net"
	"time"

	pb "github.com/r04922101/go-grpc-example/timeout"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	port = flag.Int("port", 8080, "The server port")
)

// server is used to implement HelloServiceServer
type server struct {
	// UnimplementedHelloServiceServer must be embedded to have forward compatible implementations.
	pb.UnimplementedHelloServiceServer
}

// SayHello implements HelloServiceServer.SayHello
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	name := in.GetName()
	log.Printf("received message from client \"%v\" ", name)

	// cause client timeout
	if name == "timeout" {
		time.Sleep(2 * time.Second)
	}

	// check request still valid or not
	if err := ctx.Err(); errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) {
		fmt.Printf("context canceled: %v\n", err)
		return nil, status.New(codes.DeadlineExceeded, "deadline exceeded or client canceled").Err()
	}

	return &pb.HelloResponse{Message: "Hello " + name + "!"}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen on port %d: %v", *port, err)
	}

	s := grpc.NewServer()
	pb.RegisterHelloServiceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
