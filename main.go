package main

import (
	"log"
	"redikru-test/config"
	"redikru-test/internal/domain/company"
	"redikru-test/internal/domain/job"

	"redikru-test/routes"
	"redikru-test/seeders"

	_ "redikru-test/docs" 
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)



// @title API Dokumentasi untuk Technical Test Redikru
// @version 1.0
// @description Ini adalah dokumentasi API untuk backend service lowongan pekerjaan yang dibuat sebagai bagian dari proses seleksi di Redikru.

// @contact.name (Nama Kamu)
// @contact.email (Email Kamu)

// @host localhost:8080
// @BasePath /
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

	// setup redis
	redisClient := config.ConnectRedis()

	// setup domain job
	jobRepository := job.NewRepository(db)
	jobService := job.NewService(jobRepository, redisClient)
	jobHandler := job.NewHandler(jobService)

	router := routes.SetupRoutes(jobHandler)

	// swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Run(":8080")
}
