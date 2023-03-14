package models

// CreateStudentApiParams holds the data of required fields in the request
type CreateStudentApiParams struct {
	FullName   string `json:"fullName" binding:"required,min=4"`
	Year       int    `json:"year" binding:"required,min=1,max=4"`
	Department string `json:"department" binding:"required,min=2"`
	Email      string `json:"email" binding:"required,email"`
}

// ListStudentsApiParams holds the data of required fields in the request
type ListStudentsApiParams struct {
	PageId    int `form:"pageId" binding:"required,min=1"`
	PageLimit int `form:"pageLimit" binding:"required,min=5,max=20"`
}

// UpdateStudentApiBodyParams holds the data of required fields in the request
type UpdateStudentApiBodyParams struct {
	FullName   string `json:"fullName" binding:"required,min=4"`
	Year       int    `json:"year" binding:"required,min=1,max=4"`
	Department string `json:"department" binding:"required,min=2"`
	Email      string `json:"email" binding:"required,email"`
}

// UpdateStudentApiUriParam holds the ID of the student
type UpdateStudentApiUriParam struct {
	Id int `uri:"id" binding:"required,min=1"`
}

// DeleteStudentApiParam holds the ID of the student
type DeleteStudentApiParam struct {
	Id int `uri:"id" binding:"required,min=1"`
}
