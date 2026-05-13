package model

type Section struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Order       int    `json:"section_order"`
	Class       int    `json:"class,omitempty"`
	LevelsCount int    `json:"levels_count,omitempty"`
	IsUnlocked  bool   `json:"is_unlocked,omitempty"`
}

type StudentPosition struct {
	SectionId  int `json:"section_id"`
	LevelOrder int `json:"level_order"`
}

type LevelsMap struct {
	Sections []Section       `json:"sections"`
	Position StudentPosition `json:"student_position"`
}
