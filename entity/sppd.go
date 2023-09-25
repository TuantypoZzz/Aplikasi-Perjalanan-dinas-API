package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Sppd struct {
	Id        uuid.UUID  `gorm:"primaryKey;column:sppd_id;type:varchar(36)"`
	NoSppd    string     `gorm:"column:no_sppd"`
	TglSppd   *time.Time `gorm:"column:tgl_sppd"`
	DeletedAt gorm.DeletedAt
	CreatedAt time.Time
	UpdatedAt time.Time

	PejabatSPTId uuid.UUID
	Employee     Employee `gorm:"foreignKey:PejabatSPTId"`

	BussinessTravelReportId uuid.UUID
	BussinessTravelReports  BussinessTravelReport `gorm:"foreignKey:BussinessTravelReportId"`
}

func (Sppd) TableName() string {
	return "sppd"
}
