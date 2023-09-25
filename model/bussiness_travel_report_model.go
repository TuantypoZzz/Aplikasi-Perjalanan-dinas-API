package model

import "time"

type PerdinModel struct {
	Id                     string    `json:"id"`
	NoSkMewaliki           string    `json:"no_sk_mewakili"`
	TglSkMewakili          string    `json:"tgl_sk_mewakili"`
	Keperluan              string    `json:"keperluan"`
	KotaAsal               string    `json:"kota_asal"`
	KotaTujuan             string    `json:"kota_tujuan"`
	TglBerangkat           string    `json:"tgl_berangkat"`
	TglKembali             string    `json:"tgl_kembali"`
	MaksudPerdin           string    `json:"maksud_perdin"`
	TujuanPerdin           string    `json:"tujuan_perdin"`
	WaktuPelaksanaanPerdin string    `json:"wakti_pelaksanaan_perdin"`
	SaranTindakanPerdin    string    `json:"saran_tindakan_perdin"`
	TglLapor               time.Time `json:"tgl_lapor"`
	Foto                   string    `json:"foto"`
	KetDok                 string    `json:"keterangan_dokumentasi"`
	TotalBiaya             string    `json:"total_biaya"`
	Employees              []string  `json:"nama_employee"`
}

type CreatePerdinModel struct {
	Id                     string `json:"id"`
	NoSkMewaliki           string `json:"no_sk_mewakili"`
	TglSkMewakili          string `json:"tgl_sk_mewakili"`
	Keperluan              string `json:"keperluan"`
	KotaAsal               string `json:"kota_asal"`
	KotaTujuan             string `json:"kota_tujuan"`
	TglBerangkat           string `json:"tgl_berangkat"`
	TglKembali             string `json:"tgl_kembali"`
	MaksudPerdin           string `json:"maksud_perdin"`
	TujuanPerdin           string `json:"tujuan_perdin"`
	WaktuPelaksanaanPerdin string `json:"wakti_pelaksanaan_perdin"`
	SaranTindakanPerdin    string `json:"saran_tindakan_perdin"`
	TglLapor               string `json:"tgl_lapor"`
	Foto                   string `json:"foto"`
	KetDok                 string `json:"keterangan_dokumentasi"`
	TotalBiaya             string `json:"total_biaya"`
}
