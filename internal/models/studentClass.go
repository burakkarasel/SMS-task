package models

import "time"

// StudentClassResponse holds student and his/her class
type StudentClassResponse struct {
	Id        int       `json:"id"`
	Student   Student   `json:"student"`
	Class     Class     `json:"class"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// StudentsClassesResponse holds the classes of a student
type StudentsClassesResponse struct {
	Student Student `json:"student"`
	Classes []Class `json:"classes"`
}

// ClassStudentsResponse holds the students of a class
type ClassStudentsResponse struct {
	Class    Class     `json:"class"`
	Students []Student `json:"students"`
}
