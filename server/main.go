package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	pb "go-ride/proto"
	dmn "go-ride/server/domain"
	svc "go-ride/server/service"
)

func loadTLSCredentials() (credentials.TransportCredentials, error) {
	serverCert, err := tls.LoadX509KeyPair("../keys/server-cert.pem", "../keys/server-key.pem")
	if err != nil {
		return nil, err
	}

	config := &tls.Config{
		Certificates: []tls.Certificate{serverCert},
		ClientAuth:   tls.NoClientCert,
	}

	return credentials.NewTLS(config), nil
}

func main() {
	const address = "localhost:8000"
	fmt.Printf("** Starting go-ride server on %s **\n", address)

	tlsCredentials, err := loadTLSCredentials()
	if err != nil {
		log.Fatal("[ERROR] Could not load TLS credentials: ", err)
	}

	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("[ERROR] Could not create listener: %s\n", err)
	}

	server := grpc.NewServer(grpc.Creds(tlsCredentials))

	service := &svc.RideServiceServer{
		Clients:  make(map[string]*dmn.Client),
		Sessions: make(map[string]*dmn.Session),
		Zones: map[string]*dmn.Zone{
			"Alpha":   {Name: "Alpha", Bikes: 15, Capacity: 20},
			"Bravo":   {Name: "Bravo", Bikes: 15, Capacity: 20},
			"Charlie": {Name: "Charlie", Bikes: 0, Capacity: 10},
			"Delta":   {Name: "Delta", Bikes: 5, Capacity: 10},
			"Echo":    {Name: "Echo", Bikes: 0, Capacity: 5},
			"Foxtrot": {Name: "Foxtrot", Bikes: 5, Capacity: 5},
		},
	}

	service.InitSessionVerifier()

	pb.RegisterRideServiceServer(server, service)

	err = server.Serve(lis)
	if err != nil {
		log.Fatalf("[ERROR] Could not serve listener: %s\n", err)
	}
}
