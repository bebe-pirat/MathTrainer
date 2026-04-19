package model

import "time"

type StudentProgress struct {
	Id          int
	StudentId   int
	SectionId   int
	LevelOrder  int
	CountStarts int
	FinishedAt  time.Time
}
