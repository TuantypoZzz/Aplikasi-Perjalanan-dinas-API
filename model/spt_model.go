package model

type SptModel struct {
	Id              string   `json:"spt_id"`
	NomorSurat      string   `json:"nomor_surat"`
	Keperluan       string   `json:"keperluan"`
	NamaKegiatan    string   `json:"nama_kegiatan"`
	KotaAsal        string   `json:"kota_asal"`
	KotaTujuan      string   `json:"kota_tujuan"`
	PegawaiPemimpin string   `json:"nama_pegawai_pemimpin"`
	PegawaiPengikut []string `json:"nama_pegawai_pengikut"`
	TglBerangkat    string   `json:"tgl_berangkat"`
	TglKembali      string   `json:"tgl_kembali"`
	TglLapor        string   `json:"tgl_lapor"`
	TglSpt          string   `json:"tgl_spt"`
	PejabatTtdSpt   string   `json:"pejabat_ttd_spt"`
	NoSkMewaliki    string   `json:"no_sk_mewakili"`
	TglSkMewakili   string   `json:"tgl_sk_mewakili"`
}

type CreateSptModel struct {
	NomorSurat    string `json:"nomor_surat"`
	Keperluan     string `json:"keperluan"`
	KotaAsal      string `json:"kota_asal"`
	KotaTujuan    string `json:"kota_tujuan"`
	TglLapor      string `json:"tgl_lapor"`
	TglSpt        string `json:"tgl_spt"`
	PejabatTtdSpt string `json:"pejabat_ttd_spt"`
	NoSkMewaliki  string `json:"no_sk_mewakili"`
	TglSkMewakili string `json:"tgl_sk_mewakili"`
	PerdinId      string `json:"perdin_id"`
	TelaahId      string `json:"telaah_id"`
}
