package repository

import (
	"context"
	"golang-todo-app/entity"
)

type DepartmentRepository interface {
	Insert(ctx context.Context, department entity.Department) entity.Department
	Update(ctx context.Context, department entity.Department) entity.Department
	Delete(ctx context.Context, department entity.Department)
	FindById(ctx context.Context, id string) entity.Department
	FindAll(ctx context.Context) []entity.Department
}
