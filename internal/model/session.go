package model

import "time"

type SessionData struct {
	SessionID int `json:"session_id"`
	UserID    int `json:"user_id"`
	Role      int `json:"role"`
}

type Session struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	ExpiresAt time.Time `json:"expires_at"`
}
