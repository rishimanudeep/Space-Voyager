package service

import (
	"fmt"
	"math"
	"net/http"
	"sort"

	"github.com/google/uuid"

	gofrHttp "gofr.dev/pkg/gofr/http"

	"Exo-planet/models"
)

var exoPlanets = make(map[uuid.UUID]models.Planet)

type service struct{}

// New nolint -New is a factory function that returns service.
func New() service {
	return service{}
}

// Create method of service struct responsible to create a new planet
func (s service) Create(planet models.Planet) (*models.Planet, error) {
	// checks the mandatory fields for the provided planet
	err := checkMandatory(planet)
	if err != nil {
		return nil, err
	}

	// validates planet parameters
	err = validatePlanet(planet)
	if err != nil {
		return nil, err
	}

	// generates new uuid for the planet's ID
	planet.ID = uuid.New()

	// Add the planet to the exoPlanets map using its generated ID as the key
	exoPlanets[planet.ID] = planet

	return &planet, nil
}

// GetByID retrieves planet by using id
func (s service) GetByID(id uuid.UUID) (*models.Planet, error) {
	// finds the id with the provided ID in the exoPlanets map
	planet, ok := exoPlanets[id]
	if !ok {
		return nil, gofrHttp.ErrorEntityNotFound{Name: "planet", Value: id.String()}
	}

	return &planet, nil
}

// GetAll fetches all the planets which present in the exoPlanets map
// sortBy will sort radius or mass in the ascending order
func (s service) GetAll(filter models.Filter) ([]models.Planet, error) {
	// initialise the empty slice to hold all the planets
	allPlanets := make([]models.Planet, 0)

	// iterates through the map and append each planet to slice
	for _, p := range exoPlanets {
		allPlanets = append(allPlanets, p)
	}

	// if sortBy provided in the filter
	if filter.SortBy != "" {
		switch filter.SortBy {
		case models.SortByRadius:
			// sort the planets by radius in ascending order
			sort.Slice(allPlanets, func(i, j int) bool {
				return allPlanets[i].Radius < allPlanets[j].Radius
			})
		case models.SortByMass:
			// Filter to only terrestrial planets
			terrestrialPlanets := make([]models.Planet, 0)
			for _, planet := range allPlanets {
				if planet.ExoplanetType == "Terrestrial" {
					terrestrialPlanets = append(terrestrialPlanets, planet)
				}
			}
			// sort the terrestrial planets by mass in ascending order
			sort.Slice(terrestrialPlanets, func(i, j int) bool {
				return terrestrialPlanets[i].Mass < terrestrialPlanets[j].Mass
			})
			return terrestrialPlanets, nil
		}
	}

	return allPlanets, nil
}

// Update will be responsible for updating an existing planet
func (s service) Update(planet models.Planet) error {
	// checks whether the planet exists in the map
	_, ok := exoPlanets[planet.ID]
	if !ok {
		return gofrHttp.ErrorEntityNotFound{Name: "planet", Value: planet.ID.String()}
	}

	// checks the mandatory fields for the provided planet
	err := checkMandatory(planet)
	if err != nil {
		return err
	}

	// validates planet parameters
	err = validatePlanet(planet)
	if err != nil {
		return err
	}

	// update the planet data in the exoPlanet map
	exoPlanets[planet.ID] = planet

	return nil
}

// GetByEstimation estimates the fuel consumption for space mission
func (s service) GetByEstimation(id uuid.UUID, capacity int) (float64, error) {
	// checks whether the planet exists in the ma
	planet, ok := exoPlanets[id]
	if !ok {
		return 0, gofrHttp.ErrorEntityNotFound{Name: "planet", Value: id.String()}
	}

	var gravity float64

	// calculate the gravitational force based on the planet's type
	if planet.ExoplanetType == "GasGiant" {
		gravity = 0.5 / math.Pow(planet.Radius, 2)
	} else {
		gravity = planet.Mass / math.Pow(planet.Radius, 2)
	}

	// calculate the estimated fuel
	fuel := planet.DistanceFromEarth / (math.Pow(gravity, 2) * float64(capacity))

	return fuel, nil
}

// Delete deletes the planet from the exoPlanets map
func (s service) Delete(id uuid.UUID) error {
	// checks whether the planet exists in the map
	if _, ok := exoPlanets[id]; ok {
		// deletes the value for given planetID
		delete(exoPlanets, id)
		return nil
	}

	return gofrHttp.ErrorEntityNotFound{Name: "planet", Value: id.String()}
}

// checks the mandatory fields for the provided planet
func checkMandatory(planet models.Planet) error {
	if planet.Name == "" {
		return gofrHttp.ErrorMissingParam{Params: []string{"name"}}
	}

	if planet.Description == "" {
		return gofrHttp.ErrorMissingParam{Params: []string{"description"}}
	}

	if planet.DistanceFromEarth == 0 {
		return gofrHttp.ErrorMissingParam{Params: []string{"distanceFromEarth"}}
	}

	if planet.Radius == 0 {
		return gofrHttp.ErrorMissingParam{Params: []string{"radius"}}
	}

	if planet.ExoplanetType == "" {
		return gofrHttp.ErrorMissingParam{Params: []string{"exoPlanetData"}}
	}

	if planet.Mass == 0 && planet.ExoplanetType == "Terrestrial" {
		return gofrHttp.ErrorMissingParam{Params: []string{"mass"}}
	}

	return nil
}

// validates the  required parameters for the provided planet
func validatePlanet(planet models.Planet) error {
	if planet.DistanceFromEarth < 10 || planet.DistanceFromEarth > 1000 {
		return customError{"distance must be between 10 and 1000 light years"}
	}

	if planet.ExoplanetType == "Terrestrial" && (planet.Mass < 0.1 || planet.Mass > 10) {
		return customError{"mass must be between 0.1 and 10 Earth-Mass units"}
	}

	if planet.Radius < 0.1 || planet.Radius > 10 {
		return customError{"radius must be between 0.1 and 10 Earth-radius units"}
	}

	if planet.ExoplanetType == "GasGiant" && planet.Mass != 0.0 {
		return customError{"mass exist only in the terrestrial ExoPlanet"}
	}

	return nil
}

type customError struct {
	error string
}

func (c customError) Error() string {
	return fmt.Sprintf("custom error: %s", c.error)
}

func (c customError) StatusCode() int {
	return http.StatusBadRequest
}
