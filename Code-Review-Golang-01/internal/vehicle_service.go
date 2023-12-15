package internal

import (
	"errors"
)

var (
	// ErrServiceVehicleNotFound is returned when no vehicle is found.
	ErrServiceVehicleNotFound      = errors.New("service: vehicle not found")
	ErrServiceVehicleAlreadyExists = errors.New("service: vehicle already exists")
	ErrServiceDataIsInvalid        = errors.New("service: data is invalid")
)

// ServiceVehicle is the interface that wraps the basic methods for a vehicle service.
// - conections with external apis
// - business logic
type ServiceVehicle interface {
	// FindAll returns all vehicles
	FindAll() (v []Vehicle, err error)
	//Create return a vehicle and error
	Create(v Vehicle) (Vehicle, error)
	//Find vehicle by color and year
	Find(color string, year int) (v []Vehicle, err error)
	//FindByBrandRangeYear vehicle by brand, year and year2
	FindByBrandRangeYear(brand string, year int, year2 int) (v []Vehicle, err error)
	AverageSpeedByBrand(brand string) (average float64, err error)
	UpdateSpeed(id int, speed int) (v Vehicle, err error)
	Delete(id int) (v Vehicle, err error)
	FindByFuelType(fuelType string) (v []Vehicle, err error)
	FindByTransmissionType(transmissionType string) (v []Vehicle, err error)
	CreateBatch(v []Vehicle) (vs []Vehicle, err error)
}
