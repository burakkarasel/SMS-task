package models

import "time"

type Student struct {
	Id         int
	FullName   string
	Year       int
	Department string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
