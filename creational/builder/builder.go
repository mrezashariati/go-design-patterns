/* Builder package implemetns builder design pattern*/
package builder

type BuildProcess interface {
	SetWheels() BuildProcess
	SetSeats() BuildProcess
	SetStructure() BuildProcess
	GetVehicle() VehicleProduct
}

type ManufacturingDirector struct {
	builder BuildProcess
}

func (f *ManufacturingDirector) Construct() {
	f.builder.SetSeats().SetStructure().SetWheels()
}

func (f *ManufacturingDirector) SetBuilder(b BuildProcess) {
	f.builder = b
}

type VehicleProduct struct {
	Wheels    int
	Seats     int
	Structure string
}

type CarBuilder struct {
	v VehicleProduct
}

func (c *CarBuilder) SetWheels() BuildProcess {
	c.v.Wheels = 4
	return c
}

func (c *CarBuilder) SetSeats() BuildProcess {
	c.v.Seats = 5
	return c
}

func (c *CarBuilder) SetStructure() BuildProcess {
	c.v.Structure = "Car"
	return c
}

func (c *CarBuilder) GetVehicle() VehicleProduct {
	return c.v
}

type BikeBuilder struct {
	b VehicleProduct
}

func (c *BikeBuilder) SetWheels() BuildProcess {
	c.b.Wheels = 2
	return c
}

func (c *BikeBuilder) SetSeats() BuildProcess {
	c.b.Seats = 2
	return c
}

func (c *BikeBuilder) SetStructure() BuildProcess {
	c.b.Structure = "Motobike"
	return c
}

func (c *BikeBuilder) GetVehicle() VehicleProduct {
	return c.b
}
