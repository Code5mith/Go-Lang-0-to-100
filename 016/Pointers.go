package main

import "fmt"

func main() {
	// Declare a variable
	x := 42

	// Create a pointer to x
	ptr := &x

	// Print the value of x and its address
	fmt.Println("Value of x:", x)
	fmt.Println("Address of x:", &x)
	fmt.Println("Value of ptr:", ptr) // Will print the same address as &x
	fmt.Println("Value pointed to by ptr:", *ptr)

	// Modify the value through the pointer
	*ptr = 43
	fmt.Println("Value of x after modification:", x)

	// A struct example
	type Person struct {
		Name string
		Age  int
	}

	// Create a person
	person := Person{"Alice", 30}

	// Create a pointer to the person
	personPtr := &person

	// Modify the person's age through the pointer
	personPtr.Age = 31
	fmt.Println("Person's age:", person.Age)
}
