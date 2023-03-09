package service

import (
	"context"
	"log"

	pb "go-ride/proto"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Alpha <---> Bravo <---> Charlie <---> Delta <---> Echo <---> Foxtrot
// 1 'step' costs 1 euro
// e.g. Alpha --> Foxtrot costs 5 euro

func computeCost(currZone string, nextZone string) uint {
	// get first letter of each zone
	currZoneAbrv := currZone[0]
	nextZoneAbrv := nextZone[0]

	// calculate absolute difference between ASCII codes
	diff := int(nextZoneAbrv) - int(currZoneAbrv)

	// Go does not have built-in function to calculate abs of int values
	if diff < 0 {
		return uint(-diff)
	}
	return uint(diff)
}

func (s RideServiceServer) Move(ctx context.Context, request *pb.MoveRequest) (*pb.Empty, error) {
	log.Printf("[Move] %v\n", request)

	sessionPtr, err := s.ValidateAuthToken(request.AuthToken)
	if err != nil {
		return nil, err
	}

	clientPtr := sessionPtr.Client
	if !clientPtr.OnBike {
		return nil, status.Error(codes.FailedPrecondition, "You are not riding a bike yet")
	}

	currZone := clientPtr.AtZone
	nextZone := request.Zone
	if currZone == nextZone {
		return nil, status.Errorf(codes.InvalidArgument, "You are already at %s", nextZone)
	}

	cost := computeCost(currZone, nextZone)
	if cost > clientPtr.Balance {
		return nil, status.Errorf(codes.FailedPrecondition, "Insufficient funds to move from %s to %s. You need at least %d\u20AC and you currently have %d\u20AC", currZone, nextZone, cost, clientPtr.Balance)
	}

	clientPtr.AtZone = nextZone
	clientPtr.Balance -= cost

	return &pb.Empty{}, nil
}
