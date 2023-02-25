package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Spt struct {
	Id        uuid.UUID `gorm:"primaryKey;column:spt_id;type:varchar(36)"`
	NoSpt     string    `gorm:"column:no_spt"`
	Keperluan string    `gorm:"column:keperluan"`
	TglSPT    string    `gorm:"column:tgl_spt"`
	DeletedAt gorm.DeletedAt
	CreatedAt time.Time
	UpdatedAt time.Time

	BussinessTravelReportId uuid.UUID
	BussinessTravelReports  BussinessTravelReport `gorm:"foreignKey:BussinessTravelReportId"`
}

func (Spt) TableName() string {
	return "spt"
}
