package service

import (
	"context"
	"golang-todo-app/entity"
	"golang-todo-app/model"
	"golang-todo-app/repository"
	"golang-todo-app/validation"

	"github.com/google/uuid"
)

type departmentServiceImpl struct {
	repository.DepartmentRepository
}

func NewDepartmentRepositoryImpl(departmentRepository *repository.DepartmentRepository) DepartmentService {
	return &departmentServiceImpl{DepartmentRepository: *departmentRepository}
}

func (s *departmentServiceImpl) Create(c context.Context, DepartmentModel model.CreateDepartment, user entity.User) entity.Department {
	validation.Validate(DepartmentModel)

	Department := entity.Department{
		Id:             uuid.New(),
		DepartmentName: DepartmentModel.DepartmentName,
		CreatedBy:      user.Username,
		UpdatedBy:      user.Username,
	}

	Department = s.DepartmentRepository.Insert(c, Department)

	return Department
}

func (s *departmentServiceImpl) Update(c context.Context, DepartmentModel model.UpdateDepartment, id string, user entity.User) model.UpdateDepartment {
	validation.Validate(DepartmentModel)

	Department := entity.Department{
		Id:             uuid.MustParse(id),
		DepartmentName: DepartmentModel.DepartmentName,
		UpdatedBy:      user.Username,
	}
	Department = s.DepartmentRepository.Update(c, Department)

	return model.UpdateDepartment{
		DepartmentName: Department.DepartmentName,
	}
}

func (s *departmentServiceImpl) Delete(c context.Context, id string) {
	Department := s.DepartmentRepository.FindById(c, id)
	s.DepartmentRepository.Delete(c, Department)
}

func (s *departmentServiceImpl) FindAll(c context.Context) (responses []model.DepartmentModel) {
	positions := s.DepartmentRepository.FindAll(c)

	if len(positions) == 0 {
		return []model.DepartmentModel{}
	}

	for _, Department := range positions {
		responses = append(responses, model.DepartmentModel{
			Id:             Department.Id.String(),
			DepartmentName: Department.DepartmentName,
		})
	}
	return responses
}

func (s *departmentServiceImpl) FindById(c context.Context, id string) model.DepartmentModel {
	Department := s.DepartmentRepository.FindById(c, id)

	return model.DepartmentModel{
		Id:             Department.Id.String(),
		DepartmentName: Department.DepartmentName,
	}
}
