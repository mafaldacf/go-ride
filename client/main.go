package main

import (
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"go-ride/client/menu"
	pb "go-ride/proto"
)

func main() {
	const address = "localhost:8000"

	fmt.Printf("** Starting go-ride client and setting up server connection on %s **\n", address)

	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewRideServiceClient(conn)

	menu.ShowStartMenu(&client)
}
