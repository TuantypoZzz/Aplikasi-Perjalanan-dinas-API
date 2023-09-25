package repository

import (
	"context"
	"golang-todo-app/entity"
)

type UploadRepository interface {
	CreateUpload(ctx context.Context, upload entity.Image)
}
