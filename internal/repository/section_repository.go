package repository

import (
	"MathTrainer/internal"
	"MathTrainer/internal/model"
	"context"
	"database/sql"
	"log/slog"
)

type SectionRepository interface {
	GetSectionsByClass(ctx context.Context, class int, currentSection int) ([]model.Section, error)
	GetStudentPosition(ctx context.Context, studentId int) (*model.StudentPosition, error)
	GetFirstStudentSection(ctx context.Context, studentId int) (int, error)

	CreateSection(ctx context.Context, section model.Section) error
	UpdateSection(ctx context.Context, section model.Section) error
	DeleteSection(ctx context.Context, sectionId int) error
	GetSections(ctx context.Context, class int) ([]model.Section, error)
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
		SELECT section_id, MAX(level_order) + 1
		FROM student_progress_level
		WHERE user_id = $1
		GROUP BY section_id
		ORDER BY section_id DESC
		LIMIT 1;
	`

	var pos model.StudentPosition
	err := r.db.QueryRowContext(ctx, query, studentId).Scan(&pos.SectionId, &pos.LevelOrder)
	if err == sql.ErrNoRows {
		sectionId, err := r.GetFirstStudentSection(ctx, studentId)
		if err != nil {
			return nil, err
		}
		return &model.StudentPosition{SectionId: sectionId, LevelOrder: 1}, nil
	}
	if err != nil {
		return nil, err
	}

	return &pos, nil
}

func (r *SectionRepositoryStruct) GetFirstStudentSection(ctx context.Context, studentId int) (int, error) {
	query := `
		SELECT sections.id
		FROM sections
		JOIN classes ON sections.class = classes.grade
		JOIN users ON users.class_id = classes.id
		WHERE section_order = 1 AND users.id = $1
		LIMIT 1;
	`

	var sectionId int
	err := r.db.QueryRowContext(ctx, query, studentId).Scan(&sectionId)
	if err != nil {
		return 0, err
	}

	return sectionId, nil
}

func (r *SectionRepositoryStruct) CreateSection(ctx context.Context, section model.Section) error {
	query := `
		INSERT INTO sections(name, class, section_order) 
		VALUES ($1, $2, $3);
	`

	res, err := r.db.ExecContext(ctx, query, section.Name, section.Class, section.Order)
	if err != nil {
		return err
	}

	if rows, err := res.RowsAffected(); err != nil || rows == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (r *SectionRepositoryStruct) UpdateSection(ctx context.Context, section model.Section) error {
	query := `
		UPDATE sections SET name = $1, 
							class = $2, 
							section_order = $3
		WHERE id = $4;
	`

	res, err := r.db.ExecContext(ctx, query, section.Name, section.Class, section.Order, section.Id)
	if err != nil {
		return err
	}

	if rows, err := res.RowsAffected(); err != nil || rows == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (r *SectionRepositoryStruct) DeleteSection(ctx context.Context, sectionId int) error {
	query := `
		DELETE FROM sections
		WHERE id = $1;
	`

	res, err := r.db.ExecContext(ctx, query, sectionId)
	if err != nil {
		slog.Error("failed to delete", "error", err)
		return err
	}

	if rows, err := res.RowsAffected(); err != nil || rows == 0 {
		slog.Error("failed to delete", "error", err)
		return sql.ErrNoRows
	}

	return nil
}

func (r *SectionRepositoryStruct) GetSections(ctx context.Context, class int) ([]model.Section, error) {
	query := `
		SELECT id, name, class, section_order
		FROM sections
	`
	parameters := make([]any, 0)
	if class >= 1 && class <= 4 {
		query += `WHERE class = $1;`
		parameters = append(parameters, class)
	}

	rows, err := r.db.QueryContext(ctx, query, parameters...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	sections := make([]model.Section, 0)
	for rows.Next() {
		var section model.Section

		err := rows.Scan(&section.Id, &section.Name, &section.Class, &section.Order)
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
