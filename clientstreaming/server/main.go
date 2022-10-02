package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"strings"

	pb "github.com/r04922101/go-grpc-example/clientstreaming"
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
func (s *server) SayHello(stream pb.HelloService_SayHelloServer) error {
	names := make([]string, 0)
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			// get all requests and send response
			return stream.SendAndClose(&pb.HelloResponse{
				Message: fmt.Sprintf("Hello %s!", strings.Join(names, ", ")),
			})
		} else if err != nil {
			return fmt.Errorf("failed to receive requests from client: %v", err)
		}

		fmt.Printf("received client request: %s\n", req.GetName())
		names = append(names, req.GetName())
	}
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
