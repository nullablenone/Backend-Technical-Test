package job

import (
	"math"
	"redikru-test/utils"

	"github.com/google/uuid"
)

type Service interface {
	CreateJob(request CreateJobRequest) (*Job, error)
	GetAllJob(request GetAllJobsRequest) ([]Job, utils.Pagination, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository}
}

func (s *service) CreateJob(request CreateJobRequest) (*Job, error) {
	job := Job{
		ID:          uuid.NewString(),
		Title:       request.Title,
		Description: request.Description,
		CompanyID:   request.CompanyID,
	}

	err := s.repository.CreateJob(&job)
	if err != nil {
		return nil, err
	}
	return &job, nil
}

func (s *service) GetAllJob(request GetAllJobsRequest) ([]Job, utils.Pagination, error) { // <-- Ubah return
	jobs, totalRecords, err := s.repository.GetAllJob(request)
	if err != nil {
		return nil, utils.Pagination{}, err
	}

	// Hitung total halaman
	totalPages := int(math.Ceil(float64(totalRecords) / float64(request.Limit)))

	// Siapkan struct pagination
	pagination := utils.Pagination{
		CurrentPage:  request.Page,
		PerPage:      request.Limit,
		TotalPages:   totalPages,
		TotalRecords: totalRecords,
	}

	return jobs, pagination, nil
}
