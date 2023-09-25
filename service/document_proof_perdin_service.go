package service

import (
	"context"
	"golang-todo-app/model"
)

type DocumentService interface {
	CreateDocument(ctx context.Context, document model.DokProofPerdin, files string)
}
