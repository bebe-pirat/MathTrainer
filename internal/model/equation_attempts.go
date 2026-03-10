package model

import "time"

type EquationAttempts struct {
	Id          int
	StudentId   int
	EquationId  int
	GivenAnswer string
	Correct     bool
	AttemptedAt time.Time
}
