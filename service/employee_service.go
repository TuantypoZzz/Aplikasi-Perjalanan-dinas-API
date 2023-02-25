package service

import (
	"context"
	"golang-todo-app/entity"
	"golang-todo-app/model"
)

type EmployeeService interface {
	Create(ctx context.Context, employeeModel model.CreateEmployeeModel, user entity.User) entity.Employee
	// update(ctx context.Context, employeeModel model.UpdateEmployee, id string, user entity.User) model.UpdateEmployee
	// Delete(ctx context.Context, id string)
	// FindById(ctx context.Context, id string) model.EmployeeModel
	// FindAll(ctx context.Context) []model.EmployeeModel
}
