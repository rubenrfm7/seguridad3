package models

import "time"

type Token struct {
	Username  string    `json:"username"`
	ExpiresAt time.Time `json:"expires_at"`
}
