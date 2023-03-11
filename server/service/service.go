package service

import (
	pb "go-ride/proto"
	"log"
	"time"

	dmn "go-ride/server/domain"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const MIN_BALANCE = 5

type RideServiceServer struct {
	pb.UnimplementedRideServiceServer

	// lets deal with pointers since it is more efficient and it is easier to change values
	Clients  map[string]*dmn.Client
	Sessions map[string]*dmn.Session
	Zones    map[string]*dmn.Zone
}

func (s RideServiceServer) logAction(client *dmn.Client, action int, fromZone string, toZone string, zone string) {
	log := dmn.Log{Timestamp: time.Now(), Action: action, FromZone: fromZone, ToZone: toZone, Zone: zone}
	client.Logs = append(client.Logs, log)
}

func (s RideServiceServer) verifySessions() {
	for key, value := range s.Sessions {
		offtime := time.Since(value.LastAccess)                     // already in seconds
		if offtime > 0 && time.Duration(offtime) > 60*time.Second { // only 1 minute to showcase the functionality
			log.Printf("--- Invalidating auth token of %s\n", value.Client.Username)
			delete(s.Sessions, key) // user did not perform any action so invalidate its AuthToken
		}
	}
}

func (s RideServiceServer) InitSessionVerifier() {
	go func() { // start goroutine aka thread to verify sessions every second
		for {
			s.verifySessions()
			time.Sleep(1 * time.Second)
		}
	}()
}

func (s RideServiceServer) ValidateAuthToken(token string) (*dmn.Session, error) {
	sessionPtr, exists := s.Sessions[token]

	if !exists {
		return nil, status.Errorf(codes.Unauthenticated, "Session expired. Please login again.")
	}

	sessionPtr.LastAccess = time.Now() // update access time

	return sessionPtr, nil
}
