package repository

import (
	"context"
	"golang-todo-app/entity"
	"golang-todo-app/exception"

	"gorm.io/gorm"
)

type telaahRepositoryImpl struct {
	*gorm.DB
}

func NewTelaahRepositoryImpl(DB *gorm.DB) TelaahRepository {
	return &telaahRepositoryImpl{DB: DB}
}

func (repository *telaahRepositoryImpl) InsertTelaah(ctx context.Context, telaah entity.Telaah) entity.Telaah {
	errs := repository.DB.WithContext(ctx).Create(&telaah).Error
	exception.PanicLogging(errs)
	return telaah
}

func (repository *telaahRepositoryImpl) FindAllTelaah(ctx context.Context) []entity.Telaah {
	var telaah []entity.Telaah
	result := repository.WithContext(ctx).Preload("Activity.Employee").Preload("BussinessTravelReports").Find(&telaah)
	if result.RowsAffected == 0 {
		panic(exception.NotFoundError{
			Message: "Data not found",
		})
	}
	return telaah
}
