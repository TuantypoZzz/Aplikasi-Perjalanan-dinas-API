package repository

import (
	"context"
	"golang-todo-app/entity"
)

type DocumentRepository interface {
	InsertDocument(ctx context.Context, Document entity.DokProofPerdin)
}
