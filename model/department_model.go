package model

type DepartmentModel struct {
	Id             string `json:"id"`
	DepartmentName string `json:"nama_jabatan"`
}

type CreateDepartment struct {
	DepartmentName string `json:"nama_jabatan"`
}

type UpdateDepartment struct {
	DepartmentName string `json:"nama_jabatan"`
}
