package model

type EmployeeModel struct {
	Id           string `json:"id"`
	Name         string `json:"name"`
	NIP          string `json:"nip"`
	Pangkat      string `json:"pangkat"`
	Jabatan      string `json:"jabatan"`
	Bidang       string `json:"Bidang"`
	Seksi        string `json:"seksi"`
	UnitKerja    string `json:"unit_kerja"`
	JenisKelamin string `json:"jenis_kelamin"`
	TempatLahir  string `json:"tempat_lahir"`
	TanggalLahir string `json:"tanggal_lahir"`
	NoTlpn       string `json:"nomor_tlpn"`
	Email        string `json:"email"`
	Alamat       string `json:"alamat"`
	Foto         string `json:"foto"`
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
	Id           string `json:"id"`
	Name         string `json:"nama"`
	NIP          string `json:"nip"`
	IdPangkat    string `json:"id_pangkat"`
	IdJabatan    string `json:"id_jabatan"`
	Bidang       string `json:"bidang"`
	Seksi        string `json:"seksi"`
	UnitKerja    string `json:"unit_kerja"`
	JenisKelamin string `json:"jenis_kelamin"`
	TempatLahir  string `json:"tempat_lahir"`
	TanggalLahir string `json:"tanggal_lahir"`
	NoTlpn       string `json:"nomor_tlpn"`
	Email        string `json:"email"`
	Alamat       string `json:"alamat"`
	Foto         string `json:"foto"`
}
