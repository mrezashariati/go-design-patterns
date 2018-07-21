package builder

import "testing"

func TestBuilderPattern(t *testing.T) {
	manufacturingComplex := ManufacturingDirector{}

	safsdfhello

	carBuilder := &CarBuilder{}
	manufacturingComplex.SetBuilder(carBuilder)
	manufacturingComplex.Construct()
	car := carBuilder.GetVehicle()

	if car.Wheels != 4 {
		t.Errorf("Wheels on a car must be 4")
	}
	if car.Structure != "Car" {
		t.Errorf("Structure on a car must be 'Car' ")
	}
	if car.Seats != 5 {
		t.Errorf("Seats on a car must be 5")
	}

	bikeBuilder := &BikeBuilder{}

	manufacturingComplex.SetBuilder(bikeBuilder)
	manufacturingComplex.Construct()

	motobike := bikeBuilder.GetVehicle()
	motobike.Seats = 2

	if motobike.Wheels != 2 {
		t.Errorf("Wheels on a motobike must be 2")
	}
	if motobike.Structure != "Motobike" {
		t.Errorf("Structure on a motobike must be 'Motobike' ")
	}
	if motobike.Seats != 2 {
		t.Errorf("Seats on a motobike must be 2")
	}
}
