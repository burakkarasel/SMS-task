package models

// CreateStudentClassApiParams holds the data of required fields in the request
type CreateStudentClassApiParams struct {
	StudentId int `json:"studentId" binding:"required,min=1"`
	ClassId   int `json:"classId" binding:"required,min=1"`
}

// ListStudentClassesApiParams holds the data of required fields in the request
type ListStudentClassesApiParams struct {
	PageLimit int `form:"pageLimit" binding:"required,min=5,max=20"`
	PageId    int `form:"pageId" binding:"required,min=1"`
	ClassId   int `form:"classId" binding:"min=0"`
	StudentId int `form:"studentId" binding:"min=0"`
}

// GetOneStudentClassApiParams holds the data of required fields in the request
type GetOneStudentClassApiParams struct {
	StudentClassId int `uri:"studentClassId" binding:"required,min=1"`
}

// DeleteOneStudentClassApiParams holds the data of required fields in the request
type DeleteOneStudentClassApiParams struct {
	StudentClassId int `uri:"studentClassId" binding:"required,min=1"`
}
