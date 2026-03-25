package model

type SessionData struct {
	SessionID int `json:"session_id"`
	UserID    int `json:"user_id"`
	Role      int `json:"role"`
}
