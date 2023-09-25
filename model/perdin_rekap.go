package model

type RekapPerdin struct {
	IdBussinessTravelReport string   `json:"id"`
	NoSppd                  string   `json:"no_spt"`
	Kegiatan                string   `json:"kegiatan"`
	TglLapor                string   `json:"tgl_lapor"`
	TglBerangkat            string   `json:"tgl_berangkat"`
	TglKembali              string   `json:"tgl_kembali"`
	PegawaiPemimpin         string   `json:"pegawai_pemimpin"`
	PegawaiPengikut         []string `json:"pegawai_pengikut"`
	KotaTujuan              string   `json:"tujuan"`
	LamaPerdin              int16    `json:"lama_perdin"`
}
