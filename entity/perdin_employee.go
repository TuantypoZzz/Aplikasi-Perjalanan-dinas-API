package entity

import "github.com/google/uuid"

type PerdinEmployee struct {
	Id                      uuid.UUID `gorm:"primaryKey;column:perdin_employee_id;type:varchar(36)"`
	BussinessTravelReportId uuid.UUID `gorm:"primaryKey;column:bussiness_travel_report_id;type:varchar(36)"`
	EmployeeId              uuid.UUID `gorm:"primaryKey;column:employee_id;type:varchar(36)"`
	CommandFlag             int
	Debt                    int
	TotalTransportationCost int
	TotalHotelCost          int
	TotalDailyCost          int
	TotalPerdinCost         int

	LumpsumId       uuid.UUID
	Lumpsums        Lumpsum `gorm:"foreignkey:LumpsumId"`
	OtherCostDetail []OtherCostDetail
	Transports      []Transport
	Accommodations  []Accommodation
}

func (PerdinEmployee) TableName() string {
	return "perdin_employee"
}
