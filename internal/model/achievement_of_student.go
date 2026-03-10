package model

import "time"

type AchievementOfStudent struct {
	StudentId     int
	AchievementId int
	GotAt         time.Time
}
