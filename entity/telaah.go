package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Telaah struct {
	Id        uuid.UUID  `gorm:"primaryKey;column:telaah_id;type:varchar(36)"`
	NoTelaah  string     `gorm:"column:no_telaah"`
	TglTelaah *time.Time `gorm:"column:tgl_telaah"`
	DeletedAt gorm.DeletedAt
	CreatedAt time.Time
	UpdatedAt time.Time

	ActivityId              uuid.UUID
	BussinessTravelReportId uuid.UUID

	Activity               Activity              `gorm:"foreignKey:ActivityId"`
	BussinessTravelReports BussinessTravelReport `gorm:"foreignKey:BussinessTravelReportId"`
}

func (Telaah) TableName() string {
	return "telaah"
}
