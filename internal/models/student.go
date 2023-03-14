package models

import "time"

// Student struct holds student entity
type Student struct {
	Id         int       `json:"id"`
	FullName   string    `json:"fullName"`
	Year       int       `json:"year"`
	Email      string    `json:"email"`
	Department string    `json:"department"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}
