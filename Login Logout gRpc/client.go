package main

import (
	"app/authproto"
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Connection not established : %v", conn)
	}

	defer conn.Close()

	client := authproto.NewSecurityClient(conn)

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("\n Choose Operation :")
		fmt.Println("1. Login")
		fmt.Println("2. Logout")
		fmt.Println("3. Exit")
		fmt.Println("Enter your choice: ")
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			handleLogin(client, reader)
		case "2":
			handleLogout(client, reader)
		case "3":
			fmt.Println("Exit Client...")
			return
		default:
			fmt.Println("Sorry no such operation")
		}
	}
}

func handleLogin(client authproto.SecurityClient, reader *bufio.Reader) {
	fmt.Println("Enter Username:")
	username, _ := reader.ReadString('\n')
	username = strings.TrimSpace(username)

	fmt.Print("Enter Password: ")
	password, _ := reader.ReadString('\n')
	password = strings.TrimSpace(password)

	req := &authproto.LoginRequest{
		Username: username,
		Password: password,
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	// Call the gRPC server
	res, err := client.Login(ctx, req)
	if err != nil {
		log.Printf("Login failed: %v", err)
		return
	}

	fmt.Println("Server Response:", res.GetMessage())

}

func handleLogout(client authproto.SecurityClient, reader *bufio.Reader) {
	fmt.Print("Enter Username: ")
	username, _ := reader.ReadString('\n')
	username = strings.TrimSpace(username)

	// Prepare the request
	req := &authproto.LogoutRequest{
		Username: username,
	}

	// Set timeout for the request
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	// Call the gRPC server
	res, err := client.Logout(ctx, req)
	if err != nil {
		log.Printf("Logout failed: %v", err)
		return
	}

	fmt.Println("Server Response:", res.GetMessage())
}
