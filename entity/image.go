package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Image struct {
	Id        uuid.UUID `gorm:"primaryKey;column:image_id;type:varchar(36)"`
	NameFile  string    `gorm:"colomn:name_file"`
	TipeFile  string    `gorm:"colomn:tipe_file"`
	Ukuran    string    `gorm:"colomn:ukuran_file"`
	CreatedBy string    `gorm:"column:created_by"`
	UpdatedBy string    `gorm:"column:updated_by"`

	DeletedAt gorm.DeletedAt
	CreatedAt time.Time
	UpdatedAt time.Time

	EmployeeId              uuid.UUID
	BussinessTravelReportID uuid.UUID

	BussinessTravelReport BussinessTravelReport `gorm:"foreignkey:BussinessTravelReportID"`
	Employee              Employee              `gorm:"foreignkey:EmployeeId1"`
}
