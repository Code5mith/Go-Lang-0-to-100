package main

import "fmt"

// Operation is a function type that takes two integers and returns an integer
type Operation func(int, int) int

// Calculate performs an operation on two numbers using a provided function
func Calculate(x, y int, op Operation) int {
	return op(x, y)
}

func main() {
	// Define some operations
	add := func(x, y int) int { return x + y }
	subtract := func(x, y int) int { return x - y }
	multiply := func(x, y int) int { return x * y }
	divide := func(x, y int) int { return x / y }

	// Use the Calculate function with different operations
	result1 := Calculate(5, 3, add)
	result2 := Calculate(10, 2, divide)

	fmt.Println("Addition:", result1)
	fmt.Println("Division:", result2)
}
