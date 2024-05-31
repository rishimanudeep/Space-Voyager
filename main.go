package main

import (
	"gofr.dev/pkg/gofr"

	handlerPlanet "Exo-planet/handler"
	servicePlanet "Exo-planet/service"
)

func main() {
	// Initialize a new instance of the GoFr framework
	app := gofr.New()

	// Initialise Planet Service
	planetService := servicePlanet.New()

	// injected planet service into the handler
	planetHandler := handlerPlanet.New(planetService)

	// routes for handling planet data
	// Route to get all exoplanets
	app.GET("/exoplanet", planetHandler.GetAll)

	// Route to create a new exoplanet
	app.POST("/exoplanet", planetHandler.Create)

	// Route to get exoplanet by ID
	app.GET("/exoplanet/{id}", planetHandler.GetByID)

	// Route to update an exoplanet by ID
	app.PUT("/exoplanet/{id}", planetHandler.Update)

	// Route to delete exoplanet by ID
	app.DELETE("/exoplanet/{id}", planetHandler.Delete)

	// Route to get fuel estimation for an exoplanet by ID
	app.GET("/exoplanet/{id}/fuel-estimation", planetHandler.GetEstimation)

	// Run the application
	app.Run()
}
