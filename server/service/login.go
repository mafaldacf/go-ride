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

func (s RideServiceServer) Login(ctx context.Context, request *pb.LoginRequest) (*pb.LoginResponse, error) {
	log.Printf("[Login] %v\n", request)

	var username = request.Username
	var password = request.Password

	userPtr, found := users[username]

	if !found {
		return nil, status.Errorf(codes.NotFound, "User %s does not exist", username)
	}

	var salt = userPtr.Salt
	match, err := utils.ValidatePassword(userPtr.PasswordHash, password, salt)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "Unexpected error while validating password")
	} else if !match {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid password")
	}

	token, err := utils.GenerateAuthToken(username)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "Unexpected error while generating authentication token")
	}

	var session = dmn.Session{User: userPtr, AuthToken: token, LastAccess: time.Now()}
	tokens[token] = &session

	return &pb.LoginResponse{
		AuthToken: token,
	}, nil
}
