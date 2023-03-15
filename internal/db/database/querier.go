package database

import (
	"context"
	"github.com/burakkarasel/SMS-task/internal/models"
)

type Querier interface {
	CreateStudent(ctx context.Context, newStudent models.CreateStudentParams) (models.Student, error)
	ListStudents(ctx context.Context, arg models.ListStudentsParams) ([]models.Student, error)
	GetStudent(ctx context.Context, arg models.GetOneStudentParam) (models.Student, error)
	UpdateStudent(ctx context.Context, arg models.UpdateStudentParams) (models.Student, error)
	DeleteStudent(ctx context.Context, arg models.DeleteOneStudentParam) error

	CreateClass(ctx context.Context, newClass models.CreateClassParams) (models.Class, error)
	ListClasses(ctx context.Context, arg models.ListClassesParams) ([]models.Class, error)
	GetClass(ctx context.Context, arg models.GetOneClassParam) (models.Class, error)
	UpdateClass(ctx context.Context, arg models.UpdateClassParams) (models.Class, error)
	DeleteClass(ctx context.Context, arg models.DeleteOneClassParam) error
}

var _ Querier = (*Queries)(nil)
