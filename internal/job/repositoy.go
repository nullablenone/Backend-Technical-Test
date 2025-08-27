package job

import "gorm.io/gorm"

type Repository interface {
	CreateJob(job *Job) error
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
