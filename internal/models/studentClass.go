package models

import "time"

type studentClass struct {
	Id        int
	StudentId int
	ClassId   int
	CreatedAt time.Time
	UpdatedAt time.Time
}
