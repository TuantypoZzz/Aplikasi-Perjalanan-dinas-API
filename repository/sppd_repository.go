package repository

import (
	"context"
	"golang-todo-app/entity"
)

type SppdRepository interface {
	InsertSppd(ctx context.Context, Sppd entity.Sppd) entity.Sppd
	FindAllSppd(ctx context.Context) []entity.Sppd
	GetSPPDNoByBussinessTravelReportID(ctx context.Context, bussinessTravelReportID string) entity.Sppd
}
