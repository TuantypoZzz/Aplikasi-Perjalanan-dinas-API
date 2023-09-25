package repository

import (
	"context"
	"golang-todo-app/entity"
	"golang-todo-app/exception"

	"gorm.io/gorm"
)

type lumpsumRepositoryImpl struct {
	*gorm.DB
}

func NewLumpsumRepositoryImpl(DB *gorm.DB) LumpsumRepository {
	return &lumpsumRepositoryImpl{DB: DB}
}

func (repository *lumpsumRepositoryImpl) Insert(ctx context.Context, lumpsum entity.Lumpsum) entity.Lumpsum {
	err := repository.DB.WithContext(ctx).Create(&lumpsum).Error

	exception.PanicLogging(err)
	return lumpsum
}

func (repository *lumpsumRepositoryImpl) Update(ctx context.Context, lumpsum entity.Lumpsum) entity.Lumpsum {
	err := repository.DB.WithContext(ctx).Where("lumpsum_id = ? ", lumpsum.Id).Updates(&lumpsum).Error

	exception.PanicLogging(err)
	return lumpsum
}
func (repository *lumpsumRepositoryImpl) Delete(ctx context.Context, lumpsum entity.Lumpsum) {
	repository.DB.WithContext(ctx).Where("lumpsum_id = ?", lumpsum.Id).Delete(&lumpsum)

}
func (repository *lumpsumRepositoryImpl) FindAll(ctx context.Context) []entity.Lumpsum {
	var lumpsums []entity.Lumpsum
	result := repository.DB.WithContext(ctx).Find(&lumpsums)
	if result.RowsAffected == 0 {
		panic(exception.NotFoundError{
			Message: "Data Not Found",
		})
	}
	return lumpsums
}
func (repository *lumpsumRepositoryImpl) FindById(ctx context.Context, id string) entity.Lumpsum {
	var lumpsums entity.Lumpsum
	result := repository.DB.WithContext(ctx).Where("lumpsum_id = ?", id).First(&lumpsums)
	if result.RowsAffected == 0 {
		panic(exception.NotFoundError{
			Message: "Lumpsum Not Found",
		})
	}
	return lumpsums
}
