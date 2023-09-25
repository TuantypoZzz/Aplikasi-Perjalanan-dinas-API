package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Accommodation struct {
	Id                  uuid.UUID `gorm:"primaryKey;column:acomodation_id;type:varchar(36)"`
	BiayaTransportDarat float64   `gorm:"column:biaya_transport_darat"`
	NamaHotel           string    `gorm:"column:nama_hotel"`
	TipeKamar           string    `gorm:"column:tipe_kamar"`
	NoKamar             string    `gorm:"column:no_kamar"`
	CheckInDate         time.Time `gorm:"column:checkin_date"`
	CheckOutDate        time.Time `gorm:"column:checkout_date"`
	BiayaPermalam       float64   `gorm:"column:biaya_permalam"`
	LamaMenginap        string    `gorm:"column:lama_menginap"`
	DeletedAt           gorm.DeletedAt
	CreatedAt           time.Time
	UpdatedAt           time.Time

	PerdinEmployeeId uuid.UUID
}
