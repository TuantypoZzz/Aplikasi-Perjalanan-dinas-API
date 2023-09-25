package service

import (
	"context"
	"golang-todo-app/entity"
	"golang-todo-app/model"
	"golang-todo-app/repository"
)

type perdinLaporanServiceImpl struct {
	repository.PerdinLaporanRepository
	repository.PerdinRepository
	repository.SppdRepository
}

func NewPerdinLaporanServiceImpl(perdinLaporanImpl *repository.PerdinLaporanRepository, perdinRepository *repository.PerdinRepository) PerdinLaporanService {
	return &perdinLaporanServiceImpl{PerdinLaporanRepository: *perdinLaporanImpl, PerdinRepository: *perdinRepository}
}

func (service *perdinLaporanServiceImpl) FindAllPerdin(ctx context.Context) (response []model.RekapPerdin) {
	perdin := service.PerdinLaporanRepository.FindAllPerdin(ctx)

	if len(perdin) == 0 {
		return []model.RekapPerdin{}
	}

	for _, perdin := range perdin {

		employeesPerdin := service.PerdinRepository.FindAllEmployeePerdinById(ctx, perdin.Id.String())
		pemimpin, pengikut := getEmployeePerdin(employeesPerdin)
		sppd := service.SppdRepository.GetSPPDNoByBussinessTravelReportID(ctx, perdin.Id.String())
		durasiPerdin := calculateTravelDuration(*perdin.TglBerangkat, *perdin.TglKembali)
		NoSppd := getSppd(sppd)

		response = append(response, model.RekapPerdin{
			IdBussinessTravelReport: perdin.Id.String(),
			NoSppd:                  NoSppd,
			TglLapor:                perdin.TglLapor.String(),
			TglBerangkat:            perdin.TglBerangkat.String(),
			TglKembali:              perdin.TglKembali.String(),
			PegawaiPemimpin:         pemimpin,
			PegawaiPengikut:         pengikut,
			KotaTujuan:              perdin.KotaTujuan,
			LamaPerdin:              int16(durasiPerdin),
		})
	}
	return response
}
func getSppd(sppd entity.Sppd) (noSppd string) {
	noSppd = sppd.NoSppd
	return
}
