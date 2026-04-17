package model

type Section struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Order       int    `json:"order"`
	LevelsCount int    `json:"levels_count"`
}

type StudentPosition struct {
	SectionId  int `json:"section_id"`
	LevelOrder int `json:"level_order"`
}

type LevelsMap struct {
	Sections []Section       `json:"sections"`
	Position StudentPosition `json:"student_position"`
}
