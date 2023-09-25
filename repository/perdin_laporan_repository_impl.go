package repository

import (
	"context"
	"golang-todo-app/entity"
	"golang-todo-app/exception"

	"gorm.io/gorm"
)

type perdinLaporanRepositoryImpl struct {
	*gorm.DB
}

func NewPerdinLaporanRepositoryImpl(DB *gorm.DB) PerdinLaporanRepository {
	return &perdinLaporanRepositoryImpl{DB: DB}
}

func (repo *perdinLaporanRepositoryImpl) FindAllPerdin(ctx context.Context) []entity.BussinessTravelReport {
	var perdinLaporan []entity.BussinessTravelReport
	result := repo.DB.WithContext(ctx).Find(&perdinLaporan)
	if result.RowsAffected == 0 {
		panic(exception.NotFoundError{
			Message: "Data not found",
		})
	}

	return perdinLaporan

}
