package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Telaah struct {
	Id        uuid.UUID `gorm:"primaryKey;column:telaah_id;type:varchar(36)"`
	NoTelaah  string    `gorm:"column:no_telaah"`
	TglTelaah string    `gorm:"column:tgl_telaah"`
	DeletedAt gorm.DeletedAt
	CreatedAt time.Time
	UpdatedAt time.Time

	PejabatPPTKId           uuid.UUID
	BussinessTravelReportId uuid.UUID

	Employee               Employee              `gorm:"foreignKey:PejabatPPTKId"`
	BussinessTravelReports BussinessTravelReport `gorm:"foreignKey:BussinessTravelReportId"`
}

func (Telaah) TableName() string {
	return "telaah"
}
