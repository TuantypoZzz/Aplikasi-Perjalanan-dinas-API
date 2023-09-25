package service

import (
	"context"
	"golang-todo-app/entity"
	"golang-todo-app/model"
)

type EmployeeService interface {
	Create(ctx context.Context, employeeModel model.CreateEmployeeModel, user entity.User, photo string)
	Update(ctx context.Context, employeeModel model.UpdateEmployee, id string, user entity.User, photo string)
	Delete(ctx context.Context, id string)
	FindById(ctx context.Context, id string) model.EmployeeModel
	FindAll(ctx context.Context) []model.EmployeeModel
}
