package service

import (
	"context"
	"golang-todo-app/entity"
	"golang-todo-app/model"
)

type LumpsumService interface {
	Create(ctx context.Context, lumpsumModel model.CreateLumpsum, user entity.User)
	Update(ctx context.Context, lumpsumModel model.UpdateLumpsum, id string, user entity.User)
	Delete(ctx context.Context, id string)
	FindById(ctx context.Context, id string) model.Lumpsum
	FindAll(ctx context.Context) []model.Lumpsum
}
