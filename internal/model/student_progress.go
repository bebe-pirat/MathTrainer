package model

import "time"

type StudentProgress struct {
	Id          int
	StudentId   int
	LevelId     int
	CountStarts int
	FinishedAt  time.Time
}
