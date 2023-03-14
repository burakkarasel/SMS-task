package models

import "time"

// Class holds class entity
type Class struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Professor string    `json:"professor"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
