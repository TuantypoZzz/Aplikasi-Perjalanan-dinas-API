package repository

import (
	"context"
	"golang-todo-app/entity"
	"golang-todo-app/exception"

	"gorm.io/gorm"
)

type transportRepositoryImpl struct {
	*gorm.DB
}

func NewTransportRepositoryImpl(DB *gorm.DB) TransportRepository {
	return &transportRepositoryImpl{DB: DB}
}

func (repository *transportRepositoryImpl) InsertTransport(ctx context.Context, transport []entity.Transport) {
	err := repository.DB.WithContext(ctx).Create(&transport).Error
	exception.PanicLogging(err)
}

func (repository *transportRepositoryImpl) SetTransportFlag(ctx context.Context, transportId string, flagTransport entity.Transport) {
	err := repository.DB.WithContext(ctx).Model(&flagTransport).Where("transport_id = ?", transportId).Update("flag_transport", 1).Error
	exception.PanicLogging(err)
}
