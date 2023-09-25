package model

type SppdModel struct {
	Id                           string   `json:"sppd_id"`
	NomorSurat                   string   `json:"nomor_surat"`
	PejabatTtdSppd               string   `json:"pejabat_ttd_sppd"`
	PegawaiPemimpin              string   `json:"nama_pegawai_pemimpin"`
	PangkatPegawaiPemimpin       string   `json:"pangkat_pegawai_pemimpin"`
	JabatanPegawaiPemimpin       string   `json:"jabatan_pegawai_pemimpin"`
	TingkatPerdinPegawaiPemimpin string   `json:"tingkat_perdin_pegawai_pemimpin"`
	MataAnggaran                 string   `json:"mata_anggaran"`
	Kegiatan                     string   `json:"kegiatan"`
	Instansi                     string   `json:"instansi"`
	JenisAngkutan                string   `json:"jenis_angkutan"`
	KotaAsal                     string   `json:"kota_asal"`
	KotaTujuan                   string   `json:"kota_tujuan"`
	Lama                         string   `json:"lama"`
	TglBerangkat                 string   `json:"tgl_berangkat"`
	TglKembali                   string   `json:"tgl_kembali"`
	PegawaiPengikut              []string `json:"nama_pegawai_pengikut"`
	TanggalSppd                  string   `json:"tgl_sppd"`
}

type CreateSppdModel struct {
	NomorSurat     string `json:"nomor_surat"`
	PejabatTtdSppd string `json:"pejabat_ttd_sppd"`
	TanggalSppd    string `json:"tgl_sppd"`
	PerdinId       string `json:"perdin_id"`
}
