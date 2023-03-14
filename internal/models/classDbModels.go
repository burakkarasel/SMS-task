package models

type CreateClassParams struct {
	Name      string
	Professor string
}

type UpdateClassParams struct {
	Id        int
	Name      string
	Professor string
}

type ListClassesParams struct {
	Offset int
	Limit  int
}

type GetOneClassParam struct {
	Id int
}

type DeleteOneClassParam struct {
	Id int
}
