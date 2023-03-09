package service

import (
	"context"
	"log"

	pb "go-ride/proto"
)

func (s RideServiceServer) AddFunds(ctx context.Context, request *pb.AddFundsRequest) (*pb.AddFundsResponse, error) {
	log.Printf("[Add Funds] %v\n", request)

	sessionPtr, err := s.ValidateAuthToken(request.AuthToken)
	if err != nil {
		return nil, err
	}

	clientPtr := sessionPtr.Client
	clientPtr.Balance += uint(request.Amount)
	balance := clientPtr.Balance

	return &pb.AddFundsResponse{
		Balance: uint32(balance),
	}, nil
}
