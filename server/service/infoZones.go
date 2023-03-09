package service

import (
	"context"
	"log"
	"sort"
	pb "go-ride/proto"
)

// we need to sort manually because in Go elements of a map are not sorted since they are implemented as hash tables
func (s RideServiceServer) sortZones() []string {
	zones := make([]string, 0, len(s.Zones))

	for key := range s.Zones {
		zones = append(zones, key)
	}

	sort.Strings(zones)
	return zones
}

func (s RideServiceServer) InfoZones(ctx context.Context, request *pb.AuthRequest) (*pb.InfoZonesResponse, error) {
	log.Printf("[Info Zones] %v\n", request)

	_, err := s.ValidateAuthToken(request.AuthToken)
	if err != nil {
		return nil, err
	}

	response := &pb.InfoZonesResponse{}
	zones := s.sortZones()

	for _, zone := range zones {
		response.Zones = append(response.Zones, &pb.Zone{
			Name:     zone,
			Bikes:    int32(s.Zones[zone].Bikes),
			Capacity: int32(s.Zones[zone].Capacity),
		})
	}

	return response, nil
}
