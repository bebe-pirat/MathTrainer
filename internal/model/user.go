package model

import "time"

type User struct {
	Id           int       `json:"id"`
	Email        string    `json:"email"`
	Login        string    `json:"login"`
	PasswordHash []byte    `json:"password_hash"`
	RoleId       int       `json:"role_id"`
	Blocked      bool      `json:"blocked"`
	FullName     string    `json:"fullname"`
	ClassId      int       `json:"class_id"`
	CreatedAt    time.Time `json:"created_at"`
	LastLogin    *time.Time `json:"last_login"`
}

type UserCredentials struct {
	Login    string
	Password string
}
