package service

import (
	"context"
	"golang-todo-app/entity"
	"golang-todo-app/model"
	"golang-todo-app/repository"
	"golang-todo-app/validation"

	"github.com/google/uuid"
)

type lumpsumServiceImpl struct {
	repository.LumpsumRepository
}

func NewLumpsumRepositoryImpl(lumpsumRepository *repository.LumpsumRepository) LumpsumService {
	return &lumpsumServiceImpl{LumpsumRepository: *lumpsumRepository}
}

func (service *lumpsumServiceImpl) Create(ctx context.Context, lumpsumModel model.CreateLumpsum, user entity.User) {
	validation.Validate(lumpsumModel)

	lumpsum := entity.Lumpsum{
		Id:              uuid.New(),
		Lingkup:         lumpsumModel.Lingkup,
		BiayaHarian:     lumpsumModel.BiayaHarian,
		BiayaPenginapan: lumpsumModel.BiayaPenginapan,
		CreatedBy:       user.Username,
	}
	lumpsum = service.LumpsumRepository.Insert(ctx, lumpsum)
}

func (service *lumpsumServiceImpl) Update(ctx context.Context, lumpsumModel model.UpdateLumpsum, id string, user entity.User) {
	validation.Validate(lumpsumModel)

	lumpsum := entity.Lumpsum{
		Id:              uuid.MustParse(id),
		Lingkup:         lumpsumModel.Lingkup,
		BiayaHarian:     lumpsumModel.BiayaHarian,
		BiayaPenginapan: lumpsumModel.BiayaPenginapan,
		UpdatedBy:       user.Username,
	}
	lumpsum = service.LumpsumRepository.Update(ctx, lumpsum)
}

func (service *lumpsumServiceImpl) Delete(ctx context.Context, id string) {
	lumpsum := service.LumpsumRepository.FindById(ctx, id)
	service.LumpsumRepository.Delete(ctx, lumpsum)
}

func (service *lumpsumServiceImpl) FindById(ctx context.Context, id string) model.Lumpsum {
	lumpsum := service.LumpsumRepository.FindById(ctx, id)

	return model.Lumpsum{
		Id:              lumpsum.Id.String(),
		Lingkup:         lumpsum.Lingkup,
		BiayaHarian:     lumpsum.BiayaHarian,
		BiayaPenginapan: lumpsum.BiayaPenginapan,
	}
}

func (service *lumpsumServiceImpl) FindAll(ctx context.Context) (responses []model.Lumpsum) {
	lumpsum := service.LumpsumRepository.FindAll(ctx)

	if len(lumpsum) == 0 {
		return []model.Lumpsum{}
	}

	for _, Lumpsum := range lumpsum {
		responses = append(responses, model.Lumpsum{
			Id:              Lumpsum.Id.String(),
			Lingkup:         Lumpsum.Lingkup,
			BiayaHarian:     Lumpsum.BiayaHarian,
			BiayaPenginapan: Lumpsum.BiayaPenginapan,
		})
	}
	return responses
}
