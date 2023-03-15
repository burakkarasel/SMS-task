package db

import (
	"context"
	"github.com/burakkarasel/SMS-task/internal/models"
	"github.com/burakkarasel/SMS-task/internal/util"
	"testing"
)

// createRandomClass creates a random Class in the DB
func createRandomClass(t *testing.T) models.Class {
	arg := models.CreateClassParams{
		Professor: util.RandomName(),
		Name:      util.RandomName(),
	}

	c, err := testQueries.CreateClass(context.Background(), arg)
	if err != nil {
		t.Fatal(err)
	}

	if arg.Professor != c.Professor {
		t.Fatal(err)
	}

	if arg.Name != c.Name {
		t.Fatal(err)
	}

	if c.Id <= 0 {
		t.Fatal("Id must be greater than 0")
	}

	return c
}

// TestCreateClass tests CreateClass function
func TestCreateClass(t *testing.T) {
	_ = createRandomClass(t)
}

// TestGetClass tests GetClass function
func TestGetClass(t *testing.T) {
	c := createRandomClass(t)

	arg := models.GetOneClassParam{Id: c.Id}

	got, err := testQueries.GetClass(context.Background(), arg)

	if err != nil {
		t.Fatal(err)
	}

	if got.Id != c.Id {
		t.Fatal(err)
	}

	if got.Professor != c.Professor {
		t.Fatal(err)
	}

	if got.Name != c.Name {
		t.Fatal(err)
	}

	if got.CreatedAt != c.CreatedAt {
		t.Fatal(err)
	}

	if got.UpdatedAt != c.UpdatedAt {
		t.Fatal(err)
	}
}

// TestListClasses tests ListClasses function
func TestListClasses(t *testing.T) {
	for i := 0; i < 3; i++ {
		_ = createRandomClass(t)
	}
	arg := models.ListClassesParams{
		Offset: 0,
		Limit:  3,
	}

	got, err := testQueries.ListClasses(context.Background(), arg)

	if err != nil {
		t.Fatal(err)
	}

	if len(got) != 3 {
		t.Fatal(err)
	}
}

// TestUpdateClass tests UpdateClass function
func TestUpdateClass(t *testing.T) {
	c := createRandomClass(t)

	arg := models.UpdateClassParams{
		Id:        c.Id,
		Name:      util.RandomName(),
		Professor: util.RandomName(),
	}

	got, err := testQueries.UpdateClass(context.Background(), arg)

	if err != nil {
		t.Fatal(err)
	}

	if got.Id != arg.Id {
		t.Fatal(err)
	}

	if got.Name != arg.Name {
		t.Fatal(err)
	}

	if got.Professor != arg.Professor {
		t.Fatal(err)
	}
}

// TestDeleteClass tests DeleteClass function
func TestDeleteClass(t *testing.T) {
	c := createRandomClass(t)

	arg := models.DeleteOneClassParam{Id: c.Id}

	err := testQueries.DeleteClass(context.Background(), arg)
	if err != nil {
		t.Fatal(err)
	}

	err = testQueries.DeleteClass(context.Background(), arg)
	if err == nil {
		t.Fatal(err)
	}
}
