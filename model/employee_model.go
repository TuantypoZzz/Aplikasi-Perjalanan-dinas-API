package model

type EmployeeModel struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	NIP        string `json:"nip"`
	Pangkat    string `json:"pangkat"`
	Department string `json:"Department"`
	Bidang     string `json:"bidang"`
}

type CreateEmployeeModel struct {
	Name        string `json:"name" validate:"required,max=32"`
	NIP         string `json:"nip" validate:"required,max=32"`
	HandleId    string `json:"handle_id" validate:"required,max=64"`
	DepartmenId string `json:"department_id" validate:"required,max=64"`
	Field       string `json:"field" validate:"required,max=32"`
	Section     string `json:"section" validate:"required,max=32"`
	UnitWork    string `json:"unit_work" validate:"required,max=32"`
	Gender      string `json:"gender" validate:"required,max=32"`
	BirthPlace  string `json:"birth_place" validate:"required,max=32"`
	BirthDate   string `json:"birth_date"`
	Phone       string `json:"phone" validate:"required,max=13"`
	Email       string `json:"email" validate:"required,max=32"`
	Address     string `json:"address" validate:"required,max=32"`
	Photo       string `json:"photo" `
}

type UpdateEmployee struct {
	Name         string `json:"name" validate:"required,max=32"`
	NIP          string `json:"nip" validate:"required,max=32"`
	IdPangkat    string `json:"id_pangkat" validate:"required,max=32"`
	IdDepartment string `json:"id_Department" validate:"required,max=32"`
	Bidang       string `json:"bidang" validate:"required,max=32"`
	Seksi        string `json:"seksi" validate:"required,max=32"`
	UnitKerja    string `json:"unit_kerja" validate:"required,max=32"`
	JenisKelamin string `json:"jenis_kelamin" validate:"required,max=32"`
	TempatLahir  string `json:"tempat_lahir" validate:"required,max=32"`
	TanggalLahir string `json:"tanggal_lahir"`
	NoTlpn       string `json:"no_tlpn" validate:"required,max=13"`
	Email        string `json:"email" validate:"required,max=32"`
	Alamat       string `json:"alamat" validate:"required,max=32"`
}
