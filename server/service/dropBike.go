package service

import (
	"context"
	"log"

	pb "go-ride/proto"
	dmn "go-ride/server/domain"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s RideServiceServer) DropBike(ctx context.Context, request *pb.AuthRequest) (*pb.Empty, error) {
	log.Printf("[Drop Bike] %v\n", request)

	sessionPtr, err := s.ValidateAuthToken(request.AuthToken)
	if err != nil {
		return nil, err
	}

	clientPtr := sessionPtr.Client
	if !clientPtr.OnBike {
		return nil, status.Error(codes.FailedPrecondition, "You are not riding a bike")
	}

	zone := clientPtr.AtZone
	zonePtr := s.Zones[zone]
	if zonePtr.Bikes == zonePtr.Capacity {
		return nil, status.Errorf(codes.NotFound, "There is currently no free stop to drop your bike at %s", zone)
	}

	zonePtr.Bikes++
	clientPtr.OnBike = false
	clientPtr.AtZone = "None"
	s.logAction(clientPtr, dmn.DROP_BIKE, "", "", zone)

	return &pb.Empty{}, nil
}
