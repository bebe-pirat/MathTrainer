package model

type StudentProfile struct {
	ID       int
	FullName string

	SchoolName string
	ClassName  string
}

type Recommendation struct {
	EquationType string
	Accuracy     float64
}
