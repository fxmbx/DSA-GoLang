package example

import "fmt"

type Driven interface {
	Drive()
}

type Car struct {
}

func (car *Car) Drive() {
	fmt.Println("Driver is driving the car")
}

type Driver struct {
	Age int
}

type CarProxy struct {
	Car    *Car
	Driver *Driver
}

func (carProcy *CarProxy) Drive() {
	if carProcy.Driver.Age < 18 {
		fmt.Println("infant no suppose de drive")
	} else {
		carProcy.Car.Drive()
	}
}

func NewCarProxy(driverAge int) *CarProxy {
	return &CarProxy{
		Car: &Car{},
		Driver: &Driver{
			Age: driverAge,
		},
	}
}
