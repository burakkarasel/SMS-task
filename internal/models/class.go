package models

import "time"

type Class struct {
	Id        int
	Name      string
	Professor string
	CreatedAt time.Time
	UpdatedAt time.Time
}
