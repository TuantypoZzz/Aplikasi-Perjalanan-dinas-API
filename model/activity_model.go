package model

type ActivityModel struct {
	Id           string `json:"id"`
	ActivityName string `json:"nama_kegiatan"`
	BudgetLine   string `json:"mata_anggaran"`
	EmployeeId   string `json:"nama_pptk"`
	Instances    string `json:"instansi"`
}

type CreateActivityModel struct {
	ActivityName string `json:"nama_kegiatan"`
	BudgetLine   string `json:"mata_anggaran"`
	EmployeeId   string `json:"nama_pptk"`
	Instances    string `json:"instansi"`
}

type UpdateActivityModel struct {
	ActivityName string `json:"nama_kegiatan"`
	BudgetLine   string `json:"mata_anggaran"`
	EmployeeId   string `json:"nama_pptk"`
	Instances    string `json:"instansi"`
}
