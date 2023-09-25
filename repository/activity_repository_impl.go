package repository

import (
	"context"
	"golang-todo-app/entity"
	"golang-todo-app/exception"

	"gorm.io/gorm"
)

type activityRepositoryImpl struct {
	*gorm.DB
}

func NewActivityRepositoryImpl(DB *gorm.DB) ActivityRepository {
	return &activityRepositoryImpl{DB: DB}
}

func (repository *activityRepositoryImpl) Insert(ctx context.Context, activity entity.Activity) entity.Activity {
	err := repository.DB.WithContext(ctx).Create(&activity).Error
	exception.PanicLogging(err)
	return activity
}

func (r *activityRepositoryImpl) Update(ctx context.Context, activity entity.Activity) entity.Activity {
	err := r.DB.WithContext(ctx).Where("activity_id = ?", activity.Id).Updates(&activity).Error
	exception.PanicLogging(err)
	return activity
}

func (r *activityRepositoryImpl) Delete(ctx context.Context, activity entity.Activity) {
	r.DB.WithContext(ctx).Where("activity_id =?", activity.Id).Delete(&activity)
}

func (r *activityRepositoryImpl) FindById(ctx context.Context, id string) entity.Activity {
	var activity entity.Activity
	result := r.DB.WithContext(ctx).Preload("Employee").Where("activity_id = ?", id).First(&activity)
	if result.RowsAffected == 0 {
		panic(exception.NotFoundError{
			Message: "Activity Not Found",
		})
	}
	return activity
}

func (repository *activityRepositoryImpl) FindAll(ctx context.Context) []entity.Activity {
	var activity []entity.Activity
	result := repository.DB.WithContext(ctx).Preload("Employee").Find(&activity)

	if result.RowsAffected == 0 {
		panic(exception.NotFoundError{
			Message: "Data Not Found",
		})
	}
	return activity
}
