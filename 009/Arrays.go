package main

import "fmt"

func main() {
	// Arrays have a fixed length
	var numbers [5]int
	numbers[0] = 10
	numbers[1] = 20
	numbers[2] = 30
	numbers[3] = 40
	numbers[4] = 50

	// Slices are dynamic arrays
	slice := numbers[1:4] // Create a slice referencing a portion of the array

	// Append to a slice
	slice = append(slice, 60)

	// Access elements using index
	fmt.Println(slice[0]) // Output: 20

	// Get the length of the slice
	fmt.Println(len(slice)) // Output: 4

	// Get the capacity of the slice
	fmt.Println(cap(slice)) // Output: 5 (usually the same as the length initially)

	// Create a slice with initial length and capacity
	mySlice := make([]int, 2, 5) // Length: 2, Capacity: 5

	// Append to the slice
	mySlice = append(mySlice, 30, 40, 50)

	// Access elements and length
	fmt.Println(mySlice[0])   // Output: 0
	fmt.Println(len(mySlice)) // Output: 5
	fmt.Println(cap(mySlice)) // Output: 5 (might be different depending on implementation)
}
