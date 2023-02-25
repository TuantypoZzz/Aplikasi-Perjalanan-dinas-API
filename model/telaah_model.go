package model

type TelaahModel struct {
	Id             string `json:"id"`
	IdPejabatPPTK  string `json:"id_pejabat_pptk"`
	PerdinReportId string `json:"perdin_report_id"`
	NoTelaah       string `json:"no_telaah"`
	TglTelaah      string `json:"tgl_telaah"`
}
