package database

import (
	"context"
	"github.com/burakkarasel/SMS-task/internal/models"
)

// CreateStudentClass creates a new student class mapping in the DB
func (q *Queries) CreateStudentClass(ctx context.Context, newStudentClass models.CreateStudentClassParams) (models.StudentClassResponse, error) {
	stmt := `
		INSERT INTO student_classes(student_id, class_id)
		VALUES ($1, $2)
		RETURNING id
		`
	var studentClass models.StudentClassResponse

	row := q.db.QueryRowContext(ctx, stmt, newStudentClass.StudentId, newStudentClass.ClassId)
	var id int

	err := row.Scan(&id)
	if err != nil {
		return studentClass, err
	}

	arg := models.GetOneStudentClassParam{StudentClassId: id}

	studentClass, err = q.GetStudentClass(ctx, arg)
	return studentClass, err
}

// GetStudentClass gets one studentClass form the DB with given ID
func (q *Queries) GetStudentClass(ctx context.Context, arg models.GetOneStudentClassParam) (models.StudentClassResponse, error) {
	query := `
		SELECT s.id, s.full_name, s.email, s.year, s.department, s.created_at, s.updated_at,
		       c.id, c.name, c.professor, c.created_at, c.updated_at,
		       sc.created_at, sc.updated_at
		FROM student_classes as sc
		INNER JOIN students as s on s.id = sc.student_id
		INNER JOIN classes as c on c.id = sc.class_id
		WHERE sc.id = $1
		`

	var studentClass models.StudentClassResponse

	row := q.db.QueryRowContext(ctx, query, arg.StudentClassId)
	err := row.Scan(
		&studentClass.Student.Id,
		&studentClass.Student.FullName,
		&studentClass.Student.Email,
		&studentClass.Student.Year,
		&studentClass.Student.Department,
		&studentClass.Student.CreatedAt,
		&studentClass.Student.UpdatedAt,
		&studentClass.Class.Id,
		&studentClass.Class.Name,
		&studentClass.Class.Professor,
		&studentClass.Class.CreatedAt,
		&studentClass.Class.UpdatedAt,
		&studentClass.CreatedAt,
		&studentClass.UpdatedAt,
	)
	studentClass.Id = arg.StudentClassId
	return studentClass, err
}

// ListClassesOfStudent lists given student's classes with pagination
func (q *Queries) ListClassesOfStudent(ctx context.Context, arg models.ListClassesOfStudentParams) (models.StudentsClassesResponse, error) {
	query := `
		SELECT s.id, s.full_name, s.email, s.year, s.department, s.created_at, s.updated_at,
		       c.id, c.name, c.professor, c.created_at, c.updated_at
		FROM student_classes as sc
		INNER JOIN students as s on s.id = sc.student_id
		INNER JOIN classes as c on c.id = sc.class_id
		WHERE s.id = $1
		LIMIT $2
		OFFSET $3
		`

	var studentClasses models.StudentsClassesResponse

	rows, err := q.db.QueryContext(ctx, query, arg.StudentId, arg.Limit, arg.Offset)
	if err != nil {
		return studentClasses, err
	}

	defer rows.Close()

	for rows.Next() {
		var tempClass models.Class
		err := rows.Scan(
			&studentClasses.Student.Id,
			&studentClasses.Student.FullName,
			&studentClasses.Student.Email,
			&studentClasses.Student.Year,
			&studentClasses.Student.Department,
			&studentClasses.Student.CreatedAt,
			&studentClasses.Student.UpdatedAt,
			&tempClass.Id,
			&tempClass.Name,
			&tempClass.Professor,
			&tempClass.CreatedAt,
			&tempClass.UpdatedAt,
		)
		if err != nil {
			return studentClasses, err
		}
		studentClasses.Classes = append(studentClasses.Classes, tempClass)
	}

	if err := rows.Close(); err != nil {
		return studentClasses, err
	}

	if err := rows.Err(); err != nil {
		return studentClasses, err
	}
	return studentClasses, nil
}

// ListStudentsOfClass lists given student's classes with pagination
func (q *Queries) ListStudentsOfClass(ctx context.Context, arg models.ListStudentsOfClassParams) (models.ClassStudentsResponse, error) {
	query := `
		SELECT s.id, s.full_name, s.email, s.year, s.department, s.created_at, s.updated_at,
		       c.id, c.name, c.professor, c.created_at, c.updated_at
		FROM student_classes as sc
		INNER JOIN students as s on s.id = sc.student_id
		INNER JOIN classes as c on c.id = sc.class_id
		WHERE c.id = $1
		LIMIT $2
		OFFSET $3
		`

	var classStudents models.ClassStudentsResponse

	rows, err := q.db.QueryContext(ctx, query, arg.ClassId, arg.Limit, arg.Offset)
	if err != nil {
		return classStudents, err
	}

	defer rows.Close()

	for rows.Next() {
		var tempStudent models.Student
		err := rows.Scan(
			&tempStudent.Id,
			&tempStudent.FullName,
			&tempStudent.Email,
			&tempStudent.Year,
			&tempStudent.Department,
			&tempStudent.CreatedAt,
			&tempStudent.UpdatedAt,
			&classStudents.Class.Id,
			&classStudents.Class.Name,
			&classStudents.Class.Professor,
			&classStudents.Class.CreatedAt,
			&classStudents.Class.UpdatedAt,
		)
		if err != nil {
			return classStudents, err
		}
		classStudents.Students = append(classStudents.Students, tempStudent)
	}

	if err := rows.Close(); err != nil {
		return classStudents, err
	}

	if err := rows.Err(); err != nil {
		return classStudents, err
	}
	return classStudents, nil
}

// DeleteStudentClass deletes a student class mapping from the DB
func (q *Queries) DeleteStudentClass(ctx context.Context, arg models.DeleteOneStudentClassParam) error {
	stmt := `
		DELETE FROM student_classes
		WHERE id = $1
		RETURNING id
	`
	var id int
	row := q.db.QueryRowContext(ctx, stmt, arg.StudentClassId)
	err := row.Scan(&id)
	return err
}
