package service

import (
	"context"
	"log"

	pb "go-ride/proto"
)

func (s RideServiceServer) CheckBalance(ctx context.Context, request *pb.CheckBalanceRequest) (*pb.CheckBalanceResponse, error) {
	log.Printf("[Check Balance] %v\n", request)

	sessionPtr, err := validateAuthToken(request.AuthToken)

	if err != nil {
		return nil, err
	}

	var userPtr = sessionPtr.User
	var balance = userPtr.Balance

	return &pb.CheckBalanceResponse{
		Balance: balance,
	}, nil
}
