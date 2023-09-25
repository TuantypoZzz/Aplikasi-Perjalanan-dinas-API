package model

type PerdinReportModel struct {
	Id                     string `json:"id"`
	MaksudPerdin           string `json:"maksud_perdin"`
	TujuanPerdin           string `json:"tujuan_perdin"`
	WaktuPelaksanaanPerdin string `json:"waktu_pelaksanaan_perdin"`
	HasilYangdiCapai       string `json:"hasil_yang_dicapai"`
	SaranDanTindakLanjut   string `json:"saran_dan_tindak_lanjut"`
}

type CreatePerdinReportModel struct {
	MaksudPerdin           string `json:"maksud_perdin"`
	TujuanPerdin           string `json:"tujuan_perdin"`
	WaktuPelaksanaanPerdin string `json:"waktu_pelaksanaan_perdin"`
	HasilYangdiCapai       string `json:"hasil_yang_dicapai"`
	SaranDanTindakLanjut   string `json:"saran_dan_tindak_lanjut"`
	Foto                   string `json:"foto"`
	PerdinId               string `json:"perdin_id"`
}
