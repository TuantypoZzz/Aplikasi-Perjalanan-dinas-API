package entity

import "github.com/google/uuid"

type OtherCostDetail struct {
	Id        uuid.UUID `gorm:"primaryKey;column:other_cost_detail_id;type:varchar(36)"`
	Keperluan string    `gorm:"column:keperluan"`
	Harga     string    `gorm:"column:harga"`

	PerdinEmployeeId uuid.UUID
}
