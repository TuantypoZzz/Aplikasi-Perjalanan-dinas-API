package repository

import (
	"context"
	"golang-todo-app/entity"
	"golang-todo-app/exception"

	"gorm.io/gorm"
)

type uploadRepositoryImpl struct {
	*gorm.DB
}

func NewUploadRepositoryImpl(DB *gorm.DB) UploadRepository {
	return &uploadRepositoryImpl{DB: DB}
}

func (repository *uploadRepositoryImpl) CreateUpload(ctx context.Context, upload entity.Image) {
	err := repository.DB.WithContext(ctx).Create(&upload).Error
	exception.PanicLogging(err)
}
