package service

import (
	"context"
	"golang-todo-app/entity"
	"golang-todo-app/model"
	"golang-todo-app/repository"
	"golang-todo-app/validation"

	"github.com/google/uuid"
)

type activityServiceImpl struct {
	repository.ActivityRepository
}

func NewActivityServiceImpl(activityRepository *repository.ActivityRepository) ActivityService {
	return &activityServiceImpl{ActivityRepository: *activityRepository}
}

func (service *activityServiceImpl) Create(ctx context.Context, activityModel model.CreateActivityModel, user entity.User) {
	validation.Validate(activityModel)

	activity := entity.Activity{
		Id:            uuid.New(),
		NamaKegiatan:  activityModel.ActivityName,
		MataAnggaran:  activityModel.BudgetLine,
		PejabatPPTKId: uuid.MustParse(activityModel.EmployeeId),
		Instansi:      activityModel.Instances,
		CreatedBy:     user.Username,
	}
	service.ActivityRepository.Insert(ctx, activity)
}

func (service *activityServiceImpl) Update(ctx context.Context, activityModel model.UpdateActivityModel, id string) {
	validation.Validate(activityModel)

	activities := entity.Activity{
		Id:            uuid.MustParse(id),
		NamaKegiatan:  activityModel.ActivityName,
		MataAnggaran:  activityModel.BudgetLine,
		PejabatPPTKId: uuid.MustParse(activityModel.EmployeeId),
		Instansi:      activityModel.Instances,
	}
	service.ActivityRepository.Update(ctx, activities)
}

func (service *activityServiceImpl) Delete(ctx context.Context, id string) {
	activities := service.ActivityRepository.FindById(ctx, id)
	service.ActivityRepository.Delete(ctx, activities)
}

func (service *activityServiceImpl) FindAll(ctx context.Context) (response []model.ActivityModel) {
	activity := service.ActivityRepository.FindAll(ctx)

	if len(activity) == 0 {
		return []model.ActivityModel{}
	}

	for _, activities := range activity {
		response = append(response, model.ActivityModel{
			Id:           activities.Id.String(),
			ActivityName: activities.NamaKegiatan,
			BudgetLine:   activities.MataAnggaran,
			EmployeeId:   activities.Employee.Name,
			Instances:    activities.Instansi,
		})
	}
	return response
}

func (service *activityServiceImpl) FindById(ctx context.Context, id string) model.ActivityModel {
	activities := service.ActivityRepository.FindById(ctx, id)

	return model.ActivityModel{
		Id:           activities.Id.String(),
		ActivityName: activities.NamaKegiatan,
		BudgetLine:   activities.MataAnggaran,
		EmployeeId:   activities.Employee.Name,
		Instances:    activities.Instansi,
	}
}
