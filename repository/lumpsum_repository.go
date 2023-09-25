package repository

import (
	"context"
	"golang-todo-app/entity"
)

type LumpsumRepository interface {
	Insert(ctx context.Context, lumpsum entity.Lumpsum) entity.Lumpsum
	Update(ctx context.Context, lumpsum entity.Lumpsum) entity.Lumpsum
	Delete(ctx context.Context, lumpsum entity.Lumpsum)
	FindAll(ctx context.Context) []entity.Lumpsum
	FindById(ctx context.Context, id string) entity.Lumpsum
}
