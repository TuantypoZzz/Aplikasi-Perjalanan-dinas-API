package repository

import (
	"context"
	"golang-todo-app/entity"
	"golang-todo-app/exception"

	"gorm.io/gorm"
)

type employeeRepositoryImpl struct {
	*gorm.DB
}

func NewEmployeeRepositoryImpl(DB *gorm.DB) EmployeeRepository {
	return &employeeRepositoryImpl{DB: DB}
}

func (r *employeeRepositoryImpl) Insert(ctx context.Context, employee entity.Employee) entity.Employee {
	err := r.DB.WithContext(ctx).Create(&employee).Error

	exception.PanicLogging(err)
	return employee
}

func (r *employeeRepositoryImpl) Update(ctx context.Context, employee entity.Employee) entity.Employee {
	err := r.DB.WithContext(ctx).Where("employee_id = ?", employee.Id).Updates(&employee).Error
	exception.PanicLogging(err)
	return employee
}

func (r *employeeRepositoryImpl) Delete(ctx context.Context, employee entity.Employee) {
	r.DB.WithContext(ctx).Where("employee_id =?", employee.Id).Delete(&employee)
}

func (r *employeeRepositoryImpl) FindById(ctx context.Context, id string) entity.Employee {
	var employee entity.Employee
	result := r.DB.WithContext(ctx).Preload("Department").Preload("Handle").Where("employee_id = ?", id).First(&employee)
	if result.RowsAffected == 0 {
		panic(exception.NotFoundError{
			Message: "Employee Not Found",
		})
	}
	return employee
}

func (r *employeeRepositoryImpl) FindAll(ctx context.Context) []entity.Employee {
	var employee []entity.Employee
	result := r.DB.WithContext(ctx).Preload("Department").Preload("Handle").Find(&employee)
	if result.RowsAffected == 0 {
		panic(exception.NotFoundError{
			Message: "Data Not Found",
		})
	}
	return employee
}

func (r *employeeRepositoryImpl) FindByName(ctx context.Context, name string) entity.Employee {
	var employee entity.Employee
	result := r.DB.WithContext(ctx).Preload("Department").Preload("Handle").Where("name = ?", name).First(&employee)
	if result.RowsAffected == 0 {
		panic(exception.NotFoundError{
			Message: "Employee Not Found",
		})
	}
	return employee
}
