package main

import "fmt"

// car is an interface that defines the methods that all cars must have
type car interface {
	// drive method is used to move the car
	drive()

	// stop method is used to stop the car
	stop()
}

type lambo struct {
	Model string
}

type chevy struct {
	Model string
}

var (
	_ car = (*lambo)(nil)
	_ car = (*chevy)(nil)
)

func (l *lambo) drive() {
	fmt.Printf("Lambo of model %s on the move\n", l.Model)
}

func (l *lambo) stop() {
	fmt.Printf("Lambo of model %s stopped\n", l.Model)
}

func (c *chevy) drive() {
	fmt.Printf("Chevy of model %s on the move\n", c.Model)
}

func (c *chevy) stop() {
	fmt.Printf("Chevy of model %s stopped\n", c.Model)
}

// driveCar is a function that takes a car interface and calls the drive method
func driveCar(c car) {
	c.drive()
}

// stopCar is a function that takes a car interface and calls the stop method
func stopCar(c car) {
	c.stop()
}

func main() {
	var l = lambo{
		Model: "Gallardo",
	}
	var c = chevy{
		Model: "Camaro",
	}
	driveCar(&l)
	driveCar(&c)
	stopCar(&l)
	stopCar(&c)
}
