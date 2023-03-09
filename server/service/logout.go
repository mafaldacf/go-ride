package service

import (
	"context"
	"log"

	pb "go-ride/proto"
)

func (s RideServiceServer) Logout(ctx context.Context, request *pb.AuthRequest) (*pb.Empty, error) {
	log.Printf("[Logout] %v\n", request)

	token := request.AuthToken
	delete(s.Sessions, token) // if token is not a key it will simply do nothing

	return &pb.Empty{}, nil
}
