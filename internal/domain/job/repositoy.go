package job

import "gorm.io/gorm"

type Repository interface {
	CreateJob(job *Job) error
	GetAllJob(request GetAllJobsRequest) ([]Job, error)
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
		return err
	}

	err = r.DB.Preload("Company").First(job, "id = ?", job.ID).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) GetAllJob(request GetAllJobsRequest) ([]Job, error) {

	var jobs []Job

	query := r.DB.Preload("Company").Order("created_at DESC")

	if request.Keyword != "" {
		searchKeyword := "%" + request.Keyword + "%"
		query = query.Where("title ILIKE ? OR description ILIKE ?", searchKeyword, searchKeyword)
	}

	if request.CompanyName != "" {
		query = query.Joins("JOIN companies ON companies.id = jobs.company_id").
			Where("companies.name ILIKE ?", "%"+request.CompanyName+"%")
	}

	err := query.Find(&jobs).Error
	if err != nil {
		return nil, err
	}

	return jobs, nil
}
