package repository

import (
	"context"
	"golang-todo-app/entity"
)

type PerdinRepository interface {
	InsertPegawaiPerdin(ctx context.Context, perdin entity.BussinessTravelReport, employees []entity.Employee)
	SetPerdinEmployeeCommand(ctx context.Context, perdinId string, employeeId string, employeePerdin entity.PerdinEmployee)
	//GetPejabatTTD(ctx context.Context, activityId string) entity.Activity
	// FindPerdinEmployeeById(ctx context.Context, id string) entity.BussinessTravelReport
	FindAllEmployeePerdinById(ctx context.Context, id string) []entity.PerdinEmployee
	SetSPTPerdin(ctx context.Context, perdinId string, perdinSptReport entity.BussinessTravelReport)
	FindPerdinById(ctx context.Context, id string) entity.BussinessTravelReport
	FindAllEmployeesPerdin(ctx context.Context) []entity.PerdinEmployee
	InsertPerdinReport(ctx context.Context, PerdinReport entity.BussinessTravelReport) entity.BussinessTravelReport
	FindEmployeePerdinByName(ctx context.Context, id string) entity.PerdinEmployee
}
