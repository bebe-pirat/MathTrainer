package model

import "time"

type Class struct {
	Id        int
	Name      string
	Grade     int
	SchoolId  int
	CreatedAt time.Time
}
