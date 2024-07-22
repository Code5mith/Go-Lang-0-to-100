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

// **Function: calculateTotalArea**
// This function takes a slice of `interface{}` as its parameter.
// It iterates through the slice and attempts to perform a type assertion on each element.
//   - If the assertion to the `Shape` interface succeeds (the element actually implements `Shape`),
//     the function calls the `Area()` method and accumulates the area.
//   - If the assertion fails (the element is not a `Shape`), the function logs an error message
//     and skips that element.
func calculateTotalArea(shapes []interface{}) (float64, error) {
	var totalArea float64
	var err error

	for _, shape := range shapes {
		// **Type Assertion: shape (Shape)**
		// This line attempts to assert the current element (`shape`) in the loop to the `Shape` interface.
		// If the assertion is successful, the variable `s` will hold the underlying value as a `Shape` type.
		// If the assertion fails, the `ok` variable will be set to `false` and the error will be stored in `err`.
		s, ok := shape.(Shape)
		if !ok {
			err = fmt.Errorf("invalid shape type: %T", shape)
			fmt.Println(err) // Log the error for debugging purposes
			continue         // Skip to the next element in the loop if assertion fails
		}

		// **Method Call: s.Area()**
		// Since the assertion was successful, we can now safely call the `Area()` method on the `s` variable
		// which is guaranteed to be a `Shape` type. The area is added to the total.
		totalArea += s.Area()
	}

	return totalArea, err
}

func main() {
	// **Shapes Slice**
	// Create a slice of `interface{}` containing a mix of `Square` and `string` (invalid type).
	shapes := []interface{}{Square{side: 5}, "Not a shape", Circle{radius: 10}}

	// **Function Call: calculateTotalArea**
	// Call the `calculateTotalArea` function with the slice of shapes.
	totalArea, err := calculateTotalArea(shapes)

	if err != nil {
		fmt.Println("Error:", err) // Handle the error if encountered during type assertions
	} else {
		fmt.Println("Total Area:", totalArea)
	}
}
