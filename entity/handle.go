package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Handle struct {
	Id           uuid.UUID `gorm:"primaryKey;column:handle_id;type:varchar(36)"`
	NamaPangkat  string    `gorm:"column:nama_pangkat;type:varchar(36)"`
	NamaGolongan string    `gorm:"column:nama_golongan;type:varchar(36)"`
	Tingkat      string    `gorm:"column:tingkat;type:varchar(1)"`
	CreatedBy    string    `gorm:"column:created_by"`
	UpdatedBy    string    `gorm:"column:updated_by"`
	DeletedAt    gorm.DeletedAt
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
