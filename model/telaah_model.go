package model

type TelaahModel struct {
	Id                      string   `json:"id"`
	Activity                string   `json:"nama_kegiatan"`
	NoTelaah                string   `json:"no_telaah"`
	TglTelaah               string   `json:"tgl_telaah"`
	TglBerangkat            string   `json:"tgl_berangkat"`
	TglKembali              string   `json:"tgl_kembali"`
	PegawaiPemimpin         string   `json:"nama_pegawai_pemimpin"`
	PegawaiPengikut         []string `json:"nama_pegawai_pengikut"`
	NamaPPTK                string   `json:"nama_pptk"`
	IdBussinessTravelReport string   `json:"bussiness_travel_report_id"`
	IdEmployeesPerdin       string   `json:"employees_perdin_id"`
}

type CreateTelaahModel struct {
	Activity        string   `json:"nama_kegiatan"`
	NoTelaah        string   `json:"no_telaah"`
	TglTelaah       string   `json:"tgl_telaah"`
	TglBerangkat    string   `json:"tgl_berangkat"`
	TglKembali      string   `json:"tgl_kembali"`
	PegawaiPemimpin string   `json:"nama_pegawai_pemimpin"`
	PegawaiPengikut []string `json:"nama_pegawai_pengikut"`
	NamaPPTK        string   `json:"nama_pptk"`
}
