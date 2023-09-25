package service

import (
	"context"
	"golang-todo-app/model"
)

type PerdinService interface {
	CreateTelaah(ctx context.Context, telaahModel model.CreateTelaahModel) model.TelaahModel
	FindallTelaah(ctx context.Context) []model.TelaahModel
	CreateSpt(ctx context.Context, sptModel model.CreateSptModel)
	FindAllSpt(ctx context.Context) []model.SptModel
	CreateSppd(ctx context.Context, sppdModel model.CreateSppdModel)
	FindAllSppd(ctx context.Context) []model.SppdModel
	CreateRincianPegawaiPerdin(ctx context.Context, detailPegawaiPerdinModel model.CreateDetailPerdin)
	CreatePerdinReport(ctx context.Context, perdinReport model.CreatePerdinReportModel, fotoPath string)
	GetPerdinReportByID(ctx context.Context, perdinID string) model.CreatePerdinReportModel
	GetEmployeePerdinByName(ctx context.Context, employeePerdinID string) model.EmployeePerdinModel
}
