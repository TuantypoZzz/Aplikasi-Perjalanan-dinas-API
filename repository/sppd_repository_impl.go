package repository

import (
	"context"
	"golang-todo-app/entity"
	"golang-todo-app/exception"

	"gorm.io/gorm"
)

type sppdRepositoryImpl struct {
	*gorm.DB
}

func NewSppdRepositoryImpl(DB *gorm.DB) SppdRepository {
	return &sppdRepositoryImpl{DB: DB}
}

func (repository *sppdRepositoryImpl) InsertSppd(ctx context.Context, Sppd entity.Sppd) entity.Sppd {
	err := repository.DB.WithContext(ctx).Create(&Sppd).Error
	exception.PanicLogging(err)
	return Sppd
}

func (repository *sppdRepositoryImpl) FindAllSppd(ctx context.Context) []entity.Sppd {
	var sppd []entity.Sppd
	result := repository.DB.WithContext(ctx).Preload("Employee").Preload("BussinessTravelReports").
		Find(&sppd)
	if result.RowsAffected == 0 {
		panic(exception.NotFoundError{
			Message: "Data not found",
		})
	}
	return sppd
}
func (repository *sppdRepositoryImpl) GetSPPDNoByBussinessTravelReportID(ctx context.Context, bussinessTravelReportID string) entity.Sppd {
	var sppd entity.Sppd
	result := repository.DB.WithContext(ctx).Preload("BussinessTravelReports.Telaah.Activity").Where("bussiness_travel_report_id = ?", bussinessTravelReportID).First(&sppd)
	if result.RowsAffected == 0 {
		panic(exception.NotFoundError{
			Message: "Data not found",
		})
	}
	return sppd
}
