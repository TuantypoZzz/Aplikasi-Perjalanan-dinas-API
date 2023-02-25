package repository

import (
	"context"
	"golang-todo-app/entity"
)

type HandleRepository interface {
	Insert(ctx context.Context, handle entity.Handle) entity.Handle
	Update(ctx context.Context, handle entity.Handle) entity.Handle
	Delete(ctx context.Context, handle entity.Handle)
	FindById(ctx context.Context, id string) entity.Handle
	FindAll(ctx context.Context) []entity.Handle
}
