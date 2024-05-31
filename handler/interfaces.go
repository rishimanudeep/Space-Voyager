package handler

import (
	"Exo-planet/models"
	"github.com/google/uuid"
)

//go:generate mockgen -destination=mock_interfaces.go -package=handlers -source=interfaces.go

type PlanetService interface {
	Create(planet models.Planet) (*models.Planet, error)
	GetAll(filter models.Filter) ([]models.Planet, error)
	GetByID(id uuid.UUID) (*models.Planet, error)
	Update(planet models.Planet) error
	Delete(id uuid.UUID) error
	GetByEstimation(id uuid.UUID, capacity int) (float64, error)
}
