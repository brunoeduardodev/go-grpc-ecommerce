package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	usermanagement "github.com/brunoeduardodev/go-grpc-ecommerce/protocols"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	createUserCmd := flag.NewFlagSet("user create", flag.ExitOnError)
	createUserEmail := createUserCmd.String("email", "", "The user email")
	createUserPassword := createUserCmd.String("password", "", "The user password")

	userManagementServicePort := flag.Int("user-management-service-port", 50051, "the user management service port")

	if len(os.Args) < 2 {
		log.Fatal("Expected 'user' sub command")
	}

	switch os.Args[1] {
	case "user":
		log.SetPrefix("user:")
		if len(os.Args) < 3 {
			log.Fatal("expected 'create' sub command")
		}
		switch os.Args[2] {
		case "create":
			log.SetPrefix("user create:")
			createUserCmd.Parse(os.Args[3:])
			email := *createUserEmail
			password := *createUserPassword

			if email == "" {
				log.Fatal("'email' argument is required")
			}

			if password == "" {
				log.Fatal("'password' argument is required")
			}

			servicePort := *userManagementServicePort
			conn, err := grpc.Dial(fmt.Sprintf("localhost:%d", servicePort), grpc.WithTransportCredentials(insecure.NewCredentials()))
			if err != nil {
				log.Fatal("Could not connect:", err)
			}

			defer conn.Close()
			c := usermanagement.NewUserManagementClient(conn)
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()

			response, err := c.CreateUser(ctx, &usermanagement.CreateUserRequest{Email: email, Password: password})
			if err != nil {
				log.Fatal("error: ", err)
			}

			log.Println("response: id:", response.GetId())

		default:
			log.Fatal("expected 'create' sub command")
			return
		}

	default:
		log.Fatal("Expected 'user' sub command")
	}

}
