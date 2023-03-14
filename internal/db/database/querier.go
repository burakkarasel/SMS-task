package database

import (
	"context"
	"github.com/burakkarasel/SMS-task/internal/models"
)

type Querier interface {
	CreateStudent(ctx context.Context, newStudent models.CreateStudentParams) (models.Student, error)
	ListStudents(ctx context.Context, arg models.ListStudentsParams) ([]models.Student, error)
	UpdateStudent(ctx context.Context, arg models.UpdateStudentParams) (models.Student, error)
	DeleteStudent(ctx context.Context, arg models.DeleteOneStudentParam) error
}

var _ Querier = (*Queries)(nil)
