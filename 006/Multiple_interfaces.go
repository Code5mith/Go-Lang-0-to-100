package main

import "fmt"

// **Interface: brand**
// This interface defines a contract for any type that needs to provide a way to retrieve its brand information.
// The `getBrand()` method should return a string representing the brand name.
type brand interface {
	getBrand() string
}

// **Interface: model**
// This interface defines a contract for any type that needs to provide a way to retrieve its model information.
// The `getModel()` method should return a string representing the model name or identifier.
type model interface {
	getModel() string
}

// **Struct: car**
// This struct represents a car object with two fields:
//   - `brand` (string): Stores the car brand name.
//   - `model` (string): Stores the car model name or identifier.
//
// The `car` struct implements both the `brand` and `model` interfaces, fulfilling their respective contracts.
type car struct {
	brand string
	model string
}

// **Method: (c car) getBrand() string**
// This method implements the `getBrand()` method required by the `brand` interface for the `car` struct.
// It returns a formatted string with the label "Car Brand" followed by the car's brand name retrieved from the `c.brand` field.
func (c car) getBrand() string {
	return fmt.Sprintf("Car Brand : %s", c.brand)
}

// **Method: (c car) getModel() string**
// This method implements the `getModel()` method required by the `model` interface for the `car` struct.
// It returns a formatted string with the label "Car Model is" followed by the car's model name retrieved from the `c.model` field.
func (c car) getModel() string {
	return fmt.Sprintf("Car Model is : %s", c.model)
}

//  **Comment: Interface Usage Example**
//  This comment highlights the demonstration of using multiple interfaces for a single struct.
//  The `car` struct implements both `brand` and `model` interfaces, showcasing the ability to have these contracts on one type.

func main() {

	//  **Car Instance**
	//  Create a new `car` instance with the brand set to "BMW" and model set to "2019".
	car := car{
		brand: "BMW",
		model: "2019",
	}

	//  **Function: carInfo**
	//  This function takes two arguments:
	//      - `b brand`: This parameter expects any type that implements the `brand` interface.
	//      - `m model`: This parameter expects any type that implements the `model` interface.
	//  The function calls the `getBrand()` and `getModel()` methods on the provided arguments
	//  (assuming they implement the respective interfaces) and returns both retrieved values.
	carInfo := func(b brand, m model) (string, string) {
		return b.getBrand(), m.getModel()
	}

	//  **Function Call: carInfo**
	//  We call the `carInfo` function with two arguments:
	//      - The `car` instance itself (which implements both `brand` and `model` interfaces).
	//  This demonstrates how the function can work with any object implementing the specified interfaces,
	//  promoting flexibility and code reuse.
	fmt.Printf(carInfo(car, car))
}
