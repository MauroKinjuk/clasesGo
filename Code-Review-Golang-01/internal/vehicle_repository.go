package internal

import (
	"errors"
)

var (
	// ErrRepositoryVehicleNotFound is returned when a vehicle is not found.
	ErrRepositoryVehicleNotFound      = errors.New("repository: vehicle not found")
	ErrRepositoryVehicleAlreadyExists = errors.New("service: vehicle already exists")
	ErrRepositoryDataIsInvalid        = errors.New("service: data is invalid")
)

// RepositoryVehicle is the interface that wraps the basic methods for a vehicle repository.
type RepositoryVehicle interface {
	// FindAll returns all vehicles
	FindAll() (v []Vehicle, err error)
	Create(v Vehicle) (Vehicle, error)
	Find(color string, year int) (v []Vehicle, err error)
	FindByBrandRangeYear(brand string, year int, year2 int) (v []Vehicle, err error)
	AverageSpeedByBrand(brand string) (average float64, err error)
	UpdateSpeed(id int, speed int) (v Vehicle, err error)
	Delete(id int) (v Vehicle, err error)
	FindByFuelType(fuelType string) (v []Vehicle, err error)
	FindByTransmissionType(transmissionType string) (v []Vehicle, err error)
	CreateBatch(v []Vehicle) (vs []Vehicle, err error)
}
