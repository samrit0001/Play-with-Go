package main

import (
	"basic/greetproto"
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	//connect to gRPC server
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	//Below one is depracated
	//grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect", err)
	}
	defer conn.Close()

	//create gRPC Client
	client := greetproto.NewGreetServiceClient(conn)

	//Prepare the request
	req := &greetproto.GreetRequest{Name: "ammy"}

	//call the remote function
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.SayHello(ctx, req)
	if err != nil {
		log.Fatalf("Error calling Say Hello Function : %v", err)
	}

	fmt.Println("Server Response :", res.GetMessage())

}
