package main

import (
	"basic/greetproto"
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	greetproto.UnimplementedGreetServiceServer
}

// Implement server Method
func (s *server) SayHello(ctx context.Context, req *greetproto.GreetRequest) (*greetproto.GreetResponse, error) {
	name := req.GetName()
	message := fmt.Sprintf("Hello ,%s", name)
	return &greetproto.GreetResponse{Message: message}, nil
}

func main() {
	//start listening on port 50051
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcserver := grpc.NewServer()
	greetproto.RegisterGreetServiceServer(grpcserver, &server{})

	fmt.Println("gRPC Server Runnig on port 50051...")
	if err := grpcserver.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
