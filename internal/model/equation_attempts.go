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

type AttemptForTeacher struct {
	EquationText     string    `json:"equation_text"`
	EquationTypeId   int       `json:"equation_type_id"`
	EquationTypeName string    `json:"equation_type_name"`
	GivenAnswer      int       `json:"given_answer"`
	CorrectAnswer    int       `json:"correct_answer"`
	AnsweredAt       time.Time `json:"answered_at"`
}
