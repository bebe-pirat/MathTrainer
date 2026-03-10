package model

type Equation struct {
	Id             int
	Expression     string
	CorrectAnswer  int
	EquationTypeId int
	Difficulty     int
}
