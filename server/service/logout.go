package service

import (
	"context"
	"log"

	pb "go-ride/proto"
)

func (s RideServiceServer) Logout(ctx context.Context, request *pb.LogoutRequest) (*pb.Empty, error) {
	log.Printf("[Logout] %v\n", request)

	var token = request.AuthToken
	delete(tokens, token)

	return &pb.Empty{}, nil
}
