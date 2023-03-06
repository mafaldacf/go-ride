package service

import (
	"context"
	"log"

	pb "go-ride/proto"
)

func (s RideServiceServer) AddFunds(ctx context.Context, request *pb.AddFundsRequest) (*pb.AddFundsResponse, error) {
	log.Printf("[Add Funds] %v\n", request)

	sessionPtr, err := validateAuthToken(request.AuthToken)

	if err != nil {
		return nil, err
	}

	var userPtr = sessionPtr.User
	userPtr.Balance += request.Amount
	var balance = userPtr.Balance

	return &pb.AddFundsResponse{
		Balance: balance,
	}, nil
}
