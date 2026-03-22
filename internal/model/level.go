package model

type Level struct {
	Id         int
	Name       string
	TestLevel  bool
	Difficulty int
}

type LevelSession struct {
	LevelID  int
	Equation Equation
}

type LevelResult struct {
	Stars          int
	CorrectAnswers int
	WrongAnswers   int
}
