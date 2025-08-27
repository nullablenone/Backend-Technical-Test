package main

import (
	"redikru-test/config"
	"redikru-test/routes"
)

func main() {

	// setup config
	config.LoadENV()
	config.ConnectDB()

	router := routes.SetupRoutes()
	router.Run(":8080")
}
