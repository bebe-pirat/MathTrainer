package model

import "time"

type School struct {
	Id         int       `json:"id"`
	Name       string    `json:"name"`
	Address    string    `json:"address"`
	Created_at time.Time `json:"created_at"`
}
