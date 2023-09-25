package repository

import (
	"context"
	"golang-todo-app/entity"
	"golang-todo-app/exception"

	"gorm.io/gorm"
)

type departmentRepositoryImpl struct {
	*gorm.DB
}

func NewDepartmentRepositoryImpl(DB *gorm.DB) DepartmentRepository {
	return &departmentRepositoryImpl{DB: DB}
}

func (r *departmentRepositoryImpl) Insert(ctx context.Context, department entity.Department) entity.Department {
	err := r.DB.WithContext(ctx).Create(&department).Error
	exception.PanicLogging(err)
	return department
}

func (r *departmentRepositoryImpl) Update(ctx context.Context, department entity.Department) entity.Department {
	err := r.DB.WithContext(ctx).Where("department_id = ?", department.Id).Updates(&department).Error
	exception.PanicLogging(err)
	return department
}

func (r *departmentRepositoryImpl) Delete(ctx context.Context, department entity.Department) {
	r.DB.WithContext(ctx).Where("department_id =?", department.Id).Delete(&department)
}

func (r *departmentRepositoryImpl) FindById(ctx context.Context, id string) entity.Department {
	var department entity.Department
	result := r.DB.WithContext(ctx).Where("department_id = ?", id).First(&department)
	if result.RowsAffected == 0 {
		panic(exception.NotFoundError{
			Message: "Department Not Found",
		})
	}
	return department
}

func (r *departmentRepositoryImpl) FindAll(ctx context.Context) []entity.Department {
	var positions []entity.Department
	r.DB.WithContext(ctx).Find(&positions)
	return positions
}
