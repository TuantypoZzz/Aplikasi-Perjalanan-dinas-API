package service

import (
	"context"
	"golang-todo-app/entity"
	"golang-todo-app/model"
)

type HandleService interface {
	Create(c context.Context, Handlemodel model.CreateHandle, user entity.User) entity.Handle
	Update(c context.Context, Handlemodel model.UpdateHandle, id string, user entity.User) model.UpdateHandle
	Delete(c context.Context, id string)
	FindById(c context.Context, id string) model.HandleModel
	FindAll(c context.Context) []model.HandleModel
}
