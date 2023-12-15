package handler

import (
	"app/internal"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// VehicleJSON is an struct that represents a vehicle in json format.
type VehicleJSON struct {
	ID           int     `json:"id"`
	Brand        string  `json:"brand"`
	Model        string  `json:"model"`
	Registration string  `json:"registration"`
	Year         int     `json:"year"`
	Color        string  `json:"color"`
	MaxSpeed     int     `json:"max_speed"`
	FuelType     string  `json:"fuel_type"`
	Transmission string  `json:"transmission"`
	Passengers   int     `json:"passengers"`
	Height       float64 `json:"height"`
	Width        float64 `json:"width"`
	Weight       float64 `json:"weight"`
}

// NewVehicleDefault returns a new instance of a vehicle handler.
func NewVehicleDefault(sv internal.ServiceVehicle) *VehicleDefault {
	return &VehicleDefault{sv: sv}
}

// VehicleDefault is an struct that contains handlers for vehicle.
type VehicleDefault struct {
	sv internal.ServiceVehicle
}

// GetAll returns all vehicles.
func (c *VehicleDefault) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// request
		// ...

		// process
		// - get all vehicles from the service
		vehicles, err := c.sv.FindAll()
		if err != nil {
			switch {
			case errors.Is(err, internal.ErrServiceVehicleNotFound):
				ctx.JSON(http.StatusNotFound, map[string]any{"message": "vehicles not found"})
			default:
				ctx.JSON(http.StatusInternalServerError, map[string]any{"message": "internal server error"})
			}
			return
		}

		// response
		// - serialize vehicles
		data := make([]VehicleJSON, len(vehicles))
		for i, vehicle := range vehicles {
			data[i] = VehicleJSON{
				ID:           vehicle.ID,
				Brand:        vehicle.Attributes.Brand,
				Model:        vehicle.Attributes.Model,
				Registration: vehicle.Attributes.Registration,
				Year:         vehicle.Attributes.Year,
				Color:        vehicle.Attributes.Color,
				MaxSpeed:     vehicle.Attributes.MaxSpeed,
				FuelType:     vehicle.Attributes.FuelType,
				Transmission: vehicle.Attributes.Transmission,
				Passengers:   vehicle.Attributes.Passengers,
				Height:       vehicle.Attributes.Height,
				Width:        vehicle.Attributes.Width,
				Weight:       vehicle.Attributes.Weight,
			}
		}
		ctx.JSON(http.StatusOK, map[string]any{"message": "success to find vehicles", "data": data})
	}
}

func (c *VehicleDefault) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// request
		// - deserialize request body
		var req VehicleJSON
		err := ctx.ShouldBindJSON(&req)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, map[string]any{"message": "bad request"})
			return
		}

		// process
		// - create vehicle
		vehicle, err := c.sv.Create(internal.Vehicle{
			Attributes: internal.VehicleAttributes{
				Brand:        req.Brand,
				Model:        req.Model,
				Registration: req.Registration,
				Year:         req.Year,
				Color:        req.Color,
				MaxSpeed:     req.MaxSpeed,
				FuelType:     req.FuelType,
				Transmission: req.Transmission,
				Passengers:   req.Passengers,
				Height:       req.Height,
				Width:        req.Width,
				Weight:       req.Weight,
			},
		})
		if err != nil {
			switch {
			case errors.Is(err, internal.ErrServiceVehicleAlreadyExists):
				ctx.JSON(http.StatusConflict, map[string]any{"message": "Identificador del vehículo ya existente."})
			case errors.Is(err, internal.ErrServiceDataIsInvalid):
				ctx.JSON(http.StatusBadRequest, map[string]any{"message": "Datos del vehículo mal formados o incompletos."})
			default:
				ctx.JSON(http.StatusInternalServerError, map[string]any{"message": "internal server error"})
			}
			return
		}

		// response
		// - serialize vehicle
		data := VehicleJSON{
			ID:           vehicle.ID,
			Brand:        vehicle.Attributes.Brand,
			Model:        vehicle.Attributes.Model,
			Registration: vehicle.Attributes.Registration,
			Year:         vehicle.Attributes.Year,
			Color:        vehicle.Attributes.Color,
			MaxSpeed:     vehicle.Attributes.MaxSpeed,
			FuelType:     vehicle.Attributes.FuelType,
			Transmission: vehicle.Attributes.Transmission,
			Passengers:   vehicle.Attributes.Passengers,
			Height:       vehicle.Attributes.Height,
			Width:        vehicle.Attributes.Width,
			Weight:       vehicle.Attributes.Weight,
		}
		ctx.JSON(http.StatusCreated, map[string]any{"message": "success to create vehicle", "data": data})
	}
}

// Search func for search by color and year
func (c *VehicleDefault) Search() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// request
		// - deserialize request body
		color := ctx.Param("color")
		year, err := strconv.Atoi(ctx.Param("year"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, map[string]any{"message": "bad request"})
			return
		}

		// process
		// - search vehicle
		vehicles, err := c.sv.Find(color, year)
		if err != nil {
			switch {
			case errors.Is(err, internal.ErrServiceVehicleNotFound):
				ctx.JSON(http.StatusNotFound, map[string]any{"message": "No se encontraron vehículos con esos criterios."})
			default:
				ctx.JSON(http.StatusInternalServerError, map[string]any{"message": "internal server error"})
			}
			return
		}

		// response
		// - serialize vehicles
		data := make([]VehicleJSON, len(vehicles))
		for i, vehicle := range vehicles {
			data[i] = VehicleJSON{
				ID:           vehicle.ID,
				Brand:        vehicle.Attributes.Brand,
				Model:        vehicle.Attributes.Model,
				Registration: vehicle.Attributes.Registration,
				Year:         vehicle.Attributes.Year,
				Color:        vehicle.Attributes.Color,
				MaxSpeed:     vehicle.Attributes.MaxSpeed,
				FuelType:     vehicle.Attributes.FuelType,
				Transmission: vehicle.Attributes.Transmission,
				Passengers:   vehicle.Attributes.Passengers,
				Height:       vehicle.Attributes.Height,
				Width:        vehicle.Attributes.Width,
				Weight:       vehicle.Attributes.Weight,
			}
		}
		ctx.JSON(http.StatusOK, map[string]any{"message": "success to find vehicles", "data": data})
	}
}

func (c *VehicleDefault) ByBrandRangeYear() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// request
		// - deserialize request body
		brand := ctx.Param("brand")
		yearStart, err := strconv.Atoi(ctx.Param("year_start"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, map[string]any{"message": "bad request"})
			return
		}
		yearEnd, err := strconv.Atoi(ctx.Param("year_end"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, map[string]any{"message": "bad request"})
			return
		}

		// process
		// - search vehicle
		vehicles, err := c.sv.FindByBrandRangeYear(brand, yearStart, yearEnd)
		if err != nil {
			switch {
			case errors.Is(err, internal.ErrServiceVehicleNotFound):
				ctx.JSON(http.StatusNotFound, map[string]any{"message": "No se encontraron vehiculos con los criterios."})
			default:
				ctx.JSON(http.StatusInternalServerError, map[string]any{"message": "internal server error"})
			}
			return
		}

		// response
		// - serialize vehicles
		data := make([]VehicleJSON, len(vehicles))
		for i, vehicle := range vehicles {
			data[i] = VehicleJSON{
				ID:           vehicle.ID,
				Brand:        vehicle.Attributes.Brand,
				Model:        vehicle.Attributes.Model,
				Registration: vehicle.Attributes.Registration,
				Year:         vehicle.Attributes.Year,
				Color:        vehicle.Attributes.Color,
				MaxSpeed:     vehicle.Attributes.MaxSpeed,
				FuelType:     vehicle.Attributes.FuelType,
				Transmission: vehicle.Attributes.Transmission,
				Passengers:   vehicle.Attributes.Passengers,
				Height:       vehicle.Attributes.Height,
				Width:        vehicle.Attributes.Width,
				Weight:       vehicle.Attributes.Weight,
			}
		}
		ctx.JSON(http.StatusOK, map[string]any{"message": "success to find vehicles", "data": data})
	}
}

func (c *VehicleDefault) AverageSpeedByBrand() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// request
		// - deserialize request body
		brand := ctx.Param("brand")

		// process
		// - search vehicle
		average, err := c.sv.AverageSpeedByBrand(brand)
		if err != nil {
			switch {
			case errors.Is(err, internal.ErrServiceVehicleNotFound):
				ctx.JSON(http.StatusNotFound, map[string]any{"message": "No se encontraron vehiculos de esa marca"})
			default:
				ctx.JSON(http.StatusInternalServerError, map[string]any{"message": "internal server error"})
			}
			return
		}

		// response
		// - serialize vehicles
		ctx.JSON(http.StatusOK, map[string]any{"message": "success to find vehicles", "Average": average})
	}
}

// UpdateMaxSpeed func to update vehicle by id and receive a speed and update the max speed of the vehicle
func (c *VehicleDefault) UpdateMaxSpeed() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// request
		// - deserialize request body
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, map[string]any{"message": "bad request"})
			return
		}

		speed, err := strconv.Atoi(ctx.Param("speed"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, map[string]any{"message": "velocidad mal formada o fuera de rango"})
		}

		// process
		// - update vehicle
		vehicle, err := c.sv.UpdateSpeed(id, speed)
		if err != nil {
			switch {
			case errors.Is(err, internal.ErrServiceVehicleNotFound):
				ctx.JSON(http.StatusNotFound, map[string]any{"message": "No se encontro el vehiculo"})
			default:
				ctx.JSON(http.StatusInternalServerError, map[string]any{"message": "internal server error"})
			}
			return
		}

		// response
		// - serialize vehicle
		data := VehicleJSON{
			ID:           vehicle.ID,
			Brand:        vehicle.Attributes.Brand,
			Model:        vehicle.Attributes.Model,
			Registration: vehicle.Attributes.Registration,
			Year:         vehicle.Attributes.Year,
			Color:        vehicle.Attributes.Color,
			MaxSpeed:     vehicle.Attributes.MaxSpeed,
			FuelType:     vehicle.Attributes.FuelType,
			Transmission: vehicle.Attributes.Transmission,
			Passengers:   vehicle.Attributes.Passengers,
			Height:       vehicle.Attributes.Height,
			Width:        vehicle.Attributes.Width,
			Weight:       vehicle.Attributes.Weight,
		}
		ctx.JSON(http.StatusOK, map[string]any{"message": "Velocidad del vehiculo actualizada exitosamente", "data": data})
	}
}

func (c *VehicleDefault) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// request
		// - deserialize request body
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, map[string]any{"message": "bad request"})
			return
		}

		// process
		// - delete vehicle
		vehicle, err := c.sv.Delete(id)
		if err != nil {
			switch {
			case errors.Is(err, internal.ErrServiceVehicleNotFound):
				ctx.JSON(http.StatusNotFound, map[string]any{"message": "No se encontro el vehiculo"})
			default:
				ctx.JSON(http.StatusInternalServerError, map[string]any{"message": "internal server error"})
			}
			return
		}

		// response
		// - serialize vehicle
		data := VehicleJSON{
			ID:           vehicle.ID,
			Brand:        vehicle.Attributes.Brand,
			Model:        vehicle.Attributes.Model,
			Registration: vehicle.Attributes.Registration,
			Year:         vehicle.Attributes.Year,
			Color:        vehicle.Attributes.Color,
			MaxSpeed:     vehicle.Attributes.MaxSpeed,
			FuelType:     vehicle.Attributes.FuelType,
			Transmission: vehicle.Attributes.Transmission,
			Passengers:   vehicle.Attributes.Passengers,
			Height:       vehicle.Attributes.Height,
			Width:        vehicle.Attributes.Width,
			Weight:       vehicle.Attributes.Weight,
		}

		ctx.JSON(http.StatusOK, map[string]any{"message": "success to delete vehicle", "data": data})
	}
}

// ByFuelType func to get vehicles by fuel type
func (c *VehicleDefault) ByFuelType() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// request
		// - deserialize request body
		fuelType := ctx.Param("fuel_type")

		// process
		// - search vehicle
		vehicles, err := c.sv.FindByFuelType(fuelType)
		if err != nil {
			switch {
			case errors.Is(err, internal.ErrServiceVehicleNotFound):
				ctx.JSON(http.StatusNotFound, map[string]any{"message": "No se encontraron vehículos con ese tipo de combustible."})
			default:
				ctx.JSON(http.StatusInternalServerError, map[string]any{"message": "internal server error"})
			}
			return
		}

		// response
		// - serialize vehicles
		data := make([]VehicleJSON, len(vehicles))
		for i, vehicle := range vehicles {
			data[i] = VehicleJSON{
				ID:           vehicle.ID,
				Brand:        vehicle.Attributes.Brand,
				Model:        vehicle.Attributes.Model,
				Registration: vehicle.Attributes.Registration,
				Year:         vehicle.Attributes.Year,
				Color:        vehicle.Attributes.Color,
				MaxSpeed:     vehicle.Attributes.MaxSpeed,
				FuelType:     vehicle.Attributes.FuelType,
				Transmission: vehicle.Attributes.Transmission,
				Passengers:   vehicle.Attributes.Passengers,
				Height:       vehicle.Attributes.Height,
				Width:        vehicle.Attributes.Width,
				Weight:       vehicle.Attributes.Weight,
			}
		}
		ctx.JSON(http.StatusOK, map[string]any{"message": "success to find vehicles", "data": data})
	}
}

// ByTransmission func to get vehicles by transmission
func (c *VehicleDefault) ByTransmission() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		transmision := ctx.Param("type")

		vehicles, err := c.sv.FindByTransmissionType(transmision)
		if err != nil {
			switch {
			case errors.Is(err, internal.ErrServiceVehicleNotFound):
				ctx.JSON(http.StatusNotFound, map[string]any{"message": "No se encontraron vehículos con ese tipo de transmisión."})
			default:
				ctx.JSON(http.StatusInternalServerError, map[string]any{"message": "internal server error"})
			}
			return
		}

		data := make([]VehicleJSON, len(vehicles))
		for i, vehicle := range vehicles {
			data[i] = VehicleJSON{
				ID:           vehicle.ID,
				Brand:        vehicle.Attributes.Brand,
				Model:        vehicle.Attributes.Model,
				Registration: vehicle.Attributes.Registration,
				Year:         vehicle.Attributes.Year,
				Color:        vehicle.Attributes.Color,
				MaxSpeed:     vehicle.Attributes.MaxSpeed,
				FuelType:     vehicle.Attributes.FuelType,
				Transmission: vehicle.Attributes.Transmission,
				Passengers:   vehicle.Attributes.Passengers,
				Height:       vehicle.Attributes.Height,
				Width:        vehicle.Attributes.Width,
				Weight:       vehicle.Attributes.Weight,
			}
		}
		ctx.JSON(http.StatusOK, map[string]any{"message": "success to find vehicles", "data": data})
	}
}

func (c *VehicleDefault) CreateBatch() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// request
		// - deserialize request body
		var req []VehicleJSON
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, map[string]any{"message": "bad request"})
			return
		}

		// process
		// - create vehicle
		vehiclesInsert := make([]internal.Vehicle, len(req))
		for i, r := range req {
			vehiclesInsert[i] = internal.Vehicle{
				ID:           r.ID,
				Brand:        r.Brand,
				Model:        r.Model,
				Registration: r.Registration,
				Year:         r.Year,
				Color:        r.Color,
				MaxSpeed:     r.MaxSpeed,
				FuelType:     r.FuelType,
				Transmission: r.Transmission,
				Passengers:   r.Passengers,
				Height:       r.Height,
				Width:        r.Width,
				Weight:       r.Weight,
			}
		}
		newVehicles, err := c.sv.CreateBatch(vehiclesInsert)
		if err != nil {
			switch {
			case errors.Is(err, internal.ErrServiceVehicleNotFound):
				ctx.JSON(http.StatusNotFound, map[string]any{"message": "No se encontraron vehículos con ese tipo de combustible."})
			default:
				ctx.JSON(http.StatusInternalServerError, map[string]any{"message": "internal server error"})
			}
			return
		}
		newVehicleJSON := make([]VehicleJSON, len(newVehicles))
		for i, v := range newVehicles {
			newVehicleJSON[i] = VehicleJSON{
				ID:           v.ID,
				Brand:        v.Brand,
				Model:        v.Model,
				Registration: v.Registration,
				Year:         v.Year,
				Color:        v.Color,
				MaxSpeed:     v.MaxSpeed,
				FuelType:     v.FuelType,
				Transmission: v.Transmission,
				Passengers:   v.Passengers,
				Height:       v.Height,
				Width:        v.Width,
				Weight:       v.Weight,
			}
		}
		ctx.JSON(http.StatusOK, map[string]any{"message": "success to create vehicle", "data": newVehicleJSON})
	}
}
