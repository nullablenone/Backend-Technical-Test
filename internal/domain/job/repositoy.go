package job

import (
	"errors"
	appErrors "redikru-test/internal/errors"

	"gorm.io/gorm"
)

type Repository interface {
	CreateJob(job *Job) error
	GetAllJob(request GetAllJobsRequest) ([]Job, int64, error)
}

type repository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{DB: db}
}

func (r *repository) CreateJob(job *Job) error {
	err := r.DB.Create(job).Error

	if err != nil {
		return appErrors.ErrInternalServer
	}

	err = r.DB.Preload("Company").First(job, "id = ?", job.ID).Error
	if err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return appErrors.ErrNotFound
		}

		return appErrors.ErrInternalServer
	}

	return nil
}

func (r *repository) GetAllJob(request GetAllJobsRequest) ([]Job, int64, error) { // <-- Ubah return
	var jobs []Job
	var totalRecords int64

	query := r.DB.Model(&Job{}).Preload("Company")
	countQuery := r.DB.Model(&Job{})

	// filter ke kedua query
	if request.Keyword != "" {
		searchKeyword := "%" + request.Keyword + "%"
		query = query.Where("title ILIKE ? OR description ILIKE ?", searchKeyword, searchKeyword)
		countQuery = countQuery.Where("title ILIKE ? OR description ILIKE ?", searchKeyword, searchKeyword)
	}
	if request.CompanyName != "" {
		joinQuery := "JOIN companies ON companies.id = jobs.company_id"
		whereQuery := "companies.name ILIKE ?"
		query = query.Joins(joinQuery).Where(whereQuery, "%"+request.CompanyName+"%")
		countQuery = countQuery.Joins(joinQuery).Where(whereQuery, "%"+request.CompanyName+"%")
	}

	// hitung total record
	if err := countQuery.Count(&totalRecords).Error; err != nil {
		return nil, 0, appErrors.ErrInternalServer
	}

	// hitung offset, terapkan pagination
	offset := (request.Page - 1) * request.Limit
	err := query.Order("created_at DESC").Limit(request.Limit).Offset(offset).Find(&jobs).Error
	if err != nil {
		return nil, 0, appErrors.ErrInternalServer
	}

	return jobs, totalRecords, nil
}
