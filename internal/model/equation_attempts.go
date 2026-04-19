package model

import "time"

type Attempt struct {
	Id             int
	StudentId      int
	EquationText   string
	EquationTypeId int
	GivenAnswer    int
	CorrectAnswer  int
	AnsweredAt     time.Time
}
