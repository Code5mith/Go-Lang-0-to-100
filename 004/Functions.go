package main

import (
	"errors"
	"fmt"
)

// calc adds two numbers together
func calc(a, b int) int {
	return a + b
}

// Functions with multiple returns
func get_employee_info(name string, id int) (string, int) {
	return name, id
}

// Function to increase the request limit
func incrememtRequestLimit(currentLimit, updateLimit int) int {
	return currentLimit + updateLimit
}

// Named and Naked returns in Function
func say_hello(first, last string) (first_name, last_name string) {

	first_name = first
	last_name = last

	//We dont need to return this two variables the function will
	// return what we definded in then named return part.
	// This function uses a Naked return.

	return
}
func say_hello2(first, last string) (first_name, last_name string) {

	first_name = first
	last_name = last

	//We dont need to return this two variables the function will
	// return what we definded in then named return part.
	// this method is called a naked return

	// This function explicitilly returns our values.
	return first_name, last_name
}

// MARK: Guard clause aka early return

func divide(dividend, divisor int) (int, error) {
	// Guard clause to handle division by zero
	if divisor == 0 {
		return 0, errors.New("division by zero")
	}

	return dividend / divisor, nil
}

func main() {
	// Find the sum of 7 and 8
	sum := calc(7, 8)
	fmt.Println("The sum of 7 and 8 is:", sum)

	// Variable to store the allowed request limit
	allowedRequests := 100

	// Constant value for update amount (cannot change)
	const updateLimit = 20

	// Go copies values passed to functions (not references)
	incremented := incrememtRequestLimit(allowedRequests, updateLimit) // Call function and store result

	// Original variable won't be changed because of the copy
	fmt.Println("Original allowed requests:", allowedRequests)

	// Update the actual allowedRequests variable
	allowedRequests = incremented

	// Now allowedRequests reflects the updated value
	fmt.Println("Updated allowed requests:", allowedRequests)

	// Naked return
	fmt.Println(say_hello("hanna", "larcon"))

	//Explicit return
	fmt.Println(say_hello2("hanna", "larcon"))
}
