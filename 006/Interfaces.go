package main

import "fmt"

// Define an interface "Shaper" with a method "Area"

// Interface is a collection of method signitures.

type Shaper interface {
	Area() float64 // This method calculates the area of the shape (without implementation details)
}

// Define a "Square" struct with its properties
type Square struct {
	Side float64
}

// Implement the "Shaper" interface for "Square" by defining the Area method
func (s Square) Area() float64 {
	return s.Side * s.Side // Calculate the area of a square (side * side)
}

// Define a "Circle" struct with its properties
type Circle struct {
	Radius float64
}

// Implement the "Shaper" interface for "Circle" by defining the Area method
func (c Circle) Area() float64 {
	return 3.14159 * c.Radius * c.Radius // Calculate the area of a circle (pi * radius^2)
}

func main() {
	// Create a Square instance with a side length
	square := Square{Side: 5}

	// Create a Circle instance with a radius
	circle := Circle{Radius: 10}

	// Define a function "calculateArea" that takes a "Shaper" interface as its parameter
	// This function can work with any object that implements the "Shaper" interface,
	// promoting polymorphism (ability to work with different types)

	calculateArea := func(s Shaper) {
		fmt.Println("Area:", s.Area())
	}

	// Call the "calculateArea" function with the Square instance
	calculateArea(square)

	// Call the "calculateArea" function with the Circle instance (demonstrating polymorphism)
	calculateArea(circle)
}
