package service

import (
	"context"
	"golang-todo-app/entity"
	"golang-todo-app/exception"
	"golang-todo-app/model"
	"golang-todo-app/repository"
	"golang-todo-app/validation"
	"strconv"
	"time"

	"github.com/google/uuid"
)

type perdinServiceImpl struct {
	repository.PerdinRepository
	repository.TelaahRepository
	repository.SptRepository
	repository.SppdRepository
	repository.TransportRepository
	repository.AccommodationRepository
	repository.EmployeePerdinRepository
}

func NewPerdinServiceImpl(perdinRepository *repository.PerdinRepository, telaahRepository *repository.TelaahRepository, sptRepository *repository.SptRepository, sppdRepository *repository.SppdRepository, transportRepository *repository.TransportRepository, hotelRepository *repository.AccommodationRepository, emplyPerdinRepository *repository.EmployeePerdinRepository) PerdinService {
	return &perdinServiceImpl{PerdinRepository: *perdinRepository, TelaahRepository: *telaahRepository, SptRepository: *sptRepository, SppdRepository: *sppdRepository, TransportRepository: *transportRepository, AccommodationRepository: *hotelRepository, EmployeePerdinRepository: *emplyPerdinRepository}
}

var timeLayout = "2006-01-02"

// --------------------- TELAAH ----------------------------------

func (service *perdinServiceImpl) CreateTelaah(ctx context.Context, telaahModel model.CreateTelaahModel) model.TelaahModel {
	validation.Validate(telaahModel)

	telaahDate, err := time.Parse(timeLayout, telaahModel.TglTelaah)
	exception.PanicLogging(err)

	berangkatDate, err := time.Parse(timeLayout, telaahModel.TglBerangkat)
	exception.PanicLogging(err)

	kembaliDate, err := time.Parse(timeLayout, telaahModel.TglKembali)
	exception.PanicLogging(err)

	var employees []entity.Employee

	for _, employeeId := range telaahModel.PegawaiPengikut {
		employees = append(employees, entity.Employee{Id: uuid.MustParse(employeeId)})
	}
	employees = append(employees, entity.Employee{Id: uuid.MustParse(telaahModel.PegawaiPemimpin)})

	perdin := entity.BussinessTravelReport{
		Id:           uuid.New(),
		TglBerangkat: &berangkatDate,
		TglKembali:   &kembaliDate,
	}

	telaah := entity.Telaah{
		Id:                     uuid.New(),
		ActivityId:             uuid.MustParse(telaahModel.Activity),
		NoTelaah:               telaahModel.NoTelaah,
		TglTelaah:              &telaahDate,
		BussinessTravelReports: perdin,
	}

	service.PerdinRepository.InsertPegawaiPerdin(ctx, perdin, employees)
	service.TelaahRepository.InsertTelaah(ctx, telaah)
	service.PerdinRepository.SetPerdinEmployeeCommand(ctx, perdin.Id.String(), telaahModel.PegawaiPemimpin, entity.PerdinEmployee{CommandFlag: 1})

	return model.TelaahModel{
		Id:                      telaah.Id.String(),
		Activity:                telaahModel.Activity,
		NoTelaah:                telaah.NoTelaah,
		TglTelaah:               telaahModel.TglTelaah,
		TglBerangkat:            perdin.TglBerangkat.String(),
		TglKembali:              perdin.TglKembali.String(),
		PegawaiPemimpin:         telaahModel.PegawaiPemimpin,
		PegawaiPengikut:         telaahModel.PegawaiPengikut,
		NamaPPTK:                telaahModel.Activity,
		IdBussinessTravelReport: perdin.Id.String(),
	}
}

func (service *perdinServiceImpl) FindallTelaah(ctx context.Context) (response []model.TelaahModel) {
	telaah := service.TelaahRepository.FindAllTelaah(ctx)

	if len(telaah) == 0 {
		return []model.TelaahModel{}
	}

	for _, telaah := range telaah {

		employeesPerdin := service.PerdinRepository.FindAllEmployeePerdinById(ctx, telaah.BussinessTravelReportId.String())
		pemimpin, pengikut := getEmployeePerdin(employeesPerdin)

		response = append(response, model.TelaahModel{
			Id:                      telaah.Id.String(),
			Activity:                telaah.Activity.NamaKegiatan,
			NoTelaah:                telaah.NoTelaah,
			TglTelaah:               telaah.TglTelaah.String(),
			TglBerangkat:            telaah.BussinessTravelReports.TglBerangkat.String(),
			TglKembali:              telaah.BussinessTravelReports.TglKembali.String(),
			PegawaiPemimpin:         pemimpin,
			PegawaiPengikut:         pengikut,
			NamaPPTK:                telaah.Activity.Employee.Name,
			IdBussinessTravelReport: telaah.BussinessTravelReportId.String(),
		})
	}
	return response
}

func getEmployeePerdin(employees []entity.PerdinEmployee) (pemimpin string, pengikut []string) {
	for _, emply := range employees {
		if emply.CommandFlag == 1 {
			pemimpin = emply.Employee.Name
		} else {
			pengikut = append(pengikut, emply.Employee.Name)
		}
	}
	return
}

// --------------------- SPT ----------------------------------

func (service *perdinServiceImpl) CreateSpt(ctx context.Context, sptModel model.CreateSptModel) {
	validation.Validate(sptModel)

	ReportDate, err := time.Parse(timeLayout, sptModel.TglLapor)
	exception.PanicLogging(err)
	PublishedSptDate, err := time.Parse(timeLayout, sptModel.TglSpt)
	exception.PanicLogging(err)
	SkDate, err := time.Parse(timeLayout, sptModel.TglSkMewakili)
	exception.PanicLogging(err)

	perdin := entity.BussinessTravelReport{
		Id:            uuid.MustParse(sptModel.PerdinId),
		KotaAsal:      sptModel.KotaAsal,
		KotaTujuan:    sptModel.KotaTujuan,
		TglLapor:      &ReportDate,
		NoSkMewaliki:  sptModel.NoSkMewaliki,
		TglSkMewakili: &SkDate,
		Keperluan:     sptModel.Keperluan,
	}

	spt := entity.Spt{
		Id:                     uuid.New(),
		NoSpt:                  sptModel.NomorSurat,
		TglSPT:                 &PublishedSptDate,
		PejabatSPTId:           uuid.MustParse(sptModel.PejabatTtdSpt),
		BussinessTravelReports: perdin,
		TelaahId:               uuid.MustParse(sptModel.TelaahId),
	}

	service.PerdinRepository.SetSPTPerdin(ctx, perdin.Id.String(), perdin)
	service.SptRepository.InsertSpt(ctx, spt)
}

func (service *perdinServiceImpl) FindAllSpt(ctx context.Context) (response []model.SptModel) {
	Spt := service.SptRepository.FindAllSpt(ctx)

	if len(Spt) == 0 {
		return []model.SptModel{}
	}

	for _, spt := range Spt {
		employeesPerdin := service.PerdinRepository.FindAllEmployeePerdinById(ctx, spt.BussinessTravelReportId.String())
		pemimpin, pengikut := getEmployeePerdin(employeesPerdin)

		response = append(response, model.SptModel{
			Id:              spt.Id.String(),
			NomorSurat:      spt.NoSpt,
			Keperluan:       spt.BussinessTravelReports.Keperluan,
			NamaKegiatan:    spt.Telaah.Activity.NamaKegiatan,
			KotaAsal:        spt.BussinessTravelReports.KotaAsal,
			KotaTujuan:      spt.BussinessTravelReports.KotaTujuan,
			PegawaiPemimpin: pemimpin,
			PegawaiPengikut: pengikut,
			TglBerangkat:    spt.BussinessTravelReports.TglBerangkat.String(),
			TglKembali:      spt.BussinessTravelReports.TglKembali.String(),
			TglLapor:        spt.BussinessTravelReports.TglLapor.String(),
			TglSpt:          spt.TglSPT.String(),
			PejabatTtdSpt:   spt.Employee.Name,
			NoSkMewaliki:    spt.NoSpt,
			TglSkMewakili:   spt.BussinessTravelReports.TglSkMewakili.String(),
		})
	}
	return response
}

func (service *perdinServiceImpl) CreateSppd(ctx context.Context, sppdModel model.CreateSppdModel) {
	validation.Validate(sppdModel)

	SppdDate, err := time.Parse(timeLayout, sppdModel.TanggalSppd)
	exception.PanicLogging(err)

	sppd := entity.Sppd{
		Id:                      uuid.New(),
		NoSppd:                  sppdModel.NomorSurat,
		TglSppd:                 &SppdDate,
		PejabatSPTId:            uuid.MustParse(sppdModel.PejabatTtdSppd),
		BussinessTravelReportId: uuid.MustParse(sppdModel.PerdinId),
	}

	perdin := entity.BussinessTravelReport{
		Id:     uuid.MustParse(sppdModel.PerdinId),
		SppdId: sppd.Id,
	}
	service.SppdRepository.InsertSppd(ctx, sppd)
	service.PerdinRepository.SetSPTPerdin(ctx, perdin.Id.String(), perdin)

}

func (service *perdinServiceImpl) FindAllSppd(ctx context.Context) (response []model.SppdModel) {
	sppd := service.SppdRepository.FindAllSppd(ctx)

	if len(sppd) == 0 {
		return []model.SppdModel{}
	}

	for _, Sppd := range sppd {
		employeesPerdin := service.PerdinRepository.FindAllEmployeePerdinById(ctx, Sppd.BussinessTravelReportId.String())
		pemimpin, pengikut := getEmployeePerdin(employeesPerdin)

		pangkatPemimpinPerdin, jabatanPemimpinPerdin, tingkatPemimpinPerdin := getEmployeePerdinPG(employeesPerdin)

		response = append(response, model.SppdModel{
			Id:                           Sppd.Id.String(),
			NomorSurat:                   Sppd.NoSppd,
			PejabatTtdSppd:               Sppd.Employee.Name,
			PegawaiPemimpin:              pemimpin,
			PangkatPegawaiPemimpin:       pangkatPemimpinPerdin,
			JabatanPegawaiPemimpin:       jabatanPemimpinPerdin,
			TingkatPerdinPegawaiPemimpin: tingkatPemimpinPerdin,
			//JenisAngkutan: Sppd.tran,

			PegawaiPengikut: pengikut,
			KotaAsal:        Sppd.BussinessTravelReports.KotaAsal,
			KotaTujuan:      Sppd.BussinessTravelReports.KotaTujuan,
			TglBerangkat:    Sppd.BussinessTravelReports.TglBerangkat.String(),
			TglKembali:      Sppd.BussinessTravelReports.TglKembali.String(),
			TanggalSppd:     Sppd.TglSppd.String(),
			Lama:            strconv.Itoa(calculateTravelDuration(*Sppd.BussinessTravelReports.TglBerangkat, *Sppd.BussinessTravelReports.TglKembali)),
		})
	}
	return response
}

func getEmployeePerdinPG(employees []entity.PerdinEmployee) (pangkatPemimpinPerdin string, jabatanPemimpinPerdin string, tingkatPemimpinPerdin string) {
	for _, emply := range employees {
		pangkatPemimpinPerdin = emply.Employee.Handle.NamaPangkat
		jabatanPemimpinPerdin = emply.Employee.Department.DepartmentName
		tingkatPemimpinPerdin = emply.Employee.Handle.Tingkat

	}
	return
}

func (service *perdinServiceImpl) CreateRincianPegawaiPerdin(ctx context.Context, detailPegawaiPerdinModel model.CreateDetailPerdin) {
	validation.Validate(detailPegawaiPerdinModel)

	timeLayout := "2006-01-02"

	ChekinDate, err := time.Parse(timeLayout, detailPegawaiPerdinModel.TglCekin)
	exception.PanicLogging(err)
	ToGoFlightDate, err := time.Parse(timeLayout, detailPegawaiPerdinModel.AngkutanPergi.TanggalPenerbangan)
	exception.PanicLogging(err)
	ReturnFlightDate, err := time.Parse(timeLayout, detailPegawaiPerdinModel.AngkutanPulang.TanggalPenerbangan)
	exception.PanicLogging(err)
	ChekoutDate, err := time.Parse(timeLayout, detailPegawaiPerdinModel.TglCekout)
	exception.PanicLogging(err)

	// menghitung harga tiket transportasi
	tiketPergi := float64(detailPegawaiPerdinModel.AngkutanPergi.HargaTiket)
	tiketPulang := float64(detailPegawaiPerdinModel.AngkutanPulang.HargaTiket)
	biayaTaksiKotaAsal := float64(detailPegawaiPerdinModel.AngkutanPergi.TaksiPengeluaranRill)
	biayaTaksiKotaTujuan := float64(detailPegawaiPerdinModel.AngkutanPulang.TaksiPengeluaranRill)
	biayaTransportDarat := float64(detailPegawaiPerdinModel.TiketDarat)
	totalBiayaTransport := calculateTravelTotalCost(tiketPergi, tiketPulang, biayaTaksiKotaAsal, biayaTaksiKotaTujuan, biayaTransportDarat)

	// menghitung biaya penginapan permalam
	pricePerNight := float64(detailPegawaiPerdinModel.BiayaPenginapan)
	// durasi lama mengunap
	durasi := calculateTravelDuration(ChekinDate, ChekoutDate)
	// durasi perjalanan
	durasiPerdin := calculateTravelDuration(ToGoFlightDate, ReturnFlightDate)
	dailyMoney := detailPegawaiPerdinModel.UangHarian

	totalDailyMoney := calculateAccommodationTotalNight(float64(dailyMoney), durasiPerdin)

	totalPenginapan := calculateAccommodationTotalNight(pricePerNight, durasi)

	totalPerdinCostEmployee := calculateTotalCost(tiketPergi, tiketPulang, biayaTaksiKotaAsal, biayaTaksiKotaTujuan, biayaTransportDarat, pricePerNight, float64(dailyMoney), durasi)

	perdin := entity.PerdinEmployee{
		Id:                      uuid.MustParse(detailPegawaiPerdinModel.NamaPemimpin),
		DailyMoney:              detailPegawaiPerdinModel.UangHarian,
		LodgingExpenses:         detailPegawaiPerdinModel.BiayaPenginapan,
		TotalTransportationCost: int(totalBiayaTransport),
		TotalHotelCost:          int(totalPenginapan),
		TotalDailyCost:          int(totalDailyMoney),
		TotalPerdinCost:         int(totalPerdinCostEmployee),
	}

	transport := []entity.Transport{
		{
			Id:               uuid.New(),
			HargaTiketPergi:  tiketPergi,
			NamaAngkutan:     detailPegawaiPerdinModel.AngkutanPergi.JenisAngkutan,
			NoTiket:          detailPegawaiPerdinModel.AngkutanPergi.JenisAngkutan,
			NoPenerbangan:    detailPegawaiPerdinModel.AngkutanPergi.NomorPenerbangan,
			KelasPenerbangan: detailPegawaiPerdinModel.AngkutanPergi.KelasPenerbangan,
			KodeBooking:      detailPegawaiPerdinModel.AngkutanPergi.KodeBooking,
			TglPenerbangan:   &ToGoFlightDate,
			Taksi:            biayaTaksiKotaAsal,
			PerdinEmployeeId: perdin.Id,
			FlagTransport:    1,
		},
		{
			Id:               uuid.New(),
			HargaTiketPulang: tiketPulang,
			NamaAngkutan:     detailPegawaiPerdinModel.AngkutanPulang.JenisAngkutan,
			NoTiket:          detailPegawaiPerdinModel.AngkutanPulang.JenisAngkutan,
			NoPenerbangan:    detailPegawaiPerdinModel.AngkutanPulang.NomorPenerbangan,
			KelasPenerbangan: detailPegawaiPerdinModel.AngkutanPulang.KelasPenerbangan,
			KodeBooking:      detailPegawaiPerdinModel.AngkutanPulang.KodeBooking,
			TglPenerbangan:   &ReturnFlightDate,
			Taksi:            biayaTaksiKotaTujuan,
			PerdinEmployeeId: perdin.Id,
			FlagTransport:    0,
		},
	}

	hotel := entity.Accommodation{
		Id:                  uuid.New(),
		BiayaTransportDarat: biayaTransportDarat,
		BiayaPermalam:       pricePerNight,
		NamaHotel:           detailPegawaiPerdinModel.NamaHotel,
		NoKamar:             detailPegawaiPerdinModel.NomorKamar,
		CheckInDate:         ChekinDate,
		CheckOutDate:        ChekoutDate,
		TipeKamar:           detailPegawaiPerdinModel.TipeKamar,
		LamaMenginap:        strconv.Itoa(durasi),
		PerdinEmployeeId:    perdin.Id,
	}
	perdinUpdate := entity.BussinessTravelReport{
		Id:              uuid.MustParse(detailPegawaiPerdinModel.PerdinId),
		AccommodationId: hotel.Id,
	}

	service.TransportRepository.InsertTransport(ctx, transport)
	service.AccommodationRepository.InsertHotel(ctx, hotel)
	service.EmployeePerdinRepository.SetDetailPerdin(ctx, perdin.Id.String(), perdin)
	service.PerdinRepository.SetSPTPerdin(ctx, perdinUpdate.Id.String(), perdinUpdate)
	//service.TransportRepository.SetTransportFlag(ctx, transport.Id.String(), detailPegawaiPerdinModel.AngkutanPergi, entity.Transport{FlagTransport: 1})

}

func calculateTravelDuration(keberangkatan time.Time, kembali time.Time) int {
	duration := kembali.Sub(keberangkatan)
	return int(duration.Hours() / 24)
}

func calculateTravelTotalCost(tiketPergi, tiketPulang, biayaTaksiKotaAsal, biayaTaksiKotaTujuan, biayaTransportDarat float64) float64 {
	return tiketPergi + tiketPulang + biayaTaksiKotaAsal + biayaTaksiKotaTujuan + biayaTransportDarat
}
func calculateAccommodationTotalNight(pricePerNight float64, durasi int) float64 {
	return pricePerNight * float64(durasi)
}
func calculateTotalCost(tiketPergi, tiketPulang, biayaTaksiKotaAsal, biayaTaksiKotaTujuan, biayaTransportDarat, pricePerNight, dailyMoney float64, durasi int) float64 {
	travelTotalCost := calculateTravelTotalCost(tiketPergi, tiketPulang, biayaTaksiKotaAsal, biayaTaksiKotaTujuan, biayaTransportDarat)
	accommodationTotalCost := calculateAccommodationTotalNight(pricePerNight, durasi)
	dailyTotalCost := calculateAccommodationTotalNight(float64(dailyMoney), durasi)

	totalCost := travelTotalCost + accommodationTotalCost + dailyTotalCost

	return totalCost
}

func (service *perdinServiceImpl) CreatePerdinReport(ctx context.Context, perdinReport model.CreatePerdinReportModel, fotoPath string) {
	validation.Validate(perdinReport)

	perdinReports := entity.BussinessTravelReport{
		Id:                     uuid.MustParse(perdinReport.PerdinId),
		MaksudPerdin:           perdinReport.MaksudPerdin,
		TujuanPerdin:           perdinReport.TujuanPerdin,
		WaktuPelaksanaanPerdin: perdinReport.WaktuPelaksanaanPerdin,
		HasilYangDicapai:       perdinReport.HasilYangdiCapai,
		SaranTindakanPerdin:    perdinReport.SaranDanTindakLanjut,
		Foto:                   fotoPath,
	}

	service.PerdinRepository.SetSPTPerdin(ctx, perdinReports.Id.String(), perdinReports)
}

func (service *perdinServiceImpl) GetPerdinReportByID(ctx context.Context, perdinID string) model.CreatePerdinReportModel {
	Getperdin := service.PerdinRepository.FindPerdinById(ctx, perdinID)

	return model.CreatePerdinReportModel{
		PerdinId:               Getperdin.Id.String(),
		MaksudPerdin:           Getperdin.MaksudPerdin,
		TujuanPerdin:           Getperdin.TujuanPerdin,
		WaktuPelaksanaanPerdin: Getperdin.WaktuPelaksanaanPerdin,
		Foto:                   Getperdin.Foto,
	}
}
func (service *perdinServiceImpl) GetEmployeePerdinByName(ctx context.Context, employeePerdinID string) model.EmployeePerdinModel {
	EmplyPerdin := service.PerdinRepository.FindEmployeePerdinByName(ctx, employeePerdinID)
	nama := getEmployeePerdinByName(EmplyPerdin)
	return model.EmployeePerdinModel{
		IdPerdinEmployee:        EmplyPerdin.Id.String(),
		IdBussinessTravelReport: EmplyPerdin.BussinessTravelReportId.String(),
		Nama:                    nama,
	}
}
func getEmployeePerdinByName(employees entity.PerdinEmployee) string {
	return employees.Employee.Name
}
