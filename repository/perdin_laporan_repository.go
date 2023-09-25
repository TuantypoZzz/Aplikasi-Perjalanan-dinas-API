package repository

import (
	"context"
	"golang-todo-app/entity"
)

type PerdinLaporanRepository interface {
	FindAllPerdin(ctx context.Context) []entity.BussinessTravelReport
}
