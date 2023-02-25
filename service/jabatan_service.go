package service

import (
	"context"
	"golang-todo-app/entity"
	"golang-todo-app/model"
)

type DepartmentService interface {
	Create(c context.Context, departmentModel model.CreateDepartment, user entity.User) entity.Department
	Update(c context.Context, departmentModel model.UpdateDepartment, id string, user entity.User) model.UpdateDepartment
	Delete(c context.Context, id string)
	FindById(c context.Context, id string) model.DepartmentModel
	FindAll(c context.Context) []model.DepartmentModel
}
