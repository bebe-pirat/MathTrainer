package model

import "time"

type User struct {
	Id           int
	Email        string
	Login        string
	PasswordHash string
	RoleId       int
	Blocked      bool
	FullName     string
	ClassId      int
	CreatedAt    time.Time
	LastLogin    time.Time
}
