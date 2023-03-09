package domain

import "time"

type Client struct {
	Username     string
	PasswordHash string // password hash to simulate secure database storage
	Salt         []byte

	// service info
	Balance uint
	OnBike  bool
	AtZone  string

	// logs for every ride
	Logs []*Log
}

// users can have more than one session in multiple processes
type Session struct {
	Client     *Client
	AuthToken  string
	LastAccess time.Time
}
