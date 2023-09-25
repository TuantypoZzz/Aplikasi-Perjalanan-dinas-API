package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Employee struct {
	Id         uuid.UUID `gorm:"primaryKey;column:employee_id;type:varchar(36)"`
	Name       string    `gorm:"column:name"`
	NIP        string    `gorm:"column:nip"`
	Field      string    `gorm:"column:field"`
	Section    string    `gorm:"column:section"`
	UnitWork   string    `gorm:"column:unit_work"`
	Gender     string    `gorm:"type:enum('Laki-Laki','Perempuan');column:gender"`
	BirthPlace string    `gorm:"column:birth_place"`
	BirthDate  time.Time `gorm:"column:birth_date"`
	Phone      string    `gorm:"column:phone"`
	Email      string    `gorm:"column:email"`
	Address    string    `gorm:"type:text;column:address"`
	Photo      string    `gorm:"type:text;column:photo"`
	CreatedBy  string    `gorm:"column:created_by"`
	UpdatedBy  string    `gorm:"column:updated_by"`

	DeletedAt gorm.DeletedAt
	CreatedAt time.Time
	UpdatedAt time.Time

	DepartmentId uuid.UUID
	HandleId     uuid.UUID

	Department Department `gorm:"foreignkey:DepartmentId"`
	Handle     Handle     `gorm:"foreignkey:HandleId"`

	BussinessTravelReports []BussinessTravelReport `gorm:"many2many:perdin_employee;"`
}
