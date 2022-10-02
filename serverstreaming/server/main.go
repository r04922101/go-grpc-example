package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/r04922101/go-grpc-example/serverstreaming"
	"google.golang.org/grpc"
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
func (s *server) SayHello(req *pb.HelloRequest, stream pb.HelloService_SayHelloServer) error {
	for i := 0; i < 3; i++ {
		stream.Send(&pb.HelloResponse{Message: fmt.Sprintf("[%d] Hello %s !", i, req.GetName())})
	}
	// tells gRPC to finish writing responses
	return nil
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
