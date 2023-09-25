package service

import (
	"context"
	"golang-todo-app/entity"
	"golang-todo-app/model"
	"golang-todo-app/repository"
	"golang-todo-app/validation"

	"github.com/google/uuid"
)

type documentServiceImpl struct {
	repository.DocumentRepository
	repository.EmployeePerdinRepository
	repository.PerdinRepository
}

func NewDocumentImpl(documentRepository *repository.DocumentRepository, emplyPerdinRepository *repository.EmployeePerdinRepository, perdinRepository *repository.PerdinRepository) DocumentService {
	return &documentServiceImpl{DocumentRepository: *documentRepository, EmployeePerdinRepository: *emplyPerdinRepository, PerdinRepository: *perdinRepository}
}

func (service *documentServiceImpl) CreateDocument(ctx context.Context, document model.DokProofPerdin, file string) {
	validation.Validate(document)

	perdinEmployee := entity.PerdinEmployee{
		Id: uuid.MustParse(document.EmplyPerdinId),
	}

	documentProof := entity.DokProofPerdin{
		Id:               uuid.New(),
		Files:            file,
		PerdinEmployeeId: perdinEmployee.Id,
	}

	service.DocumentRepository.InsertDocument(ctx, documentProof)
	service.EmployeePerdinRepository.SetDetailPerdin(ctx, perdinEmployee.Id.String(), perdinEmployee)
}
