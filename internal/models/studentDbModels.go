package models

type CreateStudentParams struct {
	FullName   string
	Year       int
	Department string
	Email      string
}

type UpdateStudentParams struct {
	Id         int
	FullName   string
	Year       int
	Department string
	Email      string
}

type ListStudentsParams struct {
	Offset int
	Limit  int
}

type GetOneStudentParam struct {
	Id int
}

type DeleteOneStudentParam struct {
	Id int
}
