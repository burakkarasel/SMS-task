package database

import (
	"context"
	"github.com/burakkarasel/SMS-task/internal/models"
)

// CreateClass creates a new class in the DB
func (q *Queries) CreateClass(ctx context.Context, newClass models.CreateClassParams) (models.Class, error) {
	stmt := `
		INSERT INTO classes(name, professor)
		VALUES($1, $2)
		RETURNING id, name, professor, created_at, updated_at
		`

	var class models.Class
	row := q.db.QueryRowContext(ctx, stmt, newClass.Name, newClass.Professor)
	err := row.Scan(
		&class.Id,
		&class.Name,
		&class.Professor,
		&class.CreatedAt,
		&class.UpdatedAt,
	)
	return class, err
}

// ListClasses list the classes with pagination
func (q *Queries) ListClasses(ctx context.Context, arg models.ListClassesParams) ([]models.Class, error) {
	query := `
		SELECT id, name, professor, created_at, updated_at
		FROM classes
		LIMIT $1
		OFFSET $2
		`

	classes := []models.Class{}
	rows, err := q.db.QueryContext(ctx, query, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var tempClass models.Class
		if err := rows.Scan(
			&tempClass.Id,
			&tempClass.Name,
			&tempClass.Professor,
			&tempClass.CreatedAt,
			&tempClass.UpdatedAt,
		); err != nil {
			return nil, err
		}
		classes = append(classes, tempClass)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return classes, nil
}

// UpdateClass updates a class in DB with given values
func (q *Queries) UpdateClass(ctx context.Context, arg models.UpdateClassParams) (models.Class, error) {
	stmt := `
		UPDATE classes
		SET name = $1, professor = $2, updated_at = NOW()
		WHERE id = $3
		RETURNING id, name, professor, created_at, updated_at
		`

	var class models.Class

	row := q.db.QueryRowContext(ctx, stmt, arg.Name, arg.Professor, arg.Id)
	err := row.Scan(
		&class.Id,
		&class.Name,
		&class.Professor,
		&class.CreatedAt,
		&class.UpdatedAt,
	)
	return class, err
}

// GetClass returns a single class from the DB with given ID
func (q *Queries) GetClass(ctx context.Context, arg models.GetOneClassParam) (models.Class, error) {
	query := `
		SELECT id, name, professor, created_at, updated_at
		FROM classes
		WHERE id = $1
		`
	var class models.Class
	row := q.db.QueryRowContext(ctx, query, arg.Id)
	err := row.Scan(
		&class.Id,
		&class.Name,
		&class.Professor,
		&class.CreatedAt,
		&class.UpdatedAt,
	)
	return class, err
}

// DeleteClass deletes a class from the DB with given ID
func (q *Queries) DeleteClass(ctx context.Context, arg models.DeleteOneClassParam) error {
	stmt := `
		DELETE FROM classes
		WHERE id = $1
		RETURNING id
	`
	var id int
	row := q.db.QueryRowContext(ctx, stmt, arg.Id)
	err := row.Scan(&id)
	return err
}
