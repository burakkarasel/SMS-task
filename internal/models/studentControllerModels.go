package models

type CreateStudentApiParams struct {
	FullName   string `json:"fullName" binding:"required,min=4"`
	Year       int    `json:"year" binding:"required,min=1,max=4"`
	Department string `json:"department" binding:"required,min=2"`
	Email      string `json:"email" binding:"required,email"`
}
