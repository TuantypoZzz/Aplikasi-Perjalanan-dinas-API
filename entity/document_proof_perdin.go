package entity

import (
	"time"

	"github.com/google/uuid"
)

type DokProofPerdin struct {
	Id               uuid.UUID `gorm:"primaryKey;column:dokumen_perdin_id;type:varchar(36)"`
	PerdinEmployeeId uuid.UUID
	Files            string `gorm:"column:files"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
