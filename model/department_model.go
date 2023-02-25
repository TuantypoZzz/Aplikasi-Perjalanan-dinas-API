package model

type DepartmentModel struct {
	Id             string `json:"id"`
	DepartmentName string `json:"department_name"`
}

type CreateDepartment struct {
	DepartmentName string `json:"department_name" validate:"required,max=32"`
}

type UpdateDepartment struct {
	DepartmentName string `json:"department_name" validate:"required,max=32"`
}
