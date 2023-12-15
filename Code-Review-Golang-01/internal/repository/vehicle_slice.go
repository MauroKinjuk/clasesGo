package repository

import "app/internal"

// NewVehicleSlice returns a new instance of a vehicle repository in an slice.
func NewVehicleSlice(db []internal.Vehicle, lastId int) *VehicleSlice {
	return &VehicleSlice{
		db:     db,
		lastId: lastId,
	}
}

// VehicleSlice is an struct that represents a vehicle repository in an slice.
type VehicleSlice struct {
	// db is the database of vehicles.
	db []internal.Vehicle
	// lastId is the last id of the database.
	lastId int
}

// Create creates a new vehicle
func (s *VehicleSlice) Create(v internal.Vehicle) (internal.Vehicle, error) {
	// check if the vehicle already exists
	for _, vehicle := range s.db {
		if vehicle == v {
			return internal.Vehicle{}, internal.ErrRepositoryVehicleAlreadyExists
		}
	}

	//Check if data is empty or invalid
	if v.Attributes.Brand == "" || v.Attributes.Model == "" || v.Attributes.Year == 0 || v.Attributes.Color == "" {
		return internal.Vehicle{}, internal.ErrServiceDataIsInvalid
	}
	if v.Attributes.MaxSpeed == 0 || v.Attributes.MaxSpeed < 0 {
		return internal.Vehicle{}, internal.ErrServiceDataIsInvalid
	}
	if v.Attributes.FuelType == "" || v.Attributes.Transmission == "" || v.Attributes.Passengers == 0 {
		return internal.Vehicle{}, internal.ErrServiceDataIsInvalid
	}
	if v.Attributes.Height == 0 || v.Attributes.Width == 0 || v.Attributes.Weight == 0 {
		return internal.Vehicle{}, internal.ErrServiceDataIsInvalid
	}

	// increment the last id
	s.lastId++
	v.ID = s.lastId
	// add the vehicle to the database
	s.db = append(s.db, v)
	return v, nil
}

// Find vehicle by color and year
func (s *VehicleSlice) Find(color string, year int) (v []internal.Vehicle, err error) {
	// check if the database is empty
	if len(s.db) == 0 {
		err = internal.ErrRepositoryVehicleNotFound
		return
	}

	//check if color and year are empty
	if color == "" && year == 0 {
		err = internal.ErrRepositoryDataIsInvalid
		return
	}

	//find vehicle by color and year
	for _, vehicle := range s.db {
		if vehicle.Attributes.Color == color && vehicle.Attributes.Year == year {
			v = append(v, vehicle)
		}
	}

	// check if the vehicle was found
	if len(v) == 0 {
		err = internal.ErrRepositoryVehicleNotFound
		return
	}

	return
}

// FindAll returns all vehicles
func (s *VehicleSlice) FindAll() (v []internal.Vehicle, err error) {
	// check if the database is empty
	if len(s.db) == 0 {
		err = internal.ErrRepositoryVehicleNotFound
		return
	}

	// make a copy of the database
	v = make([]internal.Vehicle, len(s.db))
	copy(v, s.db)
	return
}

func (s *VehicleSlice) FindByBrandRangeYear(brand string, year int, year2 int) (v []internal.Vehicle, err error) {
	// check if the database is empty
	if len(s.db) == 0 {
		err = internal.ErrRepositoryVehicleNotFound
		return
	}

	//check if brand, year and year2 are empty
	if brand == "" && year == 0 && year2 == 0 {
		err = internal.ErrRepositoryDataIsInvalid
		return
	}

	//find vehicle by brand, year and year2
	for _, vehicle := range s.db {
		if vehicle.Attributes.Brand == brand && vehicle.Attributes.Year >= year && vehicle.Attributes.Year <= year2 {
			v = append(v, vehicle)
		}
	}

	// check if the vehicle was found
	if len(v) == 0 {
		err = internal.ErrRepositoryVehicleNotFound
		return
	}

	return
}

func (s *VehicleSlice) AverageSpeedByBrand(brand string) (average float64, err error) {
	// check if the database is empty
	if len(s.db) == 0 {
		err = internal.ErrRepositoryVehicleNotFound
		return
	}

	//check if brand is empty
	if brand == "" {
		err = internal.ErrRepositoryDataIsInvalid
		return
	}

	//find vehicle by brand and calculate the average speed
	count := 0
	for _, vehicle := range s.db {
		if vehicle.Attributes.Brand == brand {
			average += float64(vehicle.Attributes.MaxSpeed)
			count++
		}
	}

	if count == 0 {
		err = internal.ErrRepositoryVehicleNotFound
		return
	}

	average = average / float64(count)
	return
}

func (s *VehicleSlice) UpdateSpeed(id int, speed int) (v internal.Vehicle, err error) {
	// check if the database is empty
	if len(s.db) == 0 {
		err = internal.ErrRepositoryVehicleNotFound
		return
	}

	//check if id and speed are empty
	if id == 0 && speed == 0 {
		err = internal.ErrRepositoryDataIsInvalid
		return
	}

	//find vehicle by id and update the speed
	for i, vehicle := range s.db {
		if vehicle.ID == id {
			s.db[i].Attributes.MaxSpeed = speed
			v = s.db[i]
			return
		}
	}

	// check if the vehicle was found
	if v.ID == 0 {
		err = internal.ErrRepositoryVehicleNotFound
		return
	}

	return
}

func (s *VehicleSlice) Delete(id int) (v internal.Vehicle, err error) {
	// check if the database is empty
	if len(s.db) == 0 {
		err = internal.ErrRepositoryVehicleNotFound
		return
	}

	//check if id is empty
	if id == 0 {
		err = internal.ErrRepositoryDataIsInvalid
		return
	}

	//find vehicle by id and delete it
	for i, vehicle := range s.db {
		if vehicle.ID == id {
			v = s.db[i]
			s.db = append(s.db[:i], s.db[i+1:]...)
			return
		}
	}

	// check if the vehicle was found
	if v.ID == 0 {
		err = internal.ErrRepositoryVehicleNotFound
		return
	}

	return
}

func (s *VehicleSlice) FindByFuelType(fuelType string) (v []internal.Vehicle, err error) {
	// check if the database is empty
	if len(s.db) == 0 {
		err = internal.ErrRepositoryVehicleNotFound
		return
	}

	//check if fuelType is empty
	if fuelType == "" {
		err = internal.ErrRepositoryDataIsInvalid
		return
	}

	//find vehicle by fuelType
	for _, vehicle := range s.db {
		if vehicle.Attributes.FuelType == fuelType {
			v = append(v, vehicle)
		}
	}

	// check if the vehicle was found
	if len(v) == 0 {
		err = internal.ErrRepositoryVehicleNotFound
		return
	}

	return
}

func (s *VehicleSlice) FindByTransmissionType(transmissionType string) (v []internal.Vehicle, err error) {
	// check if the database is empty
	if len(s.db) == 0 {
		err = internal.ErrRepositoryVehicleNotFound
		return
	}

	//check if transmissionType is empty
	if transmissionType == "" {
		err = internal.ErrRepositoryDataIsInvalid
		return
	}

	//find vehicle by transmissionType
	for _, vehicle := range s.db {
		if vehicle.Attributes.Transmission == transmissionType {
			v = append(v, vehicle)
		}
	}

	// check if the vehicle was found
	if len(v) == 0 {
		err = internal.ErrRepositoryVehicleNotFound
		return
	}

	return
}

// CreateBatch creates a batch of vehicles, check if the vehicle already exists and if the data is valid
// and return error and the vehicles created
func (s *VehicleSlice) CreateBatch(v []internal.Vehicle) (vehicles []internal.Vehicle, err error) {
	// check if the database is empty
	if len(s.db) == 0 {
		err = internal.ErrServiceVehicleNotFound
		return
	}

	// check if the data is valid
	for _, vehicle := range v {
		//Check if data is empty or invalid
		if vehicle.Attributes.Brand == "" || vehicle.Attributes.Model == "" || vehicle.Attributes.Year == 0 || vehicle.Attributes.Color == "" {
			err = internal.ErrServiceDataIsInvalid
			return
		}
		if vehicle.Attributes.MaxSpeed == 0 || vehicle.Attributes.MaxSpeed < 0 {
			err = internal.ErrServiceDataIsInvalid
			return
		}
		if vehicle.Attributes.FuelType == "" || vehicle.Attributes.Transmission == "" || vehicle.Attributes.Passengers == 0 {
			err = internal.ErrServiceDataIsInvalid
			return
		}
		if vehicle.Attributes.Height == 0 || vehicle.Attributes.Width == 0 || vehicle.Attributes.Weight == 0 {
			err = internal.ErrServiceDataIsInvalid
			return
		}
	}

	for _, vehicle := range v {
		// check if the vehicle already exists
		for _, v := range s.db {
			if vehicle.ID == v.ID {
				err = internal.ErrRepositoryVehicleAlreadyExists
				return
			}
		}
	}

	// add the vehicle to the database
	s.db = append(s.db, v...)
	vehicles = v
	return
}
