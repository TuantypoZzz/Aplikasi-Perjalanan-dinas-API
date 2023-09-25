package repository

import (
	"context"
	"golang-todo-app/entity"
	"golang-todo-app/exception"

	"gorm.io/gorm"
)

type handleRepositoryImpl struct {
	*gorm.DB
}

func NewHandleRepositoryImpl(DB *gorm.DB) HandleRepository {
	return &handleRepositoryImpl{DB: DB}
}

func (r *handleRepositoryImpl) Insert(ctx context.Context, handle entity.Handle) entity.Handle {
	err := r.DB.WithContext(ctx).Create(&handle).Error
	exception.PanicLogging(err)
	return handle
}

func (r *handleRepositoryImpl) Update(ctx context.Context, handle entity.Handle) entity.Handle {
	err := r.DB.WithContext(ctx).Where("handle_id = ?", handle.Id).Updates(&handle).Error
	exception.PanicLogging(err)
	return handle
}

func (r *handleRepositoryImpl) Delete(ctx context.Context, handle entity.Handle) {
	r.DB.WithContext(ctx).Where("handle_id =?", handle.Id).Delete(&handle)
}

func (r *handleRepositoryImpl) FindById(ctx context.Context, id string) entity.Handle {
	var handle entity.Handle
	result := r.DB.WithContext(ctx).Where("handle_id = ?", id).First(&handle)
	if result.RowsAffected == 0 {
		panic(exception.NotFoundError{
			Message: "Pangkat/Golongan Not Found",
		})
	}
	return handle
}

func (r *handleRepositoryImpl) FindAll(ctx context.Context) []entity.Handle {
	var handle []entity.Handle
	r.DB.WithContext(ctx).Find(&handle)
	return handle
}
