package models

// CreateClassApiParams holds the data of required fields in the request
type CreateClassApiParams struct {
	Name      string `json:"name" binding:"required,min=2"`
	Professor string `json:"professor" binding:"required,min=4"`
}

// UpdateClassApiBodyParams holds the data of required fields in the request
type UpdateClassApiBodyParams struct {
	Name      string `json:"name" binding:"required,min=2"`
	Professor string `json:"professor" binding:"required,min=4"`
}

// UpdateClassApiUriParam holds the ID of the student
type UpdateClassApiUriParam struct {
	Id int `uri:"id" binding:"required,min=1"`
}

// ListClassesApiParams holds the data of required fields in the request
type ListClassesApiParams struct {
	PageId    int `form:"pageId" binding:"required,min=1"`
	PageLimit int `form:"pageLimit" binding:"required,min=5,max=20"`
}

// GetOneClassApiParam holds the ID of the student
type GetOneClassApiParam struct {
	Id int `uri:"id" binding:"required,min=1"`
}

// DeleteClassApiParam holds the ID of the student
type DeleteClassApiParam struct {
	Id int `uri:"id" binding:"required,min=1"`
}
