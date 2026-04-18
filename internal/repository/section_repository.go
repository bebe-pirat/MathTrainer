package repository

import (
	"MathTrainer/internal"
	"MathTrainer/internal/model"
	"context"
	"database/sql"
)

type SectionRepository interface {
	GetSectionsByClass(ctx context.Context, class int, currentSection int) ([]model.Section, error)
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

func (r *SectionRepositoryStruct) GetSectionsByClass(ctx context.Context, class int, currentSection int) ([]model.Section, error) {
	query := `
		SELECT id, name, section_order, COUNT(*) as types_count, CASE WHEN sections.id <= $1 THEN TRUE ELSE FALSE END AS unlocked
		FROM sections 
		JOIN section_equation_types ON section_id = sections.id
		WHERE class = $2
		GROUP BY sections.id, name, section_order
		ORDER BY section_order;
	`

	var sections []model.Section
	rows, err := r.db.QueryContext(ctx, query, currentSection, class)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var section model.Section
		var typesCount int

		err = rows.Scan(
			&section.Id,
			&section.Name,
			&section.Order,
			&typesCount,
			&section.IsUnlocked,
		)
		if err != nil {
			return nil, err
		}

		section.LevelsCount = internal.LevelsInSectionCoef * typesCount

		sections = append(sections, section)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return sections, nil
}

func (r *SectionRepositoryStruct) GetStudentPosition(ctx context.Context, studentId int) (*model.StudentPosition, error) {
	query := `
		SELECT section_id, MAX(level_order)
		FROM student_progress_level
		WHERE user_id = $1
		GROUP BY section_id
		ORDER BY section_id DESC
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
