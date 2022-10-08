package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/r04922101/go-grpc-example/customoption"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
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
	messageAllowNil := proto.GetExtension(in.ProtoReflect().Descriptor().Options(), pb.E_AllowNilValues)
	fmt.Println("message allow nil option:", messageAllowNil)
	// TODO: use this option to do something...
	// TODO: method, field option
	return &pb.HelloResponse{Message: "Hello " + in.GetName() + "!"}, nil
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
