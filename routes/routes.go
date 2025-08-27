package routes

import (
	"redikru-test/internal/domain/job"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(jobHandler *job.Handler) *gin.Engine {
	router := gin.Default()

	router.POST("/jobs", jobHandler.CreateJobHandler)
	router.GET("/jobs", jobHandler.GetAllJobHandler)

	return router
}
