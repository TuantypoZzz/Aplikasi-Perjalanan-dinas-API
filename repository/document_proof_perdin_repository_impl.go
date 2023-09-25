package repository

import (
	"context"
	"golang-todo-app/entity"
	"golang-todo-app/exception"

	"gorm.io/gorm"
)

type documentReporsitoryImpl struct {
	*gorm.DB
}

func NewDocumentRepositoryImpl(DB *gorm.DB) DocumentRepository {
	return &documentReporsitoryImpl{DB: DB}
}

func (repository *documentReporsitoryImpl) InsertDocument(ctx context.Context, document entity.DokProofPerdin) {
	err := repository.DB.WithContext(ctx).Create(&document).Error
	exception.PanicLogging(err)
}
