package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/r04922101/go-grpc-example/oneof"
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
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	if n, ok := in.GetName().(*pb.HelloRequest_FirstName); ok {
		fmt.Printf("Get first name %s\n", n.FirstName)
	} else if n, ok := in.GetName().(*pb.HelloRequest_LastName); ok {
		fmt.Printf("Get last name %s\n", n.LastName)
	} else {
		fmt.Println("unknown name")
	}
	fmt.Printf("directly use accessor, name: %v, first name: %s, last name: %s\n", in.GetName(), in.GetFirstName(), in.GetLastName())

	if g, ok := in.GetGender().(*pb.HelloRequest_Male); ok {
		fmt.Printf("Get male %v\n", g.Male)
	} else if g, ok := in.GetGender().(*pb.HelloRequest_Female); ok {
		fmt.Printf("Get female %v\n", g.Female)
	} else {
		fmt.Println("unknown gender", in.GetGender())
	}
	fmt.Printf("directly use accessor, gender: %v, male: %v, female: %v", in.GetGender(), in.GetMale(), in.GetFemale())

	if d, ok := in.GetDefaultValues().(*pb.HelloRequest_DefaultInt); ok {
		fmt.Printf("Get int %v\n", d.DefaultInt)
	} else if d, ok := in.GetDefaultValues().(*pb.HelloRequest_DefaultString); ok {
		fmt.Printf("Get string %v\n", d.DefaultString)
	} else {
		fmt.Println("unknown default", in.GetDefaultValues())
	}
	fmt.Printf("directly use accessor, default value: %v, int: %d, string: %s", in.GetDefaultValues(), in.GetDefaultInt(), in.GetDefaultString())

	return &pb.HelloResponse{Message: "Hello!"}, nil
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
