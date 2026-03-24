package model

import "time"

type User struct {
	Id           int
	Email        string
	Login        string
	PasswordHash []byte
	RoleId       int
	Blocked      bool
	FullName     string
	ClassId      int
	CreatedAt    time.Time
	LastLogin    time.Time
}

type UserCredentials struct {
	Login    string
	Password string
}
