package entity

import "github.com/google/uuid"

type Activity struct {
	Id           uuid.UUID `gorm:"primaryKey;column:activity_id;type:varchar(36)"`
	NamaKegiatan string    `gorm:"column:nama_kegiatan"`
	MataAnggaran string    `gorm:"column:mata_anggaran"`
	Instansi     string    `gorm:"column:instansi"`
	CreatedBy    string    `gorm:"column:created_by"`
	UpdatedBy    string    `gorm:"column:updated_by"`

	PejabatPPTKId uuid.UUID
	Employee      Employee `gorm:"foreignKey:PejabatPPTKId"`
}

type ActivityWithEmployeeName struct {
	Activity
	PejabatName string `gorm:"column:name"`
}
