package model

type CreateDetailPerdin struct {
	NamaPemimpin    string   `json:"nama_pemimpin_perdin"`
	TiketDarat      int      `json:"tiket_darat"`
	TotalTransport  string   `json:"total_transport"`
	UangHarian      int      `json:"uang_harian"`
	BiayaPenginapan int      `json:"biaya_penginapan"`
	AngkutanPergi   Angkutan `json:"angkutan_pergi"`
	AngkutanPulang  Angkutan `json:"angkutan_pulang"`
	NamaHotel       string   `json:"nama_hotel"`
	TipeKamar       string   `json:"tipe_kamar"`
	NomorKamar      string   `json:"nomor_kamar"`
	TglCekin        string   `json:"tgl_cekin"`
	TglCekout       string   `json:"tgl_cekout"`
	PerdinId        string   `json:"perdin_id"`
}

type Angkutan struct {
	HargaTiket           int    `json:"harga_tiket"`
	JenisAngkutan        string `json:"jenis_angkutan"`
	NomorTiket           string `json:"nomor_tiket"`
	NomorPenerbangan     string `json:"nomor_penerbangan"`
	KelasPenerbangan     string `json:"kelas_penerbangan"`
	KodeBooking          string `json:"kode_booking"`
	TanggalPenerbangan   string `json:"tgl_penerbangan"`
	TaksiPengeluaranRill int    `json:"taksi_pengeluaran_rill"`
}
