package main

import (
	"fmt"
)

// **Interface: Shape**
// This interface defines a contract for any type that represents a shape.
// It has a single method `Area()` that should return the area of the shape (without implementation specifics).
type Shape interface {
	Area() float64
}

// **Struct: Square**
// This struct represents a square with a single field `side` (float64).
type Square struct {
	side float64
}

// **Method: (s Square) Area() float64**
// This method implements the `Area()` method for the `Square` struct, fulfilling the `Shape` interface contract.
// It calculates and returns the area of the square (side * side).
func (s Square) Area() float64 {
	return s.side * s.side
}

// **Struct: Circle**
// This struct represents a circle with a single field `radius` (float64).
type Circle struct {
	radius float64
}

// **Method: (c Circle) Area() float64**
// This method implements the `Area()` method for the `Circle` struct, fulfilling the `Shape` interface contract.
// It calculates and returns the area of the circle (pi * radius^2).
func (c Circle) Area() float64 {
	return 3.14159 * c.radius * c.radius
}

func main() {
	// **Shapes**
	// Create instances of Square and Circle
	square := Square{side: 5}
	circle := Circle{radius: 10}

	// **Shape Interface Variable**
	// Define a variable `shape` of type `Shape` interface.
	// This variable can hold any object that implements the `Shape` interface.
	var shape Shape

	// **Shape Assignment**
	// Assign the `square` instance to the `shape` variable.
	// Since `Square` implements `Shape`, this assignment is valid.
	shape = square

	// **Type Switch**
	// Use a type switch to check the dynamic type of the value held by the `shape` variable.
	switch s := shape.(type) {
	case Square:
		fmt.Println("Area of Square:", s.Area())
	case Circle:
		fmt.Println("Area of Circle:", s.Area())
	default:
		fmt.Println("Unknown shape type")
	}

	// **Shape Assignment (Circle)**
	// Assign the `circle` instance to the `shape` variable.
	shape = circle

	// **Type Switch (Again)**
	// Perform the type switch again on the updated `shape` variable.
	switch s := shape.(type) {
	case Square:
		fmt.Println("Area of Square:", s.Area()) // This won't be executed
	case Circle:
		fmt.Println("Area of Circle:", s.Area())
	default:
		fmt.Println("Unknown shape type")
	}
}
