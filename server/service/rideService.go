package service

import (
	pb "go-ride/proto"

	dmn "go-ride/server/domain"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// lets deal with pointers since it is more efficient to change values
var users = make(map[string]*dmn.User)
var tokens = make(map[string]*dmn.Session)

type RideServiceServer struct {
	pb.UnimplementedRideServiceServer
}

func validateAuthToken(token string) (*dmn.Session, error) {
	session, exists := tokens[token]

	if !exists {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid authentication token")
	}

	return session, nil
}
