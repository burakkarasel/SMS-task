package db

import (
	"context"
	"database/sql"
	"github.com/burakkarasel/SMS-task/internal/models"
	"github.com/burakkarasel/SMS-task/internal/util"
	"time"
)

type MockStore struct {
}

func (mock *MockStore) CreateClass(_ context.Context, arg models.CreateClassParams) (models.Class, error) {
	if arg.Name == "YOLO" {
		return models.Class{}, sql.ErrConnDone
	}
	return models.Class{Id: 5, Name: arg.Name, Professor: arg.Professor, CreatedAt: time.Now(), UpdatedAt: time.Now()}, nil
}

func (mock *MockStore) ListClasses(_ context.Context, arg models.ListClassesParams) ([]models.Class, error) {
	if arg.Limit == 17 {
		return nil, sql.ErrConnDone
	}
	var classes []models.Class
	for i := 0; i < 5; i++ {
		classes = append(classes, util.RandomClass())
	}
	return classes, nil
}

func (mock *MockStore) UpdateClass(_ context.Context, arg models.UpdateClassParams) (models.Class, error) {
	if arg.Name == "YOLO" {
		return models.Class{}, sql.ErrConnDone
	}
	if arg.Name == "OH NO" {
		return models.Class{}, sql.ErrNoRows
	}
	classToUpdate := util.RandomClass()
	classToUpdate.Name = arg.Name
	classToUpdate.Professor = arg.Professor
	classToUpdate.UpdatedAt = time.Now()
	return classToUpdate, nil
}

func (mock *MockStore) GetClass(_ context.Context, arg models.GetOneClassParam) (models.Class, error) {
	if arg.Id == 17 {
		return models.Class{}, sql.ErrConnDone
	}
	if arg.Id == 18 {
		return models.Class{}, sql.ErrNoRows
	}
	return util.RandomClass(), nil
}

func (mock *MockStore) DeleteClass(_ context.Context, arg models.DeleteOneClassParam) error {
	if arg.Id == 17 {
		return sql.ErrConnDone
	}
	if arg.Id == 18 {
		return sql.ErrNoRows
	}
	return nil
}

func (mock *MockStore) CreateStudent(_ context.Context, arg models.CreateStudentParams) (models.Student, error) {
	if arg.FullName == "YOLO" {
		return models.Student{}, sql.ErrConnDone
	}
	return models.Student{
		Id:         util.RandomInt(1, 1000),
		FullName:   arg.FullName,
		Year:       arg.Year,
		Email:      arg.Email,
		Department: arg.Department,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now()}, nil
}

func (mock *MockStore) ListStudents(_ context.Context, arg models.ListStudentsParams) ([]models.Student, error) {
	if arg.Limit == 17 {
		return nil, sql.ErrConnDone
	}
	var students []models.Student
	for i := 0; i < 5; i++ {
		students = append(students, util.RandomStudent())
	}
	return students, nil
}

func (mock *MockStore) GetStudent(_ context.Context, arg models.GetOneStudentParam) (models.Student, error) {
	if arg.Id == 17 {
		return models.Student{}, sql.ErrConnDone
	}

	if arg.Id == 18 {
		return models.Student{}, sql.ErrNoRows
	}
	return util.RandomStudent(), nil
}

func (mock *MockStore) UpdateStudent(_ context.Context, arg models.UpdateStudentParams) (models.Student, error) {
	if arg.FullName == "YOLO" {
		return models.Student{}, sql.ErrConnDone
	}
	if arg.FullName == "OH NO" {
		return models.Student{}, sql.ErrNoRows
	}
	studentToUpdate := util.RandomStudent()
	studentToUpdate.FullName = arg.FullName
	studentToUpdate.Year = arg.Year
	studentToUpdate.Department = arg.Department
	studentToUpdate.Email = arg.Email
	studentToUpdate.UpdatedAt = time.Now()
	return studentToUpdate, nil
}

func (mock *MockStore) DeleteStudent(_ context.Context, arg models.DeleteOneStudentParam) error {
	if arg.Id == 17 {
		return sql.ErrConnDone
	}
	if arg.Id == 18 {
		return sql.ErrNoRows
	}
	return nil
}

func (mock *MockStore) CreateStudentClass(_ context.Context, arg models.CreateStudentClassParams) (models.StudentClassResponse, error) {
	if arg.StudentId == 17 {
		return models.StudentClassResponse{}, sql.ErrConnDone
	}
	if arg.StudentId == 18 {
		return models.StudentClassResponse{}, sql.ErrNoRows
	}
	return util.RandomStudentClassResponse(), nil
}

func (mock *MockStore) GetStudentClass(_ context.Context, arg models.GetOneStudentClassParam) (models.StudentClassResponse, error) {
	if arg.StudentClassId == 17 {
		return models.StudentClassResponse{}, sql.ErrConnDone
	}
	if arg.StudentClassId == 18 {
		return models.StudentClassResponse{}, sql.ErrNoRows
	}
	return util.RandomStudentClassResponse(), nil
}

func (mock *MockStore) ListClassesOfStudent(_ context.Context, arg models.ListClassesOfStudentParams) (models.StudentsClassesResponse, error) {
	if arg.StudentId == 17 {
		return models.StudentsClassesResponse{}, sql.ErrConnDone
	}
	if arg.StudentId == 18 {
		return models.StudentsClassesResponse{}, sql.ErrNoRows
	}
	return util.RandomStudentClassesResponse(), nil
}

func (mock *MockStore) ListStudentsOfClass(_ context.Context, arg models.ListStudentsOfClassParams) (models.ClassStudentsResponse, error) {
	if arg.ClassId == 17 {
		return models.ClassStudentsResponse{}, sql.ErrConnDone
	}
	if arg.ClassId == 18 {
		return models.ClassStudentsResponse{}, sql.ErrNoRows
	}
	return util.RandomClassStudentsResponse(), nil
}

func (mock *MockStore) DeleteStudentClass(_ context.Context, arg models.DeleteOneStudentClassParam) error {
	if arg.StudentClassId == 17 {
		return sql.ErrConnDone
	}
	if arg.StudentClassId == 18 {
		return sql.ErrNoRows
	}
	return nil
}
