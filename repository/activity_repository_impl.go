package repository

import (
	"context"
	"golang-todo-app/entity"
	"golang-todo-app/exception"

	"gorm.io/gorm"
)

type activityRepositoryImpl struct {
	*gorm.DB
}

func NewActivityRepositoryImpl(DB *gorm.DB) ActivityRepository {
	return &activityRepositoryImpl{DB: DB}
}

func (repository *activityRepositoryImpl) Insert(ctx context.Context, activity entity.Activity) entity.Activity {
	// Raw SQL Insert query
	query := `
		INSERT INTO activities (activity_id, nama_kegiatan, mata_anggaran, pejabat_pptk_id, instansi, created_by)
		VALUES (?, ?, ?, ?, ?, ?)
	`

	// Execute raw SQL query
	err := repository.DB.WithContext(ctx).Exec(query, activity.Id, activity.NamaKegiatan, activity.MataAnggaran, activity.PejabatPPTKId, activity.Instansi, activity.CreatedBy).Error
	exception.PanicLogging(err)
	return activity
}

func (r *activityRepositoryImpl) Update(ctx context.Context, activity entity.Activity) entity.Activity {
	err := r.DB.WithContext(ctx).Where("activity_id = ?", activity.Id).Updates(&activity).Error
	exception.PanicLogging(err)
	return activity
}

func (r *activityRepositoryImpl) Delete(ctx context.Context, activity entity.Activity) {
	r.DB.WithContext(ctx).Where("activity_id =?", activity.Id).Delete(&activity)
}

func (r *activityRepositoryImpl) FindById(ctx context.Context, id string) entity.Activity {
	var result entity.ActivityWithEmployeeName

	query := `SELECT act.activity_id, act.nama_kegiatan, act.mata_anggaran, act.instansi, 
              emply.name FROM activities act 
              JOIN employees emply ON emply.employee_id = act.pejabat_pptk_id 
              WHERE act.activity_id = ?`

	dbResult := r.DB.WithContext(ctx).Raw(query, id).Scan(&result)
	if dbResult.RowsAffected == 0 {
		panic(exception.NotFoundError{
			Message: "Activity Not Found",
		})
	}

	// Pemetaan: Assign Employee.Name dari temporary field EmployeeName
	result.Activity.Employee.Name = result.PejabatName

	return result.Activity
}

func (repository *activityRepositoryImpl) FindAll(ctx context.Context) []entity.Activity {
	var activity []entity.Activity
	result := repository.DB.WithContext(ctx).Preload("Employee").Find(&activity)

	if result.RowsAffected == 0 {
		panic(exception.NotFoundError{
			Message: "Data Not Found",
		})
	}
	return activity
}
