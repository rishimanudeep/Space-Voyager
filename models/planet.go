package models

import (
	"Exo-planet/types"
	"fmt"
	"github.com/google/uuid"
	"net/http"
)

const (
	SortByRadius string = "radius"
	SortByMass   string = "mass"
)

type Planet struct {
	ID                uuid.UUID           `json:"id"`
	Name              string              `json:"name"`
	Description       string              `json:"description"`
	DistanceFromEarth float64             `json:"distance_from_earth"`
	Radius            float64             `json:"radius"`
	Mass              float64             `json:"mass,omitempty"`
	ExoplanetType     types.ExoPlanetType `json:"exoplanet_type"`
}

type Filter struct {
	SortBy string
}

// NewFilter creates a new Filter instance with the specified sorting criteria.
func NewFilter(sortBy string) (Filter, error) {
	var sortCriteria string
	switch sortBy {
	case "radius":
		sortCriteria = SortByRadius
	case "mass":
		sortCriteria = SortByMass
	default:
		//return Filter{}, errors.New("invalid sorting criteria: " + sortBy)
		return Filter{}, customError{"invalid sort parameter:"}
	}
	return Filter{SortBy: sortCriteria}, nil
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
