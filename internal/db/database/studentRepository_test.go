package db

import (
	"context"
	"github.com/burakkarasel/SMS-task/internal/models"
	"github.com/burakkarasel/SMS-task/internal/util"
	"log"
	"testing"
)

// createRandomStudent creates a random student in the DB
func createRandomStudent(t *testing.T) models.Student {
	arg := models.CreateStudentParams{
		FullName:   util.RandomName(),
		Year:       util.RandomInt(1, 4),
		Department: util.RandomName(),
		Email:      util.RandomEmail(),
	}

	s, err := testQueries.CreateStudent(context.Background(), arg)
	if err != nil {
		t.Fatal(err)
	}
	if s.Id <= 0 {
		t.Fatal(err)
	}
	if s.Email != arg.Email {
		t.Fatal(err)
	}
	if s.Department != arg.Department {
		t.Fatal(err)
	}
	if s.FullName != arg.FullName {
		t.Fatal(err)
	}
	if s.Year != arg.Year {
		t.Fatal(err)
	}
	return s
}

// TestCreateStudent test CreateStudent function
func TestCreateStudent(t *testing.T) {
	_ = createRandomStudent(t)
}

// TestGetStudent tests GetStudent function
func TestGetStudent(t *testing.T) {
	s := createRandomStudent(t)

	s2, err := testQueries.GetStudent(context.Background(), models.GetOneStudentParam{Id: s.Id})

	if err != nil {
		t.Fatal(err)
	}
	if s2.CreatedAt != s.CreatedAt {
		t.Fatal(err)
	}
	if s2.Id != s.Id {
		t.Fatal(err)
	}
	if s2.UpdatedAt != s.UpdatedAt {
		t.Fatal(err)
	}
	if s2.Year != s.Year {
		t.Fatal(err)
	}
	if s2.Email != s.Email {
		t.Fatal(err)
	}
	if s2.Department != s.Department {
		t.Fatal(err)
	}
	if s2.FullName != s.FullName {
		t.Fatal(err)
	}
}

// TestListStudents tests ListStudents Function
func TestListStudents(t *testing.T) {
	s := []models.Student{}
	for i := 0; i < 3; i++ {
		newS := createRandomStudent(t)
		s = append(s, newS)
	}
	got, err := testQueries.ListStudents(context.Background(), models.ListStudentsParams{Limit: 3, Offset: 0})
	if err != nil {
		t.Fatal(err)
	}
	if len(got) != len(s) {
		t.Fatal(err)
	}
}

// TestUpdateStudent tests UpdateStudent function
func TestUpdateStudent(t *testing.T) {
	s := createRandomStudent(t)

	arg := models.UpdateStudentParams{
		Id:         s.Id,
		FullName:   util.RandomName(),
		Department: util.RandomName(),
		Email:      util.RandomEmail(),
		Year:       util.RandomInt(1, 4),
	}

	got, err := testQueries.UpdateStudent(context.Background(), arg)

	if err != nil {
		log.Fatal(err)
	}
	if got.Id != s.Id {
		log.Fatal(err)
	}
	if got.FullName != arg.FullName {
		log.Fatal(err)
	}
	if got.Department != arg.Department {
		log.Fatal(err)
	}
	if got.Year != arg.Year {
		log.Fatal(err)
	}
	if got.Email != arg.Email {
		log.Fatal(err)
	}
}

// TestDeleteStudent tests DeleteStudent function
func TestDeleteStudent(t *testing.T) {
	s := createRandomStudent(t)

	err := testQueries.DeleteStudent(context.Background(), models.DeleteOneStudentParam{Id: s.Id})
	if err != nil {
		log.Fatal(err)
	}

	err = testQueries.DeleteStudent(context.Background(), models.DeleteOneStudentParam{Id: s.Id})
	if err == nil {
		log.Fatal(err)
	}
}
