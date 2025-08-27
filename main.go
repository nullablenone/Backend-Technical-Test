package main

import (
	"log"
	"redikru-test/config"
	"redikru-test/internal/domain/company"
	"redikru-test/internal/domain/job"

	"redikru-test/routes"
	"redikru-test/seeders"
)

func main() {

	// setup config
	config.LoadENV()
	db := config.ConnectDB()

	// setup table
	err := db.AutoMigrate(company.Company{}, job.Job{})
	if err != nil {
		log.Fatal(err)
	}

	// exe seeder
	seeders.SeedCompanies(db)

	// setup domain job
	jobRepository := job.NewRepository(db)
	jobService := job.NewService(jobRepository)
	jobHandler := job.NewHandler(jobService)

	router := routes.SetupRoutes(jobHandler)
	router.Run(":8080")
}
