package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"go-ride/client/menu"
	pb "go-ride/proto"
)

func loadTLSCredentials() (credentials.TransportCredentials, error) {
	// Load certificate of the CA who signed server's certificate
	pemServerCA, err := ioutil.ReadFile("../keys/ca-cert.pem")
	if err != nil {
		return nil, err
	}

	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(pemServerCA) {
		return nil, fmt.Errorf("[ERROR] Failed to add CA's certificate")
	}

	// Create the credentials and return it
	config := &tls.Config{
		RootCAs: certPool,
	}

	return credentials.NewTLS(config), nil
}

func main() {
	const address = "localhost:8000"

	fmt.Printf("** Starting go-ride client and setting up server connection on %s **\n", address)

	tlsCredentials, err := loadTLSCredentials()
	if err != nil {
		log.Fatal("[ERROR] Could not load TLS credentials: ", err)
	}

	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(tlsCredentials))
	if err != nil {
		log.Fatalf("[ERROR] Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewRideServiceClient(conn)

	menu.ShowStartMenu(&client)
}
