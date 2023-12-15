package service

import (
	"app/internal"
	"errors"
	"fmt"
)

// NewDefault returns a new instance of a vehicle service.
func NewDefault(rp internal.RepositoryVehicle) *Default {
	return &Default{rp: rp}
}

// Default is an struct that represents a vehicle service.
type Default struct {
	rp internal.RepositoryVehicle
}

// FindAll returns all vehicles.
func (s *Default) FindAll() (v []internal.Vehicle, err error) {
	// get all vehicles from the repository
	v, err = s.rp.FindAll()
	if err != nil {
		if errors.Is(err, internal.ErrRepositoryVehicleNotFound) {
			err = fmt.Errorf("%w. %v", internal.ErrServiceVehicleNotFound, err)
			return
		}
		return
	}

	return
}

// Create creates a new vehicle.
func (s *Default) Create(v internal.Vehicle) (internal.Vehicle, error) {
	// create a new vehicle in the repository
	v, err := s.rp.Create(v)
	if err != nil {
		//check different errors
		if errors.Is(err, internal.ErrRepositoryVehicleAlreadyExists) {
			err = fmt.Errorf("%w. %v", internal.ErrServiceVehicleAlreadyExists, err)
			return internal.Vehicle{}, err
		}
		if errors.Is(err, internal.ErrRepositoryDataIsInvalid) {
			err = fmt.Errorf("%w. %v", internal.ErrServiceDataIsInvalid, err)
			return internal.Vehicle{}, err
		}
		if errors.Is(err, internal.ErrRepositoryVehicleNotFound) {
			err = fmt.Errorf("%w. %v", internal.ErrServiceVehicleNotFound, err)
			return internal.Vehicle{}, err
		}
		return internal.Vehicle{}, err
	}

	return v, nil
}

// Find returns all vehicles by color and year.
func (s *Default) Find(color string, year int) (v []internal.Vehicle, err error) {
	// get all vehicles from the repository
	v, err = s.rp.Find(color, year)
	if err != nil {
		if errors.Is(err, internal.ErrRepositoryVehicleNotFound) {
			err = fmt.Errorf("%w. %v", internal.ErrServiceVehicleNotFound, err)
			return
		}
		return
	}

	return
}

func (s *Default) FindByBrandRangeYear(brand string, year int, year2 int) (v []internal.Vehicle, err error) {
	// get all vehicles from the repository
	v, err = s.rp.FindByBrandRangeYear(brand, year, year2)
	if err != nil {
		if errors.Is(err, internal.ErrRepositoryVehicleNotFound) {
			err = fmt.Errorf("%w. %v", internal.ErrServiceVehicleNotFound, err)
			return
		}
		return
	}

	return
}

func (s *Default) AverageSpeedByBrand(brand string) (average float64, err error) {
	// get all vehicles from the repository
	average, err = s.rp.AverageSpeedByBrand(brand)
	if err != nil {
		if errors.Is(err, internal.ErrRepositoryVehicleNotFound) {
			err = fmt.Errorf("%w. %v", internal.ErrServiceVehicleNotFound, err)
			return
		}
		return
	}

	return
}

func (s *Default) UpdateSpeed(id int, speed int) (v internal.Vehicle, err error) {
	// get all vehicles from the repository
	v, err = s.rp.UpdateSpeed(id, speed)
	if err != nil {
		if errors.Is(err, internal.ErrRepositoryVehicleNotFound) {
			err = fmt.Errorf("%w. %v", internal.ErrServiceVehicleNotFound, err)
			return
		}
		return
	}

	return
}

func (s *Default) Delete(id int) (v internal.Vehicle, err error) {
	// get all vehicles from the repository
	v, err = s.rp.Delete(id)
	if err != nil {
		if errors.Is(err, internal.ErrRepositoryVehicleNotFound) {
			err = fmt.Errorf("%w. %v", internal.ErrServiceVehicleNotFound, err)
			return
		}
		return
	}

	return
}

func (s *Default) FindByFuelType(fuelType string) (v []internal.Vehicle, err error) {
	// get all vehicles from the repository
	v, err = s.rp.FindByFuelType(fuelType)
	if err != nil {
		if errors.Is(err, internal.ErrRepositoryVehicleNotFound) {
			err = fmt.Errorf("%w. %v", internal.ErrServiceVehicleNotFound, err)
			return
		}
		return
	}

	return
}

func (s *Default) FindByTransmissionType(transmissionType string) (v []internal.Vehicle, err error) {
	// get all vehicles from the repository
	v, err = s.rp.FindByTransmissionType(transmissionType)
	if err != nil {
		if errors.Is(err, internal.ErrRepositoryVehicleNotFound) {
			err = fmt.Errorf("%w. %v", internal.ErrServiceVehicleNotFound, err)
			return
		}
		return
	}

	return
}

func (s *Default) CreateBatch(v []internal.Vehicle) (nvs []internal.Vehicle, err error) {
	// get all vehicles from the repository
	createdVehicles, err := s.rp.CreateBatch(v)
	if err != nil {
		if errors.Is(err, internal.ErrRepositoryVehicleNotFound) {
			err = fmt.Errorf("%w. %v", internal.ErrServiceVehicleNotFound, err)
			return
		}
		return
	}

	return createdVehicles, nil
}
