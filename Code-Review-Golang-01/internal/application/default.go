package application

import (
	"app/internal/handler"
	"app/internal/loader"
	"app/internal/repository"
	"app/internal/service"

	"github.com/gin-gonic/gin"
)

// ConfigDefaultInMemory is an struct that contains the configuration for the default application settings.
type ConfigDefaultInMemory struct {
	// FileLoader is the path to the file that contains the vehicles.
	FileLoader string
	// Addr is the address where the application will be listening.
	Addr string
}

// NewDefaultInMemory returns a new instance of a default application.
func NewDefaultInMemory(c *ConfigDefaultInMemory) *DefaultInMemory {
	// default config
	defaultCfg := &ConfigDefaultInMemory{
		FileLoader: "vehicles.json",
		Addr:       ":8080",
	}
	if c != nil {
		if c.FileLoader != "" {
			defaultCfg.FileLoader = c.FileLoader
		}
		if c.Addr != "" {
			defaultCfg.Addr = c.Addr
		}
	}

	return &DefaultInMemory{
		fileLoader: defaultCfg.FileLoader,
		addr:       defaultCfg.Addr,
	}
}

// DefaultInMemory is an struct that contains the default application settings.
type DefaultInMemory struct {
	// fileLoader is the path to the file that contains the vehicles.
	fileLoader string
	// addr is the address where the application will be listening.
	addr string
}

// Run starts the application.
func (d *DefaultInMemory) Run() (err error) {
	// dependencies initialization
	// loader
	ld := loader.NewVehicleJSON(d.fileLoader)
	data, err := ld.Load()
	if err != nil {
		return
	}

	// repository
	rp := repository.NewVehicleSlice(data.Data, data.LastId)

	// service
	sv := service.NewDefault(rp)

	// handler
	hd := handler.NewVehicleDefault(sv)

	// router
	rt := gin.New()
	// - middlewares
	rt.Use(gin.Logger())
	rt.Use(gin.Recovery())
	// - endpoints
	gr := rt.Group("/vehicles")
	{
		gr.GET("", hd.GetAll())
		// 1 Añadir un vehiculo
		// JSON for test:
		/* {"brand": "asd","model": "Fiero","registration": "6603","year": 1986,"color": "Mauv","max_speed": 85,"fuel_type": "gasoline",
		"transmission": "semi-automatic","passengers": 2,"height": 105.43,"width": 280.28,"weight": 288.8} */
		gr.POST("", hd.Create())

		// 2 Buscar vehiculo por color y año
		//GET /vehicles/color/{color}/year/{year}
		// http://localhost:8080/vehicles/color/Mauv/year/1986
		gr.GET("color/:color/year/:year", hd.Search())

		// 3 Buscar vehículos por marca y rango de años
		//GET /vehicles/brand/{brand}/between/{start_year}/{end_year}
		// http://localhost:8080/vehicles/brand/Lotus/between/2000/2010
		gr.GET("brand/:brand/between/:year_start/:year_end", hd.ByBrandRangeYear())

		// 4 Consultar velocidad promedio por marca
		// GET /vehicles/average_speed/brand/{brand}
		// http://localhost:8080/vehicles/average_speed/brand/Lotus
		gr.GET("average_speed/brand/:brand", hd.AverageSpeedByBrand())

		// 5 Añadir multiples vehiculos
		// POST /vehicles/batch
		// http://localhost:8080/vehicles/batch
		gr.POST("batch", hd.CreateBatch())

		// 6 Actualizar velocidad máxima de un vehículo
		// PUT /vehicles/{id}/update_speed/:speed
		// http://localhost:8080/vehicles/1/update_speed/100
		gr.PUT(":id/update_speed/:speed", hd.UpdateMaxSpeed())

		// 7 Listar vehiculos por tipo de combustible
		// GET /vehicles/fuel_type/{type}
		// http://localhost:8080/vehicles/fuel_type/gasoline
		gr.GET("fuel_type/:type", hd.ByFuelType())

		// 8 Eliminar un vehículo
		// DELETE /vehicles/{id}
		// http://localhost:8080/vehicles/1
		gr.DELETE(":id", hd.Delete())

		// 9 Buscar vehiculos por tipo de transmisión
		// GET /vehicles/transmission/{type}
		// http://localhost:8080/vehicles/transmission/automatic
		gr.GET("transmission/:type", hd.ByTransmission())
	}

	// run application
	err = rt.Run(d.addr)
	if err != nil {
		return
	}

	return
}
