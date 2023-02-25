package service

import (
	"context"
	"golang-todo-app/entity"
	"golang-todo-app/model"
	"golang-todo-app/repository"
	"golang-todo-app/validation"

	"github.com/google/uuid"
)

type HandleServiceImpl struct {
	repository.HandleRepository
}

func NewHandleRepositoryImpl(HandleRepository *repository.HandleRepository) HandleService {
	return &HandleServiceImpl{HandleRepository: *HandleRepository}
}

func (s *HandleServiceImpl) Create(c context.Context, pangkatgolongaModel model.CreateHandle, user entity.User) entity.Handle {
	validation.Validate(pangkatgolongaModel)

	pangkatgol := entity.Handle{
		Id:           uuid.New(),
		NamaPangkat:  pangkatgolongaModel.NamaPangkat,
		NamaGolongan: pangkatgolongaModel.NamaGolongan,
		Tingkat:      pangkatgolongaModel.Tingkat,
		CreatedBy:    user.Username,
		UpdatedBy:    user.Username,
	}

	pangkatgol = s.HandleRepository.Insert(c, pangkatgol)
	return pangkatgol
}

func (s *HandleServiceImpl) Update(c context.Context, Handlemodel model.UpdateHandle, id string, user entity.User) model.UpdateHandle {
	validation.Validate(Handlemodel)

	pangkatgol := entity.Handle{
		Id:           uuid.MustParse(id),
		NamaPangkat:  Handlemodel.NamaPangkat,
		NamaGolongan: Handlemodel.NamaGolongan,
		Tingkat:      Handlemodel.Tingkat,
		UpdatedBy:    user.Username,
	}
	pangkatgol = s.HandleRepository.Update(c, pangkatgol)
	return model.UpdateHandle{
		NamaPangkat:  pangkatgol.NamaPangkat,
		NamaGolongan: pangkatgol.NamaGolongan,
		Tingkat:      pangkatgol.Tingkat,
	}
}

func (s *HandleServiceImpl) Delete(c context.Context, id string) {
	pangkatgol := s.HandleRepository.FindById(c, id)
	s.HandleRepository.Delete(c, pangkatgol)
}

func (s *HandleServiceImpl) FindAll(c context.Context) (response []model.HandleModel) {
	pangkatgol := s.HandleRepository.FindAll(c)

	if len(pangkatgol) == 0 {
		return []model.HandleModel{}
	}

	for _, pangkatgol := range pangkatgol {
		response = append(response, model.HandleModel{
			Id:           pangkatgol.Id.String(),
			NamaPangkat:  pangkatgol.NamaPangkat,
			NamaGolongan: pangkatgol.NamaGolongan,
			Tingkat:      pangkatgol.Tingkat,
		})
	}
	return response
}

func (s *HandleServiceImpl) FindById(c context.Context, id string) model.HandleModel {
	pangkatgol := s.HandleRepository.FindById(c, id)

	return model.HandleModel{
		Id:           pangkatgol.Id.String(),
		NamaPangkat:  pangkatgol.NamaPangkat,
		NamaGolongan: pangkatgol.NamaGolongan,
		Tingkat:      pangkatgol.Tingkat,
	}
}
