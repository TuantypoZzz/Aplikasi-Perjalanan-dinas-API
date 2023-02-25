package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Department struct {
	Id             uuid.UUID `gorm:"primaryKey;column:department_id;type:varchar(36)"`
	DepartmentName string    `gorm:"column:department_name"`
	CreatedBy      string    `gorm:"column:created_by"`
	UpdatedBy      string    `gorm:"column:updated_by"`
	DeletedAt      gorm.DeletedAt
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
