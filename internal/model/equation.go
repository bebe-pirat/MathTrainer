package model

type Equation struct {
	Id             int    `json:"id"`
	Text           string `json:"equation_text"`
	CorrectAnswer  int    `json:"correct_answer"`
	EquationTypeId int    `json:"equation_type_id"`
}

type Answer struct {
	EquationId     int    `json:"equation_id"`
	Text           string `json:"equation_text"`
	CorrectAnswer  int    `json:"correct_answer"`
	UserAnswer     int    `json:"user_answer"`
	EquationTypeId int    `json:"equation_type_id"`
}

type EquationFeedback struct {
	EquationId    int    `json:"equation_id"`
	IsCorrect     bool   `json:"is_correct"`
	CorrectAnswer int    `json:"correct_answer"`
	Feedback      string `json:"feedback"`
}

type StarsAndXP struct {
	SectionId  int `json:"section_id"`
	LevelOrder int `json:"level_order"`
	Stars      int `json:"stars"`
	CommonXP   int `json:"common_xp"`
}
