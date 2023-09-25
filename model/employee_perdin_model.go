package model

type EmployeePerdinModel struct {
	IdPerdinEmployee        string   `json:"id_perdin_employee"`
	IdBussinessTravelReport string   `json:"id_bussiness_travel_report"`
	IdEmployee              []string `json:"id_employee"`
	CommandFlag             int      `json:"command_flag"`
	Debt                    int      `json:"debt"`
	TotalTransportationCost int      `json:"total_transport_cost"`
	TotalHotelCost          int      `json:"total_hotel_cost"`
	TotalDailyCost          int      `json:"total_daily_cost"`
	TotalPerdinCost         int      `json:"total_perdin_cost"`
	IdLumpsum               string   `json:"id_lumpsum"`
	Nama                    string   `json:"nama"`
}

type CreateEmployeePerdinModel struct {
	IdBussinessTravelReport string   `json:"id_bussiness_travel_report"`
	IdEmployee              []string `json:"id_employee"`
	CommandFlag             int      `json:"command_flag"`
	Debt                    int      `json:"debt"`
	TotalTransportationCost int      `json:"total_transport_cost"`
	TotalHotelCost          int      `json:"total_hotel_cost"`
	TotalDailyCost          int      `json:"total_daily_cost"`
	TotalPerdinCost         int      `json:"total_perdin_cost"`
	IdLumpsum               string   `json:"id_lumpsum"`
}
