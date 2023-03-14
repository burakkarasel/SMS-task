package models

import "time"

// studentClass struct holds student class mapping entity
type studentClass struct {
	Id        int       `json:"id"`
	StudentId int       `json:"studentId"`
	ClassId   int       `json:"classId"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
