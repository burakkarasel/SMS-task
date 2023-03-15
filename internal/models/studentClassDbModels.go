package models

// CreateStudentClassParams holds the fields that required to execute the DB function
type CreateStudentClassParams struct {
	StudentId int
	ClassId   int
}

// ListStudentsOfClassParams holds the fields that required to execute the DB function
type ListStudentsOfClassParams struct {
	ClassId int
	Limit   int
	Offset  int
}

// ListClassesOfStudentParams holds the fields that required to execute the DB function
type ListClassesOfStudentParams struct {
	StudentId int
	Limit     int
	Offset    int
}

// GetOneStudentClassParam holds the fields that required to execute the DB function
type GetOneStudentClassParam struct {
	StudentClassId int
}

// DeleteOneStudentClassParam holds the fields that required to execute the DB function
type DeleteOneStudentClassParam struct {
	StudentClassId int
}
