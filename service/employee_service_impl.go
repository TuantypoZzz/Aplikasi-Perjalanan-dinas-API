package service

import (
	"context"
	"golang-todo-app/entity"
	"golang-todo-app/enum"
	"golang-todo-app/exception"
	"golang-todo-app/model"
	"golang-todo-app/repository"
	"golang-todo-app/validation"
	"time"

	"github.com/google/uuid"
)

type employeeServiceImpl struct {
	repository.EmployeeRepository
}

func NewEmployeeRepositoryImpl(employeeRepository *repository.EmployeeRepository) EmployeeService {
	return &employeeServiceImpl{EmployeeRepository: *employeeRepository}
}

func (s *employeeServiceImpl) Create(ctx context.Context, employeeModel model.CreateEmployeeModel, user entity.User, photo string) {
	validation.Validate(employeeModel)

	gender, err := enum.NewGender(employeeModel.Gender)
	exception.PanicLogging(err)

	birtDate, err := time.Parse("2006-01-02", employeeModel.BirthDate)
	exception.PanicLogging(err)

	employee := entity.Employee{
		Id:                     uuid.New(),
		Name:                   employeeModel.Name,
		NIP:                    employeeModel.NIP,
		HandleId:               uuid.MustParse(employeeModel.HandleId),
		DepartmentId:           uuid.MustParse(employeeModel.DepartmenId),
		Field:                  employeeModel.Field,
		Section:                employeeModel.Section,
		UnitWork:               employeeModel.UnitWork,
		Gender:                 gender.String(),
		BirthPlace:             employeeModel.BirthPlace,
		BirthDate:              birtDate,
		Phone:                  employeeModel.Phone,
		Email:                  employeeModel.Email,
		Address:                employeeModel.Address,
		Photo:                  photo,
		CreatedBy:              user.Username,
		BussinessTravelReports: []entity.BussinessTravelReport{},
	}

	s.EmployeeRepository.Insert(ctx, employee)
}

func (service *employeeServiceImpl) Update(ctx context.Context, employeeModel model.UpdateEmployee, id string, user entity.User, photo string) {
	validation.Validate(employeeModel)

	gender, err := enum.NewGender(employeeModel.JenisKelamin)
	exception.PanicLogging(err)

	birtDate, err := time.Parse("2006-01-02", employeeModel.TanggalLahir)
	exception.PanicLogging(err)

	employee := entity.Employee{
		Id:                     uuid.MustParse(id),
		Name:                   employeeModel.Name,
		NIP:                    employeeModel.NIP,
		HandleId:               uuid.MustParse(employeeModel.IdPangkat),
		DepartmentId:           uuid.MustParse(employeeModel.IdJabatan),
		Field:                  employeeModel.Bidang,
		Section:                employeeModel.Seksi,
		UnitWork:               employeeModel.UnitKerja,
		Gender:                 gender.String(),
		BirthPlace:             employeeModel.TempatLahir,
		BirthDate:              birtDate,
		Phone:                  employeeModel.NoTlpn,
		Email:                  employeeModel.Email,
		Address:                employeeModel.Alamat,
		Photo:                  photo,
		UpdatedBy:              user.Username,
		BussinessTravelReports: []entity.BussinessTravelReport{},
	}
	service.EmployeeRepository.Update(ctx, employee)
	// return &model.UpdateEmployee{
	// 	Name:         employees.Name,
	// 	NIP:          employees.NIP,
	// 	IdPangkat:    employees.Handle.NamaPangkat + "/" + employee.Handle.NamaGolongan,
	// 	IdJabatan:    employees.Department.DepartmentName,
	// 	Bidang:       employees.Field,
	// 	Seksi:        employees.Section,
	// 	UnitKerja:    employees.UnitWork,
	// 	JenisKelamin: employees.Gender,
	// 	TempatLahir:  employees.BirthPlace,
	// 	TanggalLahir: employees.BirthDate.String(),
	// 	NoTlpn:       employees.Phone,
	// 	Email:        employees.Email,
	// 	Alamat:       employees.Address,
	// 	Foto:         employees.Photo,
	// }
}

func (service *employeeServiceImpl) Delete(ctx context.Context, id string) {
	employees := service.EmployeeRepository.FindById(ctx, id)
	service.EmployeeRepository.Delete(ctx, employees)
}

func (service *employeeServiceImpl) FindAll(ctx context.Context) (response []model.EmployeeModel) {
	employees := service.EmployeeRepository.FindAll(ctx)

	if len(employees) == 0 {
		return []model.EmployeeModel{}
	}

	for _, employee := range employees {
		response = append(response, model.EmployeeModel{
			Id:           employee.Id.String(),
			Name:         employee.Name,
			NIP:          employee.NIP,
			Pangkat:      employee.Handle.NamaPangkat + "/" + employee.Handle.NamaGolongan,
			Jabatan:      employee.Department.DepartmentName,
			Bidang:       employee.Field,
			Seksi:        employee.Section,
			UnitKerja:    employee.UnitWork,
			JenisKelamin: employee.Gender,
			TempatLahir:  employee.BirthPlace,
			TanggalLahir: employee.BirthDate.String(),
			NoTlpn:       employee.Phone,
			Email:        employee.Email,
			Alamat:       employee.Address,
			Foto:         employee.Photo,
		})
	}
	return response
}

func (service *employeeServiceImpl) FindById(ctx context.Context, id string) model.EmployeeModel {
	employees := service.EmployeeRepository.FindById(ctx, id)

	return model.EmployeeModel{
		Id:           employees.Id.String(),
		Name:         employees.Name,
		NIP:          employees.NIP,
		Pangkat:      employees.Handle.NamaPangkat + "/" + employees.Handle.NamaGolongan,
		Jabatan:      employees.Department.DepartmentName,
		Bidang:       employees.Field,
		Seksi:        employees.Section,
		UnitKerja:    employees.UnitWork,
		JenisKelamin: employees.Gender,
		TempatLahir:  employees.BirthPlace,
		TanggalLahir: employees.BirthDate.String(),
		NoTlpn:       employees.Phone,
		Email:        employees.Email,
		Alamat:       employees.Address,
		Foto:         employees.Photo,
	}
}
