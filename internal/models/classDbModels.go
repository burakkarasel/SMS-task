package models

// CreateClassParams holds the fields that required to execute the DB function
type CreateClassParams struct {
	Name      string
	Professor string
}

// UpdateClassParams holds the fields that required to execute the DB function
type UpdateClassParams struct {
	Id        int
	Name      string
	Professor string
}

// ListClassesParams holds the fields that required to execute the DB function
type ListClassesParams struct {
	Offset int
	Limit  int
}

// GetOneClassParam holds the fields that required to execute the DB function
type GetOneClassParam struct {
	Id int
}

// DeleteOneClassParam holds the fields that required to execute the DB function
type DeleteOneClassParam struct {
	Id int
}
