package main

import "fmt"

// car is an interface that defines the methods that all cars must have
type car interface {
	// addDoors method is used to add doors to the car
	// This is an example of the Builder pattern
	addDoors(doors uint8) car
	// addEngine method is used to add an engine to the car
	// This is an example of the Builder pattern
	addEngine(engine string) car
	// drive method is used to move the car
	drive()
	// stop method is used to stop the car
	stop()
}

// lambo is a struct that represents a Lamborghini car
type lambo struct {
	doors  uint8
	engine string
}

// newLambo is a function that returns a new lambo struct
// This implements the Factory pattern
func newLambo() car {
	return &lambo{}
}

func (l *lambo) addDoors(doors uint8) car {
	l.doors = doors
	return l
}

func (l *lambo) addEngine(engine string) car {
	l.engine = engine
	return l
}

func (l *lambo) drive() {
	fmt.Println("Lambo on the move")
}

func (l *lambo) stop() {
	fmt.Println("Lambo stopped")
}

// chevy is a struct that represents a Chevrolet car
type chevy struct {
	doors  uint8
	engine string
}

// newChevy is a function that returns a new chevy struct
// This implements the Factory pattern
func newChevy() car {
	return &chevy{}
}

func (c *chevy) addDoors(doors uint8) car {
	c.doors = doors
	return c
}

func (c *chevy) addEngine(engine string) car {
	c.engine = engine
	return c
}

func (c *chevy) drive() {
	fmt.Println("Chevy on the move")
}

func (c *chevy) stop() {
	fmt.Println("Chevy stopped")
}

// factory is an interface that defines the methods that all factories must have
// This implements the Abstract Factory pattern
type factory interface {
	makeCar(car car, doors uint8, engine string) car
}

type carFactory struct{}

func (cf *carFactory) makeCar(car car, doors uint8, engine string) car {
	switch car.(type) {
	case *lambo:
		fmt.Println("Making a Lambo")
	case *chevy:
		fmt.Println("Making a Chevy")
	}

	return car.addDoors(doors).addEngine(engine)
}

// newCarFactory is a function that returns a new carFactory struct
// This implements the Factory pattern
func newCarFactory() factory {
	return &carFactory{}
}

func main() {
	var f = newCarFactory()
	var l = newLambo()
	var c = newChevy()
	f.makeCar(l, 2, "V8")
	f.makeCar(c, 4, "V6")
	l.drive()
	c.drive()
	l.stop()
	c.stop()
}
