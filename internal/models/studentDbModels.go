package models

// CreateStudentParams holds the fields that required to execute the DB function
type CreateStudentParams struct {
	FullName   string
	Year       int
	Department string
	Email      string
}

// UpdateStudentParams holds the fields that required to execute the DB function
type UpdateStudentParams struct {
	Id         int
	FullName   string
	Year       int
	Department string
	Email      string
}

// ListStudentsParams holds the fields that required to execute the DB function
type ListStudentsParams struct {
	Offset int
	Limit  int
}

// GetOneStudentParam holds the fields that required to execute the DB function
type GetOneStudentParam struct {
	Id int
}

// DeleteOneStudentParam holds the fields that required to execute the DB function
type DeleteOneStudentParam struct {
	Id int
}
