package db

import (
	"context"
	"github.com/burakkarasel/SMS-task/internal/models"
)

// CreateStudent creates a new student in DB
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

// ListStudents list students with given parameters
func (q *Queries) ListStudents(ctx context.Context, arg models.ListStudentsParams) ([]models.Student, error) {
	query := `
		SELECT id, full_name, year, department, email, created_at, updated_at
		FROM students
		LIMIT $1
		OFFSET $2
	`

	// first we execute the query
	rows, err := q.db.QueryContext(ctx, query, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	// we defer closing rows
	defer rows.Close()

	// slice for holding the students data
	students := []models.Student{}

	// looping through the rows till there is none
	for rows.Next() {
		var tempStudent models.Student
		// scan the current row and append it to array after checking errors
		if err := rows.Scan(
			&tempStudent.Id,
			&tempStudent.FullName,
			&tempStudent.Year,
			&tempStudent.Department,
			&tempStudent.Email,
			&tempStudent.CreatedAt,
			&tempStudent.UpdatedAt,
		); err != nil {
			return nil, err
		}
		students = append(students, tempStudent)
	}
	// closing the rows anyway
	if err := rows.Close(); err != nil {
		return nil, err
	}
	// checking if there is any error
	if err := rows.Err(); err != nil {
		return nil, err
	}
	// finally returning the slice and nil as error
	return students, nil
}

// GetStudent gets a single student with given ID
func (q *Queries) GetStudent(ctx context.Context, arg models.GetOneStudentParam) (models.Student, error) {
	query := `
		SELECT id, full_name, year, department, email, created_at, updated_at
		FROM students
		WHERE id = $1
		`

	var student models.Student

	row := q.db.QueryRowContext(ctx, query, arg.Id)
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

// UpdateStudent updates the student by its id and given values
func (q *Queries) UpdateStudent(ctx context.Context, arg models.UpdateStudentParams) (models.Student, error) {
	stmt := `
		UPDATE students
		SET full_name = $1,
    	year = $2,
    	department = $3,
    	email = $4,
    	updated_at = NOW()
		WHERE id = $5
		RETURNING id, full_name, year, department, email, created_at, updated_at;
		`

	var student models.Student

	row := q.db.QueryRowContext(ctx, stmt, arg.FullName, arg.Year, arg.Department, arg.Email, arg.Id)
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

// DeleteStudent deletes student by given ID
func (q *Queries) DeleteStudent(ctx context.Context, arg models.DeleteOneStudentParam) error {
	stmt := `
			DELETE FROM students
			WHERE id = $1
			RETURNING id
		`

	var id int

	row := q.db.QueryRowContext(ctx, stmt, arg.Id)
	err := row.Scan(&id)
	return err
}
