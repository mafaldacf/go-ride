package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "go-ride/proto"

	svc "go-ride/server/service"
)

func main() {
	const address = "localhost:8000"
	fmt.Printf("** Starting go-ride server on %s **\n", address)

	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Could not create listener: %s\n", err)
	}

	server := grpc.NewServer()
	service := &svc.RideServiceServer{}

	pb.RegisterRideServiceServer(server, service)

	err = server.Serve(lis)
	if err != nil {
		log.Fatalf("Could not serve server: %s\n", err)
	}
}
