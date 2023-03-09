package domain

import "time"

type Zone struct {
	Name     string
	Bikes    int // number of current bikes
	Capacity int // maximum capacity for bikes
}

type Log struct {
	Username  string
	FromZone  string
	ToZone    string
	Timestamp time.Time
}
