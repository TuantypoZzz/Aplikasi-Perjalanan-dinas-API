package repository

import (
	"context"
	"golang-todo-app/entity"
	"golang-todo-app/exception"

	"gorm.io/gorm"
)

type employeePerdinRepositoryImpl struct {
	*gorm.DB
}

func NewEmployeePerdinRepositoryImpl(DB *gorm.DB) EmployeePerdinRepository {
	return &employeePerdinRepositoryImpl{DB: DB}
}

func (r *employeePerdinRepositoryImpl) SetDetailPerdin(ctx context.Context, perdinEmpyId string, detailEmpyPerdin entity.PerdinEmployee) {
	err := r.DB.WithContext(ctx).Where("perdin_employee_id = ?", perdinEmpyId).Updates(&detailEmpyPerdin).Error
	exception.PanicLogging(err)
}
