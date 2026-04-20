package model

type StudentStats struct {
	TotalAttempts  int     `json:"total_attempts"`
	CorrectAnswers int     `json:"correct_answers"`
	WrongAnswers   int     `json:"wrong_answers"`
	Accuracy       float32 `json:"accuracy_percent"`

	LevelsCompleted int `json:"levels_completed"`
	StarsTotal      int `json:"stars_total"`
	XP              int `json:"xp"`

	EquationTypes []ExtendedEquationTypeStats `json:"equation_type_stats"`
	Achievements  []AchievementOfStudent      `json:"achievements"`

	WeakTopics []string `json:"weak_types"`
}

type ShortEquationTypeStats struct {
	TypeId   int
	Accuracy float32
}

type ExtendedEquationTypeStats struct {
	Type     string `json:"type"`
	Attempts int    `json:"attempts"`
	Correct  int    `json:"correct"`
	Wrong    int    `json:"wrong"`
	Accuracy int    `json:"accuracy_percent"`
}

type ClassStats struct {
	StudentsCount  int     `json:"students_count"`
	TotalAttempts  int     `json:"total_attempts"`
	CorrectAnswers int     `json:"correct_answers"`
	WrongAnswers   int     `json:"wrong_answers"`
	Accuracy       float32 `json:"accuracy_percent"`

	EquationTypes []EquationTypeStats `json:"equation_types_stats"`
	Students      []StudentShortStats `json:"students"`
}

type EquationTypeStats struct {
	Type     string `json:"type"`
	Accuracy int    `json:"accuracy_percent"`
}

type StudentShortStats struct {
	StudentId       int `json:"student_id"`
	Name            int `json:"name"`
	Accuracy        int `json:"accuracy"`
	LevelsComplited int `json:"levels_complited"`
}

type SchoolStats struct {
	StudentsCount       int     `json:"students_count"`
	ClassesCount        int     `json:"classes_count"`
	TotalEquationSolved int     `json:"total_equation_solved"`
	WrongAnswers        int     `json:"wrong_answers"`
	Accuracy            float32 `json:"accuracy_percent"`

	EquationTypes []EquationTypeStats `json:"equation_types"`
	Classes       []ClassShortStats   `json:"classes"`
}

type ClassShortStats struct {
	Name     string `json:"name"`
	Accuracy int    `json:"accuracy_percent"`
}
