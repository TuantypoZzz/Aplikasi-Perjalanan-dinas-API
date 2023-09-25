package repository

import (
	"context"
	"golang-todo-app/entity"
)

type TelaahRepository interface {
	InsertTelaah(ctx context.Context, telaah entity.Telaah) entity.Telaah
	FindAllTelaah(ctx context.Context) []entity.Telaah
}
