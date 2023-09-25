package repository

import (
	"context"
	"golang-todo-app/entity"
)

type SptRepository interface {
	InsertSpt(ctx context.Context, Spt entity.Spt) entity.Spt
	FindAllSpt(ctx context.Context) []entity.Spt
}
