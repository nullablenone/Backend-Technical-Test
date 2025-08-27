package job

import "github.com/google/uuid"

type Service interface {
	CreateJob(request CreateJobRequest) (*Job, error)
	GetAllJob(request GetAllJobsRequest) ([]Job, error)
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

func (s *service) GetAllJob(request GetAllJobsRequest) ([]Job, error) {
	jobs, err := s.repository.GetAllJob(request)
	if err != nil {
		return nil, err
	}
	return jobs, nil
}
