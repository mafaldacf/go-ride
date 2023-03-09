package service

import (
	"context"
	"log"

	pb "go-ride/proto"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s RideServiceServer) PickupBike(ctx context.Context, request *pb.PickupBikeRequest) (*pb.Empty, error) {
	log.Printf("[Pickup Bike] %v\n", request)

	sessionPtr, err := s.ValidateAuthToken(request.AuthToken)
	if err != nil {
		return nil, err
	}

	zonePtr := s.Zones[request.Zone]
	if zonePtr.Bikes == 0 {
		return nil, status.Errorf(codes.NotFound, "There are currently no bikes to pickup in %s", request.Zone)
	}

	clientPtr := sessionPtr.Client
	if clientPtr.Balance < MIN_BALANCE {
		return nil, status.Errorf(codes.FailedPrecondition, "Insufficient funds! You need at least %d\u20AC to pickup a bike.", MIN_BALANCE)
	}

	zonePtr.Bikes--
	clientPtr.OnBike = true
	clientPtr.AtZone = request.Zone

	return &pb.Empty{}, nil
}
