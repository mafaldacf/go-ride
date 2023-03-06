package domain

import (
	"time"
)

// users can have more than one session in multiple processes
type Session struct {
	User       *User
	AuthToken  string
	LastAccess time.Time
}
