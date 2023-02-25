package model

type HandleModel struct {
	Id            string `json:"id"`
	NamaPangkat   string `json:"nama_pangkat"`
	NamaGolongan  string `json:"nama_golongan"`
	Tingkat       string `json:"tingkat"`
	EmployeeModel string `json:"employee_model"`
}

type CreateHandle struct {
	NamaPangkat  string `json:"nama_pangkat" validate:"required,max=32"`
	NamaGolongan string `json:"nama_golongan" validate:"required,max=32"`
	Tingkat      string `json:"tingkat" validate:"required,max=1"`
}

type UpdateHandle struct {
	NamaPangkat  string `json:"nama_pangkat" validate:"required,max=32"`
	NamaGolongan string `json:"nama_golongan" validate:"required,max=32"`
	Tingkat      string `json:"tingkat" validate:"required,max=1"`
}
