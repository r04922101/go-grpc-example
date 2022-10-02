// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.6
// source: clientstreaming/hello.proto

package clientstreaming

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// HelloServiceClient is the client API for HelloService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HelloServiceClient interface {
	// Sends a greeting
	SayHello(ctx context.Context, opts ...grpc.CallOption) (HelloService_SayHelloClient, error)
}

type helloServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHelloServiceClient(cc grpc.ClientConnInterface) HelloServiceClient {
	return &helloServiceClient{cc}
}

func (c *helloServiceClient) SayHello(ctx context.Context, opts ...grpc.CallOption) (HelloService_SayHelloClient, error) {
	stream, err := c.cc.NewStream(ctx, &HelloService_ServiceDesc.Streams[0], "/clientstreaming.HelloService/SayHello", opts...)
	if err != nil {
		return nil, err
	}
	x := &helloServiceSayHelloClient{stream}
	return x, nil
}

type HelloService_SayHelloClient interface {
	Send(*HelloRequest) error
	CloseAndRecv() (*HelloResponse, error)
	grpc.ClientStream
}

type helloServiceSayHelloClient struct {
	grpc.ClientStream
}

func (x *helloServiceSayHelloClient) Send(m *HelloRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *helloServiceSayHelloClient) CloseAndRecv() (*HelloResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(HelloResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// HelloServiceServer is the server API for HelloService service.
// All implementations must embed UnimplementedHelloServiceServer
// for forward compatibility
type HelloServiceServer interface {
	// Sends a greeting
	SayHello(HelloService_SayHelloServer) error
	mustEmbedUnimplementedHelloServiceServer()
}

// UnimplementedHelloServiceServer must be embedded to have forward compatible implementations.
type UnimplementedHelloServiceServer struct {
}

func (UnimplementedHelloServiceServer) SayHello(HelloService_SayHelloServer) error {
	return status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (UnimplementedHelloServiceServer) mustEmbedUnimplementedHelloServiceServer() {}

// UnsafeHelloServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HelloServiceServer will
// result in compilation errors.
type UnsafeHelloServiceServer interface {
	mustEmbedUnimplementedHelloServiceServer()
}

func RegisterHelloServiceServer(s grpc.ServiceRegistrar, srv HelloServiceServer) {
	s.RegisterService(&HelloService_ServiceDesc, srv)
}

func _HelloService_SayHello_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(HelloServiceServer).SayHello(&helloServiceSayHelloServer{stream})
}

type HelloService_SayHelloServer interface {
	SendAndClose(*HelloResponse) error
	Recv() (*HelloRequest, error)
	grpc.ServerStream
}

type helloServiceSayHelloServer struct {
	grpc.ServerStream
}

func (x *helloServiceSayHelloServer) SendAndClose(m *HelloResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *helloServiceSayHelloServer) Recv() (*HelloRequest, error) {
	m := new(HelloRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// HelloService_ServiceDesc is the grpc.ServiceDesc for HelloService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HelloService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "clientstreaming.HelloService",
	HandlerType: (*HelloServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SayHello",
			Handler:       _HelloService_SayHello_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "clientstreaming/hello.proto",
}
