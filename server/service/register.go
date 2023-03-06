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

	var username = request.Username
	var password = request.Password

	_, found := users[username]

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

	var user = dmn.User{Username: username, PasswordHash: passwordHash, Salt: salt}
	users[username] = &user

	token, err := utils.GenerateAuthToken(username)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "Unexpected error while generating authentication token")
	}

	var session = dmn.Session{User: &user, AuthToken: token, LastAccess: time.Now()}
	tokens[token] = &session

	return &pb.RegisterResponse{
		AuthToken: token,
	}, nil
}
