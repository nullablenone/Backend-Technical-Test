package job

import (
	"net/http"

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
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	newJob, err := h.Service.CreateJob(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Job posting created successfully",
		"data":    newJob,
	})

}

func (h *Handler) GetAllJobHandler(c *gin.Context) {

	var request GetAllJobsRequest

	err := c.ShouldBindQuery(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	jobs, err := h.Service.GetAllJob(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve job postings"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": jobs,
	})
}
