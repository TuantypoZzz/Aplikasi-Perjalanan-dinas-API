package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Transport struct {
	Id               uuid.UUID `gorm:"primaryKey;column:transport_id;type:varchar(36)"`
	NamaAngkutan     string    `gorm:"column:nama_angkutan"`
	FlagTransport    string    `gorm:"column:flag_transport"`
	NoTiket          string    `gorm:"column:No_tiket"`
	NoPenerbangan    string    `gorm:"column:No_penerbangan"`
	KodeBooking      string    `gorm:"column:Kode_booking"`
	TglPenerbangan   string    `gorm:"column:tgl_penerbangan"`
	KelasPenerbangan string    `gorm:"column:kelp_penerbangan"`
	DeletedAt        gorm.DeletedAt
	CreatedAt        time.Time
	UpdatedAt        time.Time

	PerdinEmployeeId uuid.UUID
}
