package model

type StudentProfile struct {
	ID       int    `json:"id"`
	FullName string `json:"fullname"`

	SchoolName string `json:"school_name"`
	ClassName  string `json:"class_name"`

	XP int `json:"xp"`
}

type Recommendation struct {
	EquationType string  `json:"equation_type"`
	Accuracy     float64 `json:"accuracy"`
}
