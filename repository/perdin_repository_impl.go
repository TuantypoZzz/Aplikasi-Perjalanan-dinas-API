package repository

import (
	"context"
	"golang-todo-app/entity"
	"golang-todo-app/exception"

	"gorm.io/gorm"
)

type perdinRepositoryImpl struct {
	*gorm.DB
}

func NewPerdinRepositoryImpl(DB *gorm.DB) PerdinRepository {
	return &perdinRepositoryImpl{DB: DB}
}

func (repository *perdinRepositoryImpl) InsertPegawaiPerdin(ctx context.Context, perdin entity.BussinessTravelReport, employees []entity.Employee) {
	db := repository.DB.WithContext(ctx).Omit("Employees.*").Create(&perdin)
	err := db.Model(&perdin).Association("Employees").Append(employees)

	exception.PanicLogging(err)
}

func (repository *perdinRepositoryImpl) SetPerdinEmployeeCommand(ctx context.Context, perdinId string, employeeId string, employeePerdin entity.PerdinEmployee) {
	err := repository.DB.WithContext(ctx).Model(&employeePerdin).Where("bussiness_travel_report_id = ?", perdinId).Where("employee_id = ?", employeeId).Update("command_flag", 1).Error
	exception.PanicLogging(err)
}

// func (repository *perdinRepositoryImpl) GetPejabatTTD(ctx context.Context, activityId string) entity.Activity {
// 	var TtdPPTK entity.Activity
// 	result := repository.DB.WithContext(ctx).Where("activity_id = ? AND pejabat_pptk_id =? ", activityId).First(&TtdPPTK)
// 	if result.RowsAffected == 0 {
// 		panic(exception.NotFoundError{
// 			Message: "Data Not Found",
// 		})
// 	}
// 	return TtdPPTK
// }

// func (repository *perdinRepositoryImpl) FindPerdinEmployeeById(ctx context.Context, id string) entity.BussinessTravelReport {
// 	var perdin entity.BussinessTravelReport
// 	// var employeePerdin []entity.PerdinEmployee
// 	result := repository.WithContext(ctx).Preload("Employees").Where("bussiness_travel_report_id = ?", id).First(&perdin)
// 	if result.RowsAffected == 0 {
// 		panic(exception.NotFoundError{
// 			Message: "Data not found",
// 		})
// 	}
// 	return perdin
// }

func (r *perdinRepositoryImpl) FindAllEmployeesPerdin(ctx context.Context) []entity.PerdinEmployee {
	var emplyPerdin []entity.PerdinEmployee
	r.DB.WithContext(ctx).Find(&emplyPerdin)
	return emplyPerdin
}

func (repository *perdinRepositoryImpl) FindAllEmployeePerdinById(ctx context.Context, id string) []entity.PerdinEmployee {
	var employeePerdin []entity.PerdinEmployee
	result := repository.WithContext(ctx).Where("bussiness_travel_report_id =?", id).Preload("Employee").Find(&employeePerdin)
	if result.RowsAffected == 0 {
		panic(exception.NotFoundError{
			Message: "Data not found",
		})
	}
	return employeePerdin
}

func (r *perdinRepositoryImpl) SetSPTPerdin(ctx context.Context, perdinId string, perdinSptReport entity.BussinessTravelReport) {
	err := r.DB.WithContext(ctx).Model(&perdinSptReport).Where("bussiness_travel_report_id = ?", perdinId).Updates(&perdinSptReport).Error
	exception.PanicLogging(err)
}

func (repository *perdinRepositoryImpl) FindPerdinById(ctx context.Context, id string) entity.BussinessTravelReport {
	var idperdin entity.BussinessTravelReport
	result := repository.DB.WithContext(ctx).Where("bussiness_travel_report_id = ?", id).Find(&idperdin)
	if result.RowsAffected == 0 {
		panic(exception.NotFoundError{
			Message: "Perdin Not Found",
		})
	}
	return idperdin
}

func (repository *perdinRepositoryImpl) FindEmployeePerdinByName(ctx context.Context, id string) entity.PerdinEmployee {
	var emplyPerdin entity.PerdinEmployee
	result := repository.DB.WithContext(ctx).
		Joins("JOIN employees ON employees.employee_id = perdin_employee.employee_id").
		Where("perdin_employee.perdin_employee_id = ?", id).
		Preload("Employee").
		Find(&emplyPerdin)

	if result.RowsAffected == 0 {
		panic(exception.NotFoundError{
			Message: "Employee Perdin Not Found",
		})
	}

	return emplyPerdin
}

func (repository *perdinRepositoryImpl) InsertPerdinReport(ctx context.Context, perdinReport entity.BussinessTravelReport) entity.BussinessTravelReport {
	err := repository.DB.WithContext(ctx).Create(&perdinReport).Error
	exception.PanicLogging(err)
	return perdinReport
}
