package model

import "time"

type EquationAttempts struct {
	Id             int
	StudentId      int
	EquationTypeId int
	GivenAnswer    int
	CorrectAnswer  int
	AnsweredAt     time.Time
}
