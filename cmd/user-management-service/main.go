package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	usermanagement "github.com/brunoeduardodev/go-grpc-ecommerce/protocols"
	"google.golang.org/grpc"
)

var port = flag.Int("port", 50051, "The server port")

type server struct {
	usermanagement.UserManagementServer
}

func (s *server) CreateUser(ctx context.Context, input *usermanagement.CreateUserRequest) (*usermanagement.CreateUserResponse, error) {
	return &usermanagement.CreateUserResponse{
		Id: "123",
	}, nil
}

func main() {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))

	if err != nil {
		log.Fatalf("Could not start tcp listener: %v", err)
	}

	s := grpc.NewServer()
	usermanagement.RegisterUserManagementServer(s, &server{})
	log.Printf("server listening at %v", listener.Addr())

	err = s.Serve(listener)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
