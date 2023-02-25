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

func (s *employeeServiceImpl) Create(ctx context.Context, employeeModel model.CreateEmployeeModel, user entity.User) entity.Employee {
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
		DepartmenId:            uuid.MustParse(employeeModel.DepartmenId),
		Field:                  employeeModel.Field,
		Section:                employeeModel.Section,
		UnitWork:               employeeModel.UnitWork,
		Gender:                 gender.String(),
		BirthPlace:             employeeModel.BirthPlace,
		BirthDate:              birtDate,
		Phone:                  employeeModel.Phone,
		Email:                  employeeModel.Email,
		Address:                employeeModel.Address,
		Photo:                  employeeModel.Photo,
		BussinessTravelReports: []entity.BussinessTravelReport{},
	}

	employee = s.EmployeeRepository.Insert(ctx, employee)

	return employee
}

// func (service *employeeServiceImpl) Update(ctx context.Context, employeeModel model.UpdateEmployee, id string, user entity.User) model.UpdateEmployee {
// 	validation.Validate(employeeModel)

// 	employee := entity.Employee{
// 		Id:           uuid.MustParse(id),
// 		Name:         employeeModel.Name,
// 		NIP:          employeeModel.NIP,
// 		Bidang:       employeeModel.Bidang,
// 		Seksi:        employeeModel.Seksi,
// 		UnitKerja:    employeeModel.UnitKerja,
// 		JenisKelamin: employeeModel.JenisKelamin,
// 		TempatLahir:  employeeModel.TempatLahir,
// 		TanggalLahir: employeeModel.TanggalLahir,
// 		NoTlpn:       employeeModel.NoTlpn,
// 		Email:        employeeModel.Email,
// 		Alamat:       employeeModel.Alamat,
// 	}
// 	employees := service.EmployeeRepository.Update(ctx, employee)
// 	return model.UpdateEmployee{
// 		Name:         employees.Name,
// 		NIP:          employees.NIP,
// 		Bidang:       employees.Bidang,
// 		Seksi:        employees.Seksi,
// 		UnitKerja:    employees.UnitKerja,
// 		JenisKelamin: employees.JenisKelamin,
// 		TempatLahir:  employees.TempatLahir,
// 		TanggalLahir: employees.TanggalLahir,
// 		NoTlpn:       employees.NoTlpn,
// 		Email:        employees.Email,
// 		Alamat:       employees.Alamat,
// 	}
// }

// func (service *employeeServiceImpl) Delete(ctx context.Context, id string) {
// 	employees := service.EmployeeRepository.FindById(ctx, id)
// 	service.EmployeeRepository.Delete(ctx, employees)
// }

// func (service *employeeServiceImpl) FindAll(ctx context.Context) (response []model.EmployeeModel) {
// 	employees := service.EmployeeRepository.FindAll(ctx)

// 	if len(employees) == 0 {
// 		return []model.EmployeeModel{}
// 	}

// 	for _, employee := range employees {
// 		response = append(response, model.EmployeeModel{
// 			Id:     employee.Id.String(),
// 			Name:   employee.Name,
// 			NIP:    employee.NIP,
// 			Bidang: employee.Bidang,
// 		})
// 	}
// 	return response
// }

// func (service *employeeServiceImpl) FindById(ctx context.Context, id string) model.EmployeeModel {
// 	employees := service.EmployeeRepository.FindById(ctx, id)

// 	return model.EmployeeModel{
// 		Id:     employees.Id.String(),
// 		Name:   employees.Name,
// 		NIP:    employees.NIP,
// 		Bidang: employees.Bidang,
// 	}
// }
