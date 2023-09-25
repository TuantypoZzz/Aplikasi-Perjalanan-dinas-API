package service

import (
	"context"
	"golang-todo-app/entity"
	"golang-todo-app/model"
)

type ActivityService interface {
	Create(ctx context.Context, activityModel model.CreateActivityModel, user entity.User)
	Update(ctx context.Context, activityModel model.UpdateActivityModel, id string)
	Delete(ctx context.Context, id string)
	FindById(ctx context.Context, id string) model.ActivityModel
	FindAll(ctx context.Context) []model.ActivityModel
}
