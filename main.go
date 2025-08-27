package main

import (
	"log"
	"redikru-test/config"
	"redikru-test/internal/company"
	"redikru-test/routes"
	"redikru-test/seeders"
)

func main() {

	// setup config
	config.LoadENV()
	db := config.ConnectDB()

	// setup table
	err := db.AutoMigrate(company.Company{})
	if err != nil {
		log.Fatal(err)
	}

	// exe seeder
	seeders.SeedCompanies(db)


	router := routes.SetupRoutes()
	router.Run(":8080")
}
