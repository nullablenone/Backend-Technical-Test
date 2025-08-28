package utils

import (
	"errors"
	"log"
	"net/http"

	appErrors "redikru-test/internal/errors"

	"github.com/gin-gonic/gin"
)

func RespondError(c *gin.Context, status int, message string) {
	c.JSON(status, gin.H{
		"status":  "error",
		"message": message,
	})
}

func RespondSuccess(c *gin.Context,  status int, data any, message string) {
	c.JSON(status, gin.H{
		"status":  "success",
		"message": message,
		"data":    data,
	})
}

func HandleError(c *gin.Context, err error) {
	log.Printf("[ERROR] Detail: %v\n", err)

	switch {
	case errors.Is(err, appErrors.ErrNotFound):
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "error",
			"message": "Data yang diminta tidak ditemukan.",
		})

	case errors.Is(err, appErrors.ErrInternalServer):
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Terjadi kesalahan internal pada server.",
		})

	default:
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Terjadi kesalahan yang tidak diketahui.",
		})
	}
}
