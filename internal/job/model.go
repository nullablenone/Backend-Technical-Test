package job

import (
	"redikru-test/internal/company"
	"time"
)

type Job struct {
	ID          string `gorm:"type:uuid;primary_key"`
	Title       string `gorm:"type:text;not null"`
	Description string `gorm:"type:text;not null"`
	CreatedAt   time.Time

	// foreign key & relationship
	CompanyID string          `gorm:"type:uuid;not null"`
	Company   company.Company `gorm:"foreignKey:CompanyID"`
}
