package job

import (
	"net/http"
	"redikru-test/utils"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{Service: service}
}

func (h *Handler) CreateJobHandler(c *gin.Context) {
	var request CreateJobRequest

	err := c.ShouldBindJSON(&request)

	if err != nil {
		utils.RespondError(c, http.StatusBadRequest, "Invalid request payload")
		return
	}

	newJob, err := h.Service.CreateJob(request)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	utils.RespondSuccess(c, http.StatusCreated, newJob, "Job posting created successfully")

}

func (h *Handler) GetAllJobHandler(c *gin.Context) {
	var request GetAllJobsRequest

	err := c.ShouldBindQuery(&request)
	if err != nil {
		utils.RespondError(c, http.StatusBadRequest, "Parameter query tidak valid")
		return
	}

	if request.Page <= 0 {
		request.Page = 1
	}
	if request.Limit <= 0 {
		request.Limit = 10
	}

	if request.Limit > 100 {
		request.Limit = 100
	}

	jobs, pagination, err := h.Service.GetAllJob(request)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	utils.RespondSuccessWithPagination(c, http.StatusOK, jobs, pagination, "Permintaan berhasil")

}
