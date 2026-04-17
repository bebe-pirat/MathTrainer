package repository

import (
	"MathTrainer/internal/model"
	"context"
	"database/sql"
)

type SectionRepository interface {
	GetSectionByClass(ctx context.Context, class int) ([]model.Section, error)
	GetStudentPosition(ctx context.Context, studentId int) (*model.StudentPosition, error)
}

type SectionRepositoryStruct struct {
	db *sql.DB
}

func NewSectionRepositoryStruct(db *sql.DB) *SectionRepositoryStruct {
	return &SectionRepositoryStruct{
		db: db,
	}
}

func (r *SectionRepositoryStruct) GetSectionByClass(ctx context.Context, class int) ([]model.Section, error) {
	query := `
		SELECT id, name, section_order, COUNT(*) as types_count
		FROM sections 
		JOIN section_equation_types ON section_id = sections.id
		WHERE class = $1;
	`

	var sections []model.Section
	rows, err := r.db.QueryContext(ctx, query, class)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var section model.Section

		err = rows.Scan(
			&section.Id,
			&section.Name,
			&section.Order,
			&section.LevelsCount,
		)
		if err != nil {
			return nil, err
		}

		sections = append(sections, section)
	}

	if rows.Err() != nil {
		return nil, err
	}

	return sections, nil
}

func (r *SectionRepositoryStruct) GetStudentPosition(ctx context.Context, studentId int) (*model.StudentPosition, error) {
	query := `
		SELECT section_id, level_order
		FROM student_progress_level
		WHERE user_id = $1
		ORDER BY section_id DESC, level_order DESC
		LIMIT 1;
	`

	var pos model.StudentPosition
	err := r.db.QueryRowContext(ctx, query, studentId).Scan(&pos.SectionId, &pos.LevelOrder)
	if err == sql.ErrNoRows {
		return &model.StudentPosition{SectionId: 1, LevelOrder: 0}, nil
	}
	if err != nil {
		return nil, err
	}

	return &pos, nil
}
