package job

import "redikru-test/utils"

type CreateJobRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	CompanyID   string `json:"company_id" binding:"required,uuid"`
}

type GetAllJobsRequest struct {
	Keyword     string `form:"keyword"`
	CompanyName string `form:"companyName"`
	Page        int    `form:"page"`
	Limit       int    `form:"limit"`
}

// cache
type cachedJobsResponse struct {
	Jobs       []Job            `json:"jobs"`
	Pagination utils.Pagination `json:"pagination"`
}

// khusus untuk swaggo
type ResponseSuccessJob struct {
	Status  string `json:"status" example:"success"`
	Message string `json:"message" example:"Job posting created successfully"`
	Data    Job    `json:"data"`
}

type ResponseSuccessGetJobs struct {
	Status     string           `json:"status" example:"sukses"`
	Message    string           `json:"message" example:"Permintaan berhasil"`
	Data       []Job            `json:"data"`
	Pagination utils.Pagination `json:"pagination"`
}


