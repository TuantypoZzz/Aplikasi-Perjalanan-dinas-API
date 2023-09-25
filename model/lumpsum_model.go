package model

type Lumpsum struct {
	Id              string `json:"id"`
	Lingkup         string `json:"lingkup"`
	BiayaHarian     string `json:"biaya_harian"`
	BiayaPenginapan string `json:"biaya_penginapan"`
}

type CreateLumpsum struct {
	Lingkup         string `json:"lingkup" validate:"required,max=32"`
	BiayaHarian     string `json:"biaya_harian" validate:"required,max=32"`
	BiayaPenginapan string `json:"biaya_penginapan" validate:"required,max=32"`
}

type UpdateLumpsum struct {
	Lingkup         string `json:"lingkup" validate:"required,max=32"`
	BiayaHarian     string `json:"biaya_harian" validate:"required,max=32"`
	BiayaPenginapan string `json:"biaya_penginapan" validate:"required,max=32"`
}
