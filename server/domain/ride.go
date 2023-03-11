package domain

import "time"

type Zone struct {
	Name     string
	Bikes    int // number of current bikes
	Capacity int // maximum capacity for bikes
}

// enumeration of actions
const (
	MOVE        = 0
	PICKUP_BIKE = 1
	DROP_BIKE   = 2
)

type Log struct {
	Timestamp time.Time
	Action    int

	// optionals depending on the context
	FromZone string
	ToZone   string
	Zone     string
}
