package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PerdinEmployee struct {
	Id                      uuid.UUID `gorm:"primaryKey;column:perdin_employee_id;type:varchar(36)"`
	BussinessTravelReportId uuid.UUID `gorm:"primaryKey;column:bussiness_travel_report_id;type:varchar(36)"`
	EmployeeId              uuid.UUID `gorm:"primaryKey;column:employee_id;type:varchar(36)"`
	CommandFlag             int
	Debt                    int
	DailyMoney              int
	LodgingExpenses         int
	TotalTransportationCost int
	TotalHotelCost          int
	TotalDailyCost          int
	TotalPerdinCost         int

	LumpsumId       uuid.UUID `gorm:"default:null"`
	Lumpsums        Lumpsum   `gorm:"foreignkey:LumpsumId"`
	OtherCostDetail []OtherCostDetail
	Transports      []Transport
	Accommodations  []Accommodation
	DokProofPerdin  []DokProofPerdin

	Employee Employee
}

func (PerdinEmployee) TableName() string {
	return "perdin_employee"
}

func (perdinEmployee *PerdinEmployee) BeforeCreate(tx *gorm.DB) (err error) {
	perdinEmployee.Id = uuid.New()
	return
}
