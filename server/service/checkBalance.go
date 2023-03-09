package service

import (
	"context"
	"log"

	pb "go-ride/proto"
)

func (s RideServiceServer) CheckBalance(ctx context.Context, request *pb.AuthRequest) (*pb.CheckBalanceResponse, error) {
	log.Printf("[Check Balance] %v\n", request)

	sessionPtr, err := s.ValidateAuthToken(request.AuthToken)
	if err != nil {
		return nil, err
	}

	clientPtr := sessionPtr.Client
	balance := clientPtr.Balance

	return &pb.CheckBalanceResponse{
		Balance: uint32(balance),
	}, nil
}
