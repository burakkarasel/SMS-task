package database

import (
	"context"
	"github.com/burakkarasel/SMS-task/internal/models"
)

func (q *Queries) CreateStudent(ctx context.Context, newStudent models.CreateStudentParams) (models.Student, error) {
	stmt := `
	INSERT INTO students (full_name, year, department, email)
	VALUES ($1, $2, $3, $4)
	RETURNING id, full_name, year, department, email, created_at, updated_at
	`

	row := q.db.QueryRowContext(ctx, stmt, newStudent.FullName, newStudent.Year, newStudent.Department, newStudent.Email)
	var student models.Student
	err := row.Scan(
		&student.Id,
		&student.FullName,
		&student.Year,
		&student.Department,
		&student.Email,
		&student.CreatedAt,
		&student.UpdatedAt,
	)
	return student, err
}
