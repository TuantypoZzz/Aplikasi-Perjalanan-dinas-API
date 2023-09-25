package repository

import (
	"context"
	"golang-todo-app/entity"
	"golang-todo-app/exception"

	"gorm.io/gorm"
)

type accommodationRepositoryImpl struct {
	*gorm.DB
}

func NewAccomodationRepositoryImpl(DB *gorm.DB) AccommodationRepository {
	return &accommodationRepositoryImpl{DB: DB}
}

func (repository *accommodationRepositoryImpl) InsertHotel(ctx context.Context, hotel entity.Accommodation) entity.Accommodation {
	err := repository.DB.WithContext(ctx).Create(&hotel).Error
	exception.PanicLogging(err)
	return hotel
}
