package job

import (
	"redikru-test/internal/domain/company"
	"time"
)

type CreateJobRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	CompanyID   string `json:"company_id" binding:"required"`
}

type GetAllJobsRequest struct {
	Keyword     string `form:"keyword"`
	CompanyName string `form:"companyName"`
}

type JobResponse struct {
	ID          string                  `json:"id"`
	Title       string                  `json:"title"`
	Description string                  `json:"description"`
	CreatedAt   time.Time               `json:"created_at"`
	Company     company.CompanyResponse `json:"company"`
}
