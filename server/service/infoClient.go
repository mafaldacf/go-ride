package service

import (
	"context"
	"log"

	pb "go-ride/proto"
)

func (s RideServiceServer) InfoClient(ctx context.Context, request *pb.AuthRequest) (*pb.InfoClientResponse, error) {
	log.Printf("[Info Client] %v\n", request)

	sessionPtr, err := s.ValidateAuthToken(request.AuthToken)
	if err != nil {
		return nil, err
	}

	clientPtr := sessionPtr.Client
	return &pb.InfoClientResponse{
		OnBike: clientPtr.OnBike,
		AtZone: clientPtr.AtZone,
	}, nil
}
