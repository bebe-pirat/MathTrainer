package model

import "time"

type Class struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Grade     int       `json:"grade"`
	SchoolId  int       `json:"school_id"`
	CreatedAt time.Time `json:"created_at"`
}
