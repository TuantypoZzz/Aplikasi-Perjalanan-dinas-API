package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BussinessTravelReport struct {
	Id                     uuid.UUID `gorm:"primaryKey;column:bussiness_travel_report_id;type:varchar(36)"`
	NoSkMewaliki           string    `gorm:"column:no_sk_mewakili"`
	TglSkMewakili          string    `gorm:"column:tgl_sk_mewakili"`
	Keperluan              string    `gorm:"column:keperluan"`
	KotaAsal               string    `gorm:"column:kota_asal"`
	KotaTujuan             string    `gorm:"column:kota_tujuan"`
	TglBerangkat           string    `gorm:"column:tgl_berangkat"`
	TglKembali             string    `gorm:"column:tgl_kembali"`
	NamaKegiatan           string    `gorm:"column:nama_kegiatan"`
	MataAnggaran           string    `gorm:"column:mata_anggaran"`
	MaksudPerdin           string    `gorm:"column:maksud_perdin"`
	TujuanPerdin           string    `gorm:"column:tujuan_perdin"`
	WaktuPelaksanaanPerdin string    `gorm:"column:waktu_pelaksanaan_perdin"`
	SaranTindakanPerdin    string    `gorm:"column:saran_tindakan_perdin"`
	Instansi               string    `gorm:"column:instansi"`
	TglLapor               string    `gorm:"column:tgl_lapor"`
	Foto                   string    `gorm:"column:foto"`
	KetDok                 string    `gorm:"column:ket_dok"`
	TotalBiaya             string    `gorm:"column:total_biaya"`
	DeletedAt              gorm.DeletedAt
	CreatedAt              time.Time
	UpdatedAt              time.Time

	Employees       []Employee `gorm:"many2many:perdin_employee;"`
	SppdId          uuid.UUID  `gorm:"foreignKey;column:sppd_id;type:varchar(36)"`
	TransportId     uuid.UUID  `gorm:"foreignKey;column:transport_id;type:varchar(36)"`
	AccommodationId uuid.UUID  `gorm:"foreignKey;column:accommodation_id;type:varchar(36)"`
}
