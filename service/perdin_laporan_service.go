package service

import (
	"context"
	"golang-todo-app/model"
)

type PerdinLaporanService interface {
	FindAllPerdin(ctx context.Context) []model.RekapPerdin
}
