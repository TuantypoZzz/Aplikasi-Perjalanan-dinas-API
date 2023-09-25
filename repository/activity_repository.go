package repository

import (
	"context"
	"golang-todo-app/entity"
)

type ActivityRepository interface {
	Insert(ctx context.Context, activity entity.Activity) entity.Activity
	Update(ctx context.Context, activity entity.Activity) entity.Activity
	Delete(ctx context.Context, activity entity.Activity)
	FindAll(ctx context.Context) []entity.Activity
	FindById(ctx context.Context, id string) entity.Activity
}
