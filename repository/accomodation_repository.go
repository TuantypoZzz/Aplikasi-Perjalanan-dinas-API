package repository

import (
	"context"
	"golang-todo-app/entity"
)

type AccommodationRepository interface {
	InsertHotel(ctx context.Context, hotel entity.Accommodation) entity.Accommodation
}
