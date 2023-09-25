package repository

import (
	"context"
	"golang-todo-app/entity"
)

type TransportRepository interface {
	InsertTransport(ctx context.Context, transport []entity.Transport)
	SetTransportFlag(ctx context.Context, transportId string, flagTransport entity.Transport)
}
