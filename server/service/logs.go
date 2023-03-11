package service

import (
	"context"
	pb "go-ride/proto"
	"log"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s RideServiceServer) CheckLogs(ctx context.Context, request *pb.AuthRequest) (*pb.CheckLogsResponse, error) {
	log.Printf("[Check Logs] %v\n", request)

	sessionPtr, err := s.ValidateAuthToken(request.AuthToken)
	if err != nil {
		return nil, err
	}

	clientPtr := sessionPtr.Client
	response := &pb.CheckLogsResponse{}

	for _, log := range clientPtr.Logs {
		response.Logs = append(response.Logs, &pb.Log{
			Timestamp: 	timestamppb.New(log.Timestamp),
			Action:    	pb.Action(log.Action),
			Zone:		log.Zone,
			FromZone: 	log.FromZone,
			ToZone: 	log.ToZone,
		})
	}

	return response, nil
}
