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

	username := request.Username
	password := request.Password

	clientPtr, found := s.Clients[username]
	if !found {
		return nil, status.Errorf(codes.NotFound, "User %s does not exist", username)
	}

	salt := clientPtr.Salt
	match, err := utils.ValidatePassword(clientPtr.PasswordHash, password, salt)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Unexpected error while validating password")
	} else if !match {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid password")
	}

	token, err := utils.GenerateAuthToken(username)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Unexpected error while generating authentication token")
	}

	session := dmn.Session{Client: clientPtr, AuthToken: token, LastAccess: time.Now()}
	s.Sessions[token] = &session

	return &pb.LoginResponse{
		AuthToken: token,
	}, nil
}
