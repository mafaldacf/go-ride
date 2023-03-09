package service

import (
	"context"
	"log"
	"time"

	pb "go-ride/proto"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	dmn "go-ride/server/domain"
	"go-ride/server/utils"
)

func (s RideServiceServer) Register(ctx context.Context, request *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	log.Printf("[Register] %v\n", request)

	username := request.Username
	password := request.Password

	_, found := s.Clients[username]

	if found {
		return nil, status.Errorf(codes.AlreadyExists, "User '%s' already exists", username)
	}

	salt, err := utils.GenerateRandomSalt()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Unexpected error while processing password")
	}

	passwordHash, err := utils.HashPasswordSalt(password, salt)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Unexpected error while processing password")
	}

	client := dmn.Client{Username: username, PasswordHash: passwordHash, Salt: salt, OnBike: false, AtZone: "None", Logs: make([]*dmn.Log, 0)}
	s.Clients[username] = &client

	token, err := utils.GenerateAuthToken(username)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "Unexpected error while generating authentication token")
	}

	session := dmn.Session{Client: &client, AuthToken: token, LastAccess: time.Now()}
	s.Sessions[token] = &session

	return &pb.RegisterResponse{
		AuthToken: token,
	}, nil
}
