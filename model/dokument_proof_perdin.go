package model

type DokProofPerdin struct {
	TiketPergi         string `json:"tiket_pergi"`
	BoardingPassPergi  string `json:"boading_pass_pergi"`
	TiketPulang        string `json:"tiket_pulang"`
	BoardingPassPulang string `json:"boarding_pass_pulang"`
	BillHotel          string `json:"bill_hotel"`
	KwitansiLainnya    string `json:"kwitansi_lainnya"`
	EmplyPerdinId      string `json:"employee_perdin_id"`
	PerdinId           string `json:"perdin_id"`
	Nama               string `json:"nama"`
}
