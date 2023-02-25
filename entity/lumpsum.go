package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Lumpsum struct {
	Id              uuid.UUID `gorm:"primaryKey;column:lumpsum_id;type:varchar(36)"`
	Lingkup         string    `gorm:"column:lingkup"`
	BiayaHarian     string    `gorm:"column:biaya_harian"`
	BiayaPenginapan string    `gorm:"column:biaya_penginapan"`
	DeletedAt       gorm.DeletedAt
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

func (Lumpsum) TableName() string {
	return "lumpsum"
}
