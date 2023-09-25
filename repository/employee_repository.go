package repository

import (
	"context"
	"golang-todo-app/entity"
)

type EmployeeRepository interface {
	Insert(ctx context.Context, employee entity.Employee) entity.Employee
	Update(ctx context.Context, employee entity.Employee) entity.Employee
	Delete(ctx context.Context, employee entity.Employee)
	FindById(ctx context.Context, id string) entity.Employee
	FindAll(ctx context.Context) []entity.Employee
	FindByName(ctx context.Context, name string) entity.Employee
}
