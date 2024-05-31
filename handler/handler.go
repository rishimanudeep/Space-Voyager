package handler

import (
	"strconv"

	"github.com/google/uuid"

	"Exo-planet/models"
	"gofr.dev/pkg/gofr"
)

type handler struct {
	planet PlanetService
}

// New nolint -New is a factory function that returns handler.
func New(planetService PlanetService) handler {
	return handler{planet: planetService}
}

// Create is a method of handler struct that handles the creation of a new planet
func (h handler) Create(ctx *gofr.Context) (interface{}, error) {
	var planet models.Planet

	// binds the request body into planet variable struct
	err := ctx.Bind(&planet)
	if err != nil {
		return nil, err
	}

	// call the create method to create the planet
	resp, err := h.planet.Create(planet)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// GetAll is a method of handler struct that handles the fetching of all planets
// it enables sorting if the sortBy parameter is provided
func (h handler) GetAll(ctx *gofr.Context) (interface{}, error) {
	// sortBy will be taken from the query param
	sortBy := ctx.Param("sortBy")

	var (
		filter models.Filter
		err    error
	)

	// If a sorting parameter is provided
	if sortBy != "" {
		filter, err = models.NewFilter(sortBy)
		if err != nil {
			return nil, err
		}
	}

	// call the GetAll method of the planet service, passing the filter
	resp, err := h.planet.GetAll(filter)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// GetByID is a method of handler struct that handles the fetching planet by ID
func (h handler) GetByID(ctx *gofr.Context) (interface{}, error) {
	// id will be taken from the path param
	id := ctx.PathParam("id")

	// checks whether the string is an uuid or not
	planetID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	// call the GetByID method of the planet service, passing the planetID
	resp, err := h.planet.GetByID(planetID)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// Update is a method of handler struct that handles the update planet by ID
func (h handler) Update(ctx *gofr.Context) (interface{}, error) {
	var planet models.Planet

	// binds the request body into planet variable struct
	err := ctx.Bind(&planet)
	if err != nil {
		return nil, err
	}

	// id will be taken from the path param
	id := ctx.PathParam("id")

	// checks whether the string is an uuid or not
	planet.ID, err = uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	// call the Update method of the planet service, passing the planet struct
	err = h.planet.Update(planet)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

// Delete is a method of handler struct that handles the delete planet by ID
func (h handler) Delete(ctx *gofr.Context) (interface{}, error) {
	// id will be taken from the path param
	id := ctx.PathParam("id")

	// checks whether the string is an uuid or not
	planetID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	// call the Delete method of the planet service, passing the planet id
	err = h.planet.Delete(planetID)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

// GetEstimation is a method of handler struct that handles the estimating the fuel of the planet
// for the given crew capacity
func (h handler) GetEstimation(ctx *gofr.Context) (interface{}, error) {
	// id will be taken from the query param
	crewCapacity := ctx.Param("crewCapacity")
	id := ctx.PathParam("id")

	// checks whether the string is an uuid or not
	planetID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	// checks for the valid integer or not
	capacity, err := strconv.Atoi(crewCapacity)
	if err != nil {
		return nil, err
	}

	// call the GetByEstimation method of the planet service, passing the planet id and capacity
	fuelEstimation, err := h.planet.GetByEstimation(planetID, capacity)
	if err != nil {
		return nil, err
	}

	return fuelEstimation, nil
}
