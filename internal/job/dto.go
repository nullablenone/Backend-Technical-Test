package job

type CreateJobRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	CompanyID   string `json:"company_id" binding:"required"`
}
