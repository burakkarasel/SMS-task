package util

import (
	"fmt"
	"github.com/burakkarasel/SMS-task/internal/models"
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

// init runs as program starts and enables true randomness
func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomInt generates a random int between given arguments
func RandomInt(min, max int) int {
	return min + int(rand.Int63n(int64(max-min+1)))
}

// RandomString generates a random string with given char count
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// RandomName returns a 6 random digits string
func RandomName() string {
	return RandomString(4)
}

// RandomEmail creates a random email
func RandomEmail() string {
	return fmt.Sprintf("%s@email.com", RandomString(6))
}

// RandomClass creates a Class instance with random values
func RandomClass() models.Class {
	return models.Class{
		Id:        RandomInt(1, 1000),
		Name:      RandomName(),
		Professor: RandomName(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

// RandomStudent creates a Student instance with random values
func RandomStudent() models.Student {
	return models.Student{
		Id:         RandomInt(1, 1000),
		FullName:   RandomName(),
		Year:       RandomInt(1, 4),
		Department: RandomName(),
		Email:      RandomEmail(),
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
}

// RandomStudentClassesResponse create a StudentClassesResponse instance with random values
func RandomStudentClassesResponse() models.StudentsClassesResponse {
	studentClasses := models.StudentsClassesResponse{
		Student: RandomStudent(),
	}
	for i := 0; i < 5; i++ {
		studentClasses.Classes = append(studentClasses.Classes, RandomClass())
	}
	return studentClasses
}

// RandomStudentClassResponse creates a StudentClassResponse instance with random values
func RandomStudentClassResponse() models.StudentClassResponse {
	return models.StudentClassResponse{
		Id:        RandomInt(1, 1000),
		Student:   RandomStudent(),
		Class:     RandomClass(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

// RandomClassStudentsResponse creates a ClassStudentResponse instance with random values
func RandomClassStudentsResponse() models.ClassStudentsResponse {
	classStudentsResponse := models.ClassStudentsResponse{
		Class: RandomClass(),
	}
	for i := 0; i < 5; i++ {
		classStudentsResponse.Students = append(classStudentsResponse.Students, RandomStudent())
	}
	return classStudentsResponse
}
