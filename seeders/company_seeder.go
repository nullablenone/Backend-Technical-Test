package seeders

import (
	"log"
	"redikru-test/internal/company"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func SeedCompanies(db *gorm.DB) {
	
	companies := []company.Company{
		{Name: "Redikru"},
		{Name: "Google"},
		{Name: "Microsoft"},
	}

	for _, c := range companies {

		var companyRecord company.Company

		result := db.Where(company.Company{Name: c.Name}).
			Attrs(company.Company{ID: uuid.NewString()}).
			FirstOrCreate(&companyRecord)

		if result.Error != nil {
			log.Printf("Gagal menjalankan seeder untuk company '%s': %v\n", c.Name, result.Error)
			continue
		}

		if result.RowsAffected > 0 {
			log.Printf("Seeder: Company '%s' berhasil dibuat.", c.Name)
		}

	}
}
