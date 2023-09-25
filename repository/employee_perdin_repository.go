package repository

import (
	"context"
	"golang-todo-app/entity"
)

type EmployeePerdinRepository interface {
	SetDetailPerdin(ctx context.Context, perdinEmpyId string, detailEmpyPerdin entity.PerdinEmployee)
}
