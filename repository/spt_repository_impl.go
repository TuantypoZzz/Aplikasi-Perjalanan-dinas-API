package repository

import (
	"context"
	"golang-todo-app/entity"
	"golang-todo-app/exception"

	"gorm.io/gorm"
)

type sptRepositoryImpl struct {
	*gorm.DB
}

func NewSptRepositoryImpl(DB *gorm.DB) SptRepository {
	return &sptRepositoryImpl{DB: DB}
}

func (repository *sptRepositoryImpl) InsertSpt(ctx context.Context, Spt entity.Spt) entity.Spt {
	err := repository.DB.WithContext(ctx).Create(&Spt).Error
	exception.PanicLogging(err)
	return Spt
}

func (repository *sptRepositoryImpl) FindAllSpt(ctx context.Context) []entity.Spt {
	var spt []entity.Spt
	result := repository.WithContext(ctx).Preload("Telaah.Activity").Preload("BussinessTravelReports").Preload("Employee").Find(&spt)
	if result.RowsAffected == 0 {
		panic(exception.NotFoundError{
			Message: "Data not found",
		})
	}
	return spt
}
