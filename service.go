package main

import (
	"context"
	example "example_generated"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"log"
	"net"
)

type exampleServer struct{}

func newExampleServer() example.OperationServiceServer {
	return new(exampleServer)
}

func (s* exampleServer) PerformOperation(ctx context.Context, msg *example.OperationRequest) (*example.OperationResponse, error) {
	var result int64 = 0
	permitted := true
	switch msg.OperationType {
	case "Addition":
		result = int64(msg.Val1 + msg.Val2)
	case "Subtraction":
		result = int64(msg.Val1 - msg.Val2)
	case "Division":
		result = int64(msg.Val1 / msg.Val2)
	case "Multiplication":
		result = int64(msg.Val1 * msg.Val2)
	default:
		permitted = false
	}
	grpc.SendHeader(ctx, metadata.New(map[string]string{
		"foo": "foo1",
		"bar": "bar1",
	}))
	grpc.SetTrailer(ctx, metadata.New(map[string]string{
		"foo": "foo2",
		"bar": "bar2",
	}))
	return &example.OperationResponse {
		Response: result,
		OperationValid: permitted} , nil
}

func main() {
	listen, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatal(err)
	}
	server := grpc.NewServer()
	example.RegisterOperationServiceServer(server, newExampleServer())
	log.Printf("Started gRPC server on port 9090...\n")
	log.Fatal(server.Serve(listen))
}