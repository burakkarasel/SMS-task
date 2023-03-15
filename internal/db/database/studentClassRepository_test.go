package db

import (
	"context"
	"github.com/burakkarasel/SMS-task/internal/models"
	"testing"
)

// createRandomStudentClass creates a random studentClass in DB
func createRandomStudentClass(t *testing.T) (models.Student, models.Class, models.StudentClassResponse) {
	s := createRandomStudent(t)
	c := createRandomClass(t)

	arg := models.CreateStudentClassParams{
		StudentId: s.Id,
		ClassId:   c.Id,
	}

	sc, err := testQueries.CreateStudentClass(context.Background(), arg)

	if err != nil {
		t.Fatal(err)
	}

	if sc.Id <= 0 {
		t.Fatal(err)
	}

	if sc.Student.Id != arg.StudentId {
		t.Fatal(err)
	}

	if sc.Class.Id != arg.ClassId {
		t.Fatal(err)
	}

	return s, c, sc
}

// TestCreateStudentClass test CreateStudentClass function
func TestCreateStudentClass(t *testing.T) {
	_, _, _ = createRandomStudentClass(t)
}

func TestGetStudentClass(t *testing.T) {
	_, _, sc := createRandomStudentClass(t)

	arg := models.GetOneStudentClassParam{StudentClassId: sc.Id}

	got, err := testQueries.GetStudentClass(context.Background(), arg)

	if err != nil {
		t.Fatal(err)
	}

	if got.Student != sc.Student {
		t.Fatal(err)
	}

	if got.Class != sc.Class {
		t.Fatal(err)
	}

	if got.Id != sc.Id {
		t.Fatal(err)
	}

	if got.CreatedAt != sc.CreatedAt {
		t.Fatal(err)
	}

	if got.UpdatedAt != sc.UpdatedAt {
		t.Fatal(err)
	}
}

// TestListClassesOfStudent tests ListClassesOfStudent function
func TestListClassesOfStudent(t *testing.T) {
	s, _, _ := createRandomStudentClass(t)

	arg := models.ListClassesOfStudentParams{StudentId: s.Id, Limit: 1, Offset: 0}
	got, err := testQueries.ListClassesOfStudent(context.Background(), arg)
	if err != nil {
		t.Fatal(err)
	}

	if len(got.Classes) != 1 {
		t.Fatal(err)
	}

	if got.Student.Id != s.Id {
		t.Fatal(err)
	}
}

// TestListStudentsOfClass tests ListStudentsOfClass function
func TestListStudentsOfClass(t *testing.T) {
	_, c, _ := createRandomStudentClass(t)

	arg := models.ListStudentsOfClassParams{ClassId: c.Id, Limit: 1, Offset: 0}
	got, err := testQueries.ListStudentsOfClass(context.Background(), arg)
	if err != nil {
		t.Fatal(err)
	}

	if len(got.Students) != 1 {
		t.Fatal(err)
	}

	if got.Class.Id != c.Id {
		t.Fatal(err)
	}
}

// TestDeleteStudentClass tests DeleteStudentClass function
func TestDeleteStudentClass(t *testing.T) {
	_, _, sc := createRandomStudentClass(t)

	arg := models.DeleteOneStudentClassParam{StudentClassId: sc.Id}
	err := testQueries.DeleteStudentClass(context.Background(), arg)
	if err != nil {
		t.Fatal(err)
	}

	err = testQueries.DeleteStudentClass(context.Background(), arg)
	if err == nil {
		t.Fatal(err)
	}

}
