package model

import "time"

type School struct {
	Id         int
	Name       string
	Address    string
	Created_at time.Time
}
